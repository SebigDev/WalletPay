package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("SecretKey")

func GenerateSecureToken() string {
	b := make([]byte, 24)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func GetToken(c *fiber.Ctx) (*jwt.Token, error) {
	authorization := c.Get("Authorization")
	jwtToken, err := jwt.Parse(
		authorization, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid signing method")
			}
			return SecretKey, nil
		})
	if err != nil {
		log.Println("Token parsing error")
		return nil, err
	}

	return jwtToken, nil
}

func GetUserIdFromToken(ctx *fiber.Ctx) (string, error) {
	token, err := GetToken(ctx)
	if err != nil {
		log.Fatal(err)
	}
	userId := ""
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userId = claims["sub"].(string)

		if !ok {
			log.Fatal("Error getting claim")
		}
	}
	return userId, nil
}
