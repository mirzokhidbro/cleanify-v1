package handlers

import (
	"bw-erp/api/http"
	"bw-erp/models"
	"bw-erp/pkg/utils"
	"context"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-telegram/bot"
	tgmodels "github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/keyboard/inline"
	"github.com/google/uuid"
)

var kb = &tgmodels.ReplyKeyboardMarkup{
	Keyboard: [][]tgmodels.KeyboardButton{
		{
			{Text: "Zayavkalar"},
			// {Text: "Tanlanganlar"},
		},
	},
	ResizeKeyboard:  true,
	OneTimeKeyboard: true,
}

func (h *Handler) CreateCompanyBotModel(c *gin.Context) {
	var body models.CreateCompanyBotModel
	if err := c.ShouldBindJSON(&body); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	opts := []bot.Option{
		bot.WithDefaultHandler(h.Handler),
	}

	b, err := bot.New(body.BotToken, opts...)
	if err != nil {
		panic(err)
	}
	botData, err := b.GetMe(c)

	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	body.Firstname = botData.FirstName
	body.Lastname = botData.LastName
	body.Username = botData.Username
	body.BotID = int(botData.ID)

	id := uuid.New()

	err = h.Stg.TelegramBot().Create(id.String(), body)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	h.handleResponse(c, http.Created, id)
}

func (h *Handler) BotStart(c *gin.Context) {
	// bots, err := h.Stg.TelegramBot().GetOrderBot()
	// if err != nil {
	// 	h.handleResponse(c, http.BadRequest, err.Error())
	// 	return
	// }

	go func() {
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
		defer cancel()

		var wg sync.WaitGroup

		// for _, bot_config := range bots {

		opts := []bot.Option{
			bot.WithDefaultHandler(h.Handler),
		}

		b, err := bot.New(h.Cfg.BotToken, opts...)
		if err != nil {
			panic(err)
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			// b.RegisterHandler(bot.HandlerTypeMessageText, "/olishkerak", bot.MatchTypeExact, h.newApplicationHandler)
			b.RegisterHandler(bot.HandlerTypeMessageText, "/code", bot.MatchTypeExact, h.telegramGroupVerificationHandler)
			b.Start(ctx)
		}()
		// }

		wg.Wait()
	}()

	h.handleResponse(c, http.OK, "OK!")
}

func (h *Handler) telegramGroupVerificationHandler(ctx context.Context, b *bot.Bot, update *tgmodels.Update) {
	// botData, _ := b.GetMe(ctx)
	if update.Message.Chat.ID < 0 {
		rand.Seed(time.Now().UnixNano())
		randomNumber := rand.Intn(900000) + 100000

		err := h.Stg.TelegramGroup().Create(models.CreateTelegramGroupRequest{
			ChatID: int(update.Message.Chat.ID),
			Name:   update.Message.Chat.Title,
			Code:   randomNumber,
		})
		if err != nil {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   err.Error(),
			})
		} else {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:    update.Message.Chat.ID,
				Text:      "<code>" + strconv.Itoa(randomNumber) + "</code>",
				ParseMode: tgmodels.ParseModeHTML,
			})
		}
	}
}

// func (h *Handler) newApplicationHandler(ctx context.Context, b *bot.Bot, update *tgmodels.Update) {
// 	botData, _ := b.GetMe(ctx)
// 	botID := botData.ID
// 	user, err := h.Stg.GetBotUserByChatIDModel(update.Message.Chat.ID, botID)
// 	orders, _ := h.Stg.GetOrdersByStatus(user.CompanyID, 83)

// 	if err != nil || user.UserID == nil {
// 		b.SendMessage(ctx, &bot.SendMessageParams{
// 			ChatID: update.Message.Chat.ID,
// 			Text:   "Bu botdan foydalanish uchun avtorizatsiyadan o'tish kerak!",
// 		})
// 	} else {
// 		if err != nil {
// 			b.SendMessage(ctx, &bot.SendMessageParams{
// 				ChatID: update.Message.Chat.ID,
// 				Text:   err.Error(),
// 			})
// 		} else {
// 			if len(orders) != 0 {
// 				kb := inline.New(b)
// 				for _, order := range orders {
// 					buttonName := *order.Address + " || " + order.Phone
// 					kb.Row().Button(buttonName, []byte(order.Phone), h.onInlineKeyboardSelect)
// 				}

// 				b.SendMessage(ctx, &bot.SendMessageParams{
// 					ChatID:      update.Message.Chat.ID,
// 					Text:        "Buyurtmani tanlang",
// 					ReplyMarkup: kb,
// 				})
// 			} else {
// 				b.SendMessage(ctx, &bot.SendMessageParams{
// 					ChatID: update.Message.Chat.ID,
// 					Text:   "Mavjud emas ❌",
// 				})
// 			}
// 		}
// 	}

// }

func (h *Handler) onInlineKeyboardSelect(ctx context.Context, b *bot.Bot, mes tgmodels.InaccessibleMessage, data []byte) {
	botData, _ := b.GetMe(ctx)
	botID := botData.ID
	user, err := h.Stg.BotUser().GetByChatID(mes.Chat.ID, botID)
	if err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: mes.Chat.ID,
			Text:   err.Error(),
		})
	} else {
		order, err := h.Stg.Order().GetByPhone(user.CompanyID, string(data))
		if err != nil {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: mes.Chat.ID,
				Text:   err.Error(),
			})
		} else {
			if order.PhoneNumber != "" {
				botData, _ := b.GetMe(ctx)
				botID := botData.ID
				user, _ := h.Stg.BotUser().GetByChatID(mes.Chat.ID, botID)
				dialogStep := "asked order slug"
				Page := "Order"
				h.Stg.BotUser().Update(models.BotUser{
					UserID:     user.UserID,
					BotID:      int(botID),
					ChatID:     mes.Chat.ID,
					DialogStep: &dialogStep,
					Page:       &Page,
					Firstname:  mes.Chat.FirstName,
					Lastname:   mes.Chat.LastName,
					Username:   mes.Chat.Username,
				})
				err = h.Stg.TelegramSession().Create(models.TelegramSessionModel{
					BotID:   botID,
					ChatID:  mes.Chat.ID,
					OrderID: order.ID,
				})
				if err != nil {
					b.SendMessage(ctx, &bot.SendMessageParams{
						ChatID: mes.Chat.ID,
						Text:   "create telegram session " + err.Error(),
					})
				}
				if err != nil {
					b.SendMessage(ctx, &bot.SendMessageParams{
						ChatID: mes.Chat.ID,
						Text:   err.Error(),
					})
				} else {
					b.SendMessage(ctx, &bot.SendMessageParams{
						ChatID: mes.Chat.ID,
						Text:   string(data) + " tanlandi. Birka nomini kiriting",
					})
				}
			}
		}
	}
}

func (h *Handler) Handler(ctx context.Context, b *bot.Bot, update *tgmodels.Update) {
	botData, _ := b.GetMe(ctx)
	botID := botData.ID
	if update != nil && update.Message != nil {
		user, err := h.Stg.BotUser().GetByChatID(update.Message.Chat.ID, botID)

		if err != nil {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   err.Error(),
			})
			err = h.Stg.BotUser().Create(models.CreateBotUserModel{
				BotID:      int(botData.ID),
				ChatID:     int(update.Message.Chat.ID),
				Page:       "Registration",
				DialogStep: "AskPhoneNumber",
				Firstname:  update.Message.From.FirstName,
				Lastname:   update.Message.From.LastName,
				Username:   update.Message.From.Username,
			})
			if err != nil {
				b.SendMessage(ctx, &bot.SendMessageParams{
					ChatID: update.Message.Chat.ID,
					Text:   err.Error(),
				})
			} else {
				text := "Salom " + update.Message.From.FirstName + " Iltimos tizimdagi telefon raqamingizni kiriting"
				b.SendMessage(ctx, &bot.SendMessageParams{
					ChatID: update.Message.Chat.ID,
					Text:   text,
				})
			}
		} else {
			switch *user.Page {
			case "Registration":
				h.RegistrationPage(ctx, b, update, botID)
			case "Order":
				h.OrderPage(ctx, b, update, user)
				// case "SetLocation":
				// 	h.SetClientLocation(ctx, b, update, user)
			}
		}
	}
}

func (h *Handler) RegistrationPage(ctx context.Context, b *bot.Bot, update *tgmodels.Update, botID int64) {
	phone := update.Message.Text
	chatID := update.Message.Chat.ID
	user, err := h.Stg.BotUser().GetSelectedBotUser(botID, phone)
	if err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   err.Error(),
		})
	} else {
		page := "Order"
		DialogStep := ""
		_, err = h.Stg.BotUser().Update(models.BotUser{
			UserID:     user.UserID,
			BotID:      int(botID),
			ChatID:     chatID,
			Page:       &page,
			DialogStep: &DialogStep,
		})
		if err != nil {
			h.handleError(ctx, b, update, err.Error(), update.Message.MessageThreadID)
		} else {
			text := "Foydalanuvchi muvaffaqiyatli ro'yxatdan o'tkazildi"
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:      update.Message.Chat.ID,
				Text:        text,
				ReplyMarkup: kb,
			})
		}
	}
}

func (h *Handler) OrderPage(ctx context.Context, b *bot.Bot, update *tgmodels.Update, user models.BotUser) {
	if update.Message.Text == "Zayavkalar" && *user.DialogStep == "" {
		h.Applications(ctx, b, update, user)
	} else {
		if *user.DialogStep == "" {

		} else {
			switch *user.DialogStep {
			case "asked order slug":
				// h.AskedOrderSlug(ctx, b, update, user)
			// case "order count asked":
			// 	h.AskedOrderCount(ctx, b, update, user)
			case "order location asked":
				// h.AskedOrderLocation(ctx, b, update, user)
			}
		}
	}
}

func (h *Handler) Applications(ctx context.Context, b *bot.Bot, update *tgmodels.Update, user models.BotUser) {
	botData, _ := b.GetMe(ctx)
	botID := botData.ID
	user, err := h.Stg.BotUser().GetByChatID(update.Message.Chat.ID, botID)
	orders, _ := h.Stg.Order().GetByStatus(user.CompanyID, 83)
	if err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "user " + err.Error(),
		})
	}

	if err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   err.Error(),
		})
	} else {
		if len(orders) != 0 {
			kb := inline.New(b)
			for _, order := range orders {
				buttonName := *order.Address + " || " + order.PhoneNumber
				kb.Row().Button(buttonName, []byte(order.PhoneNumber), h.onInlineKeyboardSelect)
			}

			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:      update.Message.Chat.ID,
				Text:        "Buyurtmani tanlang",
				ReplyMarkup: kb,
			})
		} else {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "Mavjud emas ❌",
			})
		}
	}

}

// func (h *Handler) AskedOrderSlug(ctx context.Context, b *bot.Bot, update *tgmodels.Update, user models.BotUser) {
// 	botData, _ := b.GetMe(ctx)
// 	botID := botData.ID
// 	chatID := update.Message.Chat.ID
// 	dialogStep := "order count asked"
// 	_, err := h.Stg.BotUser().Update(models.BotUser{
// 		UserID:     user.UserID,
// 		BotID:      int(botID),
// 		ChatID:     chatID,
// 		DialogStep: &dialogStep,
// 	})
// 	if err != nil {
// 		b.SendMessage(ctx, &bot.SendMessageParams{
// 			ChatID: update.Message.Chat.ID,
// 			Text:   "user " + err.Error(),
// 		})
// 	} else {
// 		session, err := h.Stg.TelegramSession().GetByChatIDBotID(chatID, botID)
// 		if err != nil {
// 			b.SendMessage(ctx, &bot.SendMessageParams{
// 				ChatID: update.Message.Chat.ID,
// 				Text:   "create telegram session " + err.Error(),
// 			})
// 		} else {

// 			orderID := session.OrderID
// 			h.Stg.Order().Update(&models.UpdateOrderRequest{
// 				ID:   orderID,
// 				Slug: update.Message.Text,
// 			})

// 			b.SendMessage(ctx, &bot.SendMessageParams{
// 				ChatID: update.Message.Chat.ID,
// 				Text:   "Buyurtma sonini kiriting.",
// 			})
// 		}
// 	}
// }

// func (h *Handler) AskedOrderCount(ctx context.Context, b *bot.Bot, update *tgmodels.Update, user models.BotUser) {
// 	botData, _ := b.GetMe(ctx)
// 	botID := botData.ID
// 	chatID := update.Message.Chat.ID

// 	dialogStep := "order location asked"
// 	_, err := h.Stg.UpdateBotUserModel(models.BotUser{
// 		UserID:     user.UserID,
// 		BotID:      int(botID),
// 		ChatID:     chatID,
// 		DialogStep: &dialogStep,
// 	})
// 	if err != nil {
// 		b.SendMessage(ctx, &bot.SendMessageParams{
// 			ChatID: update.Message.Chat.ID,
// 			Text:   err.Error(),
// 		})
// 	} else {
// 		session, err := h.Stg.GetTelegramSessionByChatIDBotID(chatID, botID)
// 		if err != nil {
// 			b.SendMessage(ctx, &bot.SendMessageParams{
// 				ChatID: update.Message.Chat.ID,
// 				Text:   err.Error(),
// 			})
// 		} else {
// 			orderID := session.OrderID
// 			h.Stg.UpdateOrder(&models.UpdateOrderRequest{
// 				ID:    orderID,
// 				Count: update.Message.Text,
// 			})

// 			kb := &tgmodels.ReplyKeyboardMarkup{
// 				Keyboard: [][]tgmodels.KeyboardButton{
// 					{
// 						{Text: "Lokatsiya", RequestLocation: true},
// 					},
// 				},
// 				ResizeKeyboard:  true,
// 				OneTimeKeyboard: true,
// 			}
// 			b.SendMessage(ctx, &bot.SendMessageParams{
// 				ChatID:      update.Message.Chat.ID,
// 				Text:        "Lokatsiya kiriting!",
// 				ReplyMarkup: kb,
// 			})
// 		}
// 	}
// }

// func (h *Handler) AskedOrderLocation(ctx context.Context, b *bot.Bot, update *tgmodels.Update, user models.BotUser) {
// 	botData, _ := b.GetMe(ctx)
// 	botID := botData.ID
// 	chatID := update.Message.Chat.ID

// 	session, err := h.Stg.TelegramSession().GetByChatIDBotID(chatID, botID)
// 	if err != nil {
// 		b.SendMessage(ctx, &bot.SendMessageParams{
// 			ChatID: update.Message.Chat.ID,
// 			Text:   err.Error(),
// 		})
// 	} else {
// 		if update.Message.Location == nil {
// 			kb := &tgmodels.ReplyKeyboardMarkup{
// 				Keyboard: [][]tgmodels.KeyboardButton{
// 					{
// 						{Text: "Lokatsiya", RequestLocation: true},
// 					},
// 				},
// 				ResizeKeyboard:  true,
// 				OneTimeKeyboard: true,
// 			}
// 			b.SendMessage(ctx, &bot.SendMessageParams{
// 				ChatID:      update.Message.Chat.ID,
// 				Text:        "Iltimos lokatsiya kiriting!",
// 				ReplyMarkup: kb,
// 			})
// 		} else {
// 			_, err := h.Stg.Order().Update(&models.UpdateOrderRequest{
// 				ID:        session.OrderID,
// 				Latitute:  update.Message.Location.Latitude,
// 				Longitude: update.Message.Location.Longitude,
// 				Status:    1,
// 			})
// 			if err != nil {
// 				b.SendMessage(ctx, &bot.SendMessageParams{
// 					ChatID: update.Message.Chat.ID,
// 					Text:   err.Error(),
// 				})
// 			} else {
// 				dialogStep := ""
// 				_, err := h.Stg.BotUser().Update(models.BotUser{
// 					UserID:     user.UserID,
// 					BotID:      int(botID),
// 					ChatID:     update.Message.Chat.ID,
// 					DialogStep: &dialogStep,
// 				})
// 				if err != nil {
// 					b.SendMessage(ctx, &bot.SendMessageParams{
// 						ChatID: update.Message.Chat.ID,
// 						Text:   err.Error(),
// 					})
// 				} else {
// 					h.Stg.TelegramSession().Delete(session.ID)
// 					b.SendMessage(ctx, &bot.SendMessageParams{
// 						ChatID:      update.Message.Chat.ID,
// 						Text:        "Buyurtma qabul qilindi",
// 						ReplyMarkup: kb,
// 					})
// 				}
// 			}
// 		}
// 	}
// }

// func (h *Handler) SetClientLocation(ctx context.Context, b *bot.Bot, update *tgmodels.Update, user models.BotUser) {
// 	botData, _ := b.GetMe(ctx)
// 	botID := botData.ID
// 	chatID := update.Message.Chat.ID

// 	session, err := h.Stg.TelegramSession().GetByChatIDBotID(chatID, botID)
// 	if err != nil {
// 		b.SendMessage(ctx, &bot.SendMessageParams{
// 			ChatID: update.Message.Chat.ID,
// 			Text:   err.Error(),
// 		})
// 	} else {
// 		if update.Message.Location == nil {
// 			kb := &tgmodels.ReplyKeyboardMarkup{
// 				Keyboard: [][]tgmodels.KeyboardButton{
// 					{
// 						{Text: "Lokatsiya", RequestLocation: true},
// 					},
// 				},
// 				ResizeKeyboard:  true,
// 				OneTimeKeyboard: true,
// 			}
// 			b.SendMessage(ctx, &bot.SendMessageParams{
// 				ChatID:      update.Message.Chat.ID,
// 				Text:        "Iltimos lokatsiya kiriting!",
// 				ReplyMarkup: kb,
// 			})
// 		} else {
// 			_, err := h.Stg.Client().Update(models.UpdateClientRequest{
// 				ID:        session.OrderID,
// 				Latitute:  update.Message.Location.Latitude,
// 				Longitude: update.Message.Location.Longitude,
// 			})
// 			if err != nil {
// 				b.SendMessage(ctx, &bot.SendMessageParams{
// 					ChatID: update.Message.Chat.ID,
// 					Text:   err.Error(),
// 				})
// 			} else {
// 				dialogStep := ""
// 				_, err := h.Stg.BotUser().Update(models.BotUser{
// 					UserID:     user.UserID,
// 					BotID:      int(botID),
// 					ChatID:     update.Message.Chat.ID,
// 					DialogStep: &dialogStep,
// 				})
// 				if err != nil {
// 					b.SendMessage(ctx, &bot.SendMessageParams{
// 						ChatID: update.Message.Chat.ID,
// 						Text:   err.Error(),
// 					})
// 				} else {
// 					h.Stg.TelegramSession().Delete(session.ID)
// 					b.SendMessage(ctx, &bot.SendMessageParams{
// 						ChatID: update.Message.Chat.ID,
// 						Text:   "Klient lokatsiyasi belgilandi!",
// 					})
// 				}
// 			}
// 		}
// 	}
// }

func (h *Handler) handleError(ctx context.Context, b *bot.Bot, update *tgmodels.Update, errorMessage string, replyToMessageID int) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        errorMessage,
		ReplyMarkup: replyToMessageID,
	})
}

func (h *Handler) SendLocation(c *gin.Context) {
	var query models.OrderSendLocationRequest
	if err := c.ShouldBindQuery(&query); err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	user, _ := utils.ExtractTokenID(c)
	order, err := h.Stg.Order().GetLocation(query.OrderID)
	if err != nil {
		h.handleResponse(c, http.OK, err.Error())
		return
	}

	botUser, _ := h.Stg.BotUser().GetByUserID(user.UserID)

	if order.Latitute == nil || order.Longitude == nil {
		h.handleResponse(c, http.OK, "Bu buyurtma lokatsiyasi mavjud emas!")
		return
	}

	opts := []bot.Option{
		bot.WithDefaultHandler(h.Handler),
	}

	b, err := bot.New(botUser.BotToken, opts...)
	if err != nil {
		h.handleResponse(c, http.OK, err.Error())
		return
	}

	b.SendLocation(c, &bot.SendLocationParams{
		ChatID:    botUser.ChatID,
		Latitude:  *order.Latitute,
		Longitude: *order.Longitude,
	})
	b.SendMessage(c, &bot.SendMessageParams{
		ChatID: botUser.ChatID,
		Text:   "Buyurtma birkasi: " + order.Slug + "\nBuyurtmachi telefon raqami: " + order.PhoneNumber,
	})

	h.handleResponse(c, http.OK, "Lokatsiya jo'natildi!")
}
