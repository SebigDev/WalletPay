package middlewares

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
)

type OperationContext struct{}

func NewOperationContext() OperationContext {
	return OperationContext{}
}

func (*OperationContext) SetXRequestIDContext(app *fiber.App) {
	app.Use(requestid.New(requestid.Config{
		Next:       nil,
		Header:     fiber.HeaderXRequestID,
		Generator:  utils.UUID,
		ContextKey: "requestid",
	}))
}

func (*OperationContext) LivenessAndHealthCheck(app *fiber.App) {
	app.Use(healthcheck.New(healthcheck.Config{
		LivenessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		LivenessEndpoint: "/live",
		ReadinessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		ReadinessEndpoint: "/ready",
	}))
}

func HandleTimeOutFunc(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 100*time.Millisecond)
	defer cancel()

	for {
		<-c.Done()
		return fiber.NewError(fiber.StatusRequestTimeout, "operation timed out")
	}
}
