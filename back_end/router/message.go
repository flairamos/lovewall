package router

import (
	"github.com/gofiber/fiber/v2"
	"shoollove/handler"
)

func Message(app fiber.Router, group string) {
	baseRouter := app.Group(group)
	{
		baseRouter.Get("/unread_message", handler.GetUnReadMessage)
		baseRouter.Get("/get_message", handler.GetMessage)
		baseRouter.Post("/send_message", handler.SendMessage)
		baseRouter.Post("/clear_unread_message", handler.ClearUnreadMessage)
	}
}
