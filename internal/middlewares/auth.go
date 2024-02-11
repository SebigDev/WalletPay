package middlewares

import (
	"CrashCourse/GoApp/internal/utils"
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	ctx context.Context
}

func NewAuthMiddleware(ctx context.Context) *AuthMiddleware {
	return &AuthMiddleware{
		ctx: ctx,
	}
}

func (mw *AuthMiddleware) UserAuthMiddlewareHandler(c *fiber.Ctx) error {

	jwtToken, err := utils.GetToken(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	if jwtToken.Valid {
		if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok {
			expirationFloat := claims["exp"].(float64)

			if !ok {
				log.Println("Expiration not found in claims")
				return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
			}

			expiration := time.Unix(int64(expirationFloat), 0)

			if time.Now().UTC().After(expiration) {
				log.Println("Token has expired")
				return fiber.NewError(fiber.StatusUnauthorized, "Token has expired")
			}
		} else {
			log.Println("Invalid claims")
			return err
		}
	} else {
		log.Println("Token invalid")
	}
	return c.Next()
}
