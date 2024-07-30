package utils

import (
	"bw-erp/models"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(user_id string, phone string) (string, string, error) {
	accessToken, err := createToken(user_id, phone, 1) // 1 hour lifespan for access token
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
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["phone"] = phone
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(lifespan)).Unix()
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
		jwtdata.UserID, _ = claims["user_id"].(string)
		return jwtdata, nil
	}

	return jwtdata, err
}
