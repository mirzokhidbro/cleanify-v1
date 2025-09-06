package utils

import (
	"bw-erp/models"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func GenerateToken(user_id string, phone string) (string, string, error) {
	accessToken, err := createToken(user_id, phone, 6) // 1 hour lifespan for access token
	if err != nil {
		return "", "", err
	}

	refreshToken, err := createToken(user_id, phone, 720) // 720 hours (30 days) lifespan for refresh token
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func createToken(user_id string, phone string, lifespan int) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"iss":   "bw-erp",                                            // issuer
		"iat":   now.Unix(),                                          // issued at
		"exp":   now.Add(time.Hour * time.Duration(lifespan)).Unix(), // expiration
		"nbf":   now.Unix(),                                          // not before
		"sub":   user_id,                                             // subject (user ID)
		"jti":   uuid.New().String(),                                 // JWT ID
		"prv":   []string{"*"},                                       // Laravel JWT specific
		"phone": phone,                                               // custom claim
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenID(c *gin.Context) (models.JWTData, error) {
	var jwtdata models.JWTData
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return jwtdata, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		jwtdata.Phone, _ = claims["phone"].(string)
		user_id, _ := claims["sub"].(string)

		userId, err := strconv.Atoi(user_id)
		if err != nil {
			return jwtdata, err
		}

		jwtdata.UserID = int64(userId)
		return jwtdata, nil
	}

	return jwtdata, err
}
