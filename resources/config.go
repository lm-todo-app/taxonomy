package resources

import "github.com/gofiber/fiber/v2"

func InitConfig() fiber.Config {
	return fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			return ctx.Status(code).JSON(fiber.Map{
				"code":    code,
				"message": err.Error(),
			})
		},
	}
}
