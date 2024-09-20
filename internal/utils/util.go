package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/sebigdev/walletpay/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateSecureToken() string {
	b := make([]byte, 24)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func GetToken(c *fiber.Ctx) (*jwt.Token, error) {
	authorization := c.Get("Authorization")
	if !strings.Contains(authorization, "Bearer") {
		return nil, fmt.Errorf("unrecognized token")
	}

	arrayToken := strings.Split(authorization, " ")
	if len(arrayToken) != 2 {
		return nil, fmt.Errorf("unrecognized token")
	}

	encodedToken := arrayToken[1]
	jwtToken, err := jwt.Parse(
		encodedToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid signing method")
			}
			return []byte(config.GoEnv("SECRET_KEY")), nil
		})
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("oops!!! something went wrong")
	}

	return jwtToken, nil
}

func GetUserIdFromToken(ctx *fiber.Ctx) (string, error) {
	token, err := GetToken(ctx)
	if err != nil {
		log.Fatal(err)
	}

	userId, err := token.Claims.(jwt.MapClaims).GetSubject()

	if err != nil {
		log.Fatal(err)
	}

	return userId, nil
}

func Length(p string) int {
	return len(strings.TrimSpace(p))
}

func ToStringArray(value string) []string {
	return []string{value}
}
