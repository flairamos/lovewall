package respond

import "github.com/gofiber/fiber/v2"

func Fail(c *fiber.Ctx, msg string, err error) error {
	return c.JSON(fiber.Map{
		"code":  false,
		"msg":   msg,
		"error": err.Error(),
	})
}

func Ok(c *fiber.Ctx, data any, msg ...string) error {
	var message string
	if len(msg) == 0 {
		message = "success"
	} else {
		message = ""
		for _, s := range msg {
			message += s
		}
	}
	return c.JSON(fiber.Map{
		"code": true,
		"msg":  message,
		"data": data,
	})
}
