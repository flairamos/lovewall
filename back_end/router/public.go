package router

import (
	"shoollove/handler"

	"github.com/gofiber/fiber/v2"
)

func Public(app fiber.Router, group string) {
	publicRouter := app.Group(group)
	{
		publicRouter.Post("/login", handler.Login)
		publicRouter.Post("/register", handler.AccountAdd)
		publicRouter.Post("/send_code", handler.SendRegisterCode)
		publicRouter.Post("/check_code", handler.CheckRegisterCode)
		publicRouter.Post("/check_login", handler.CheckLogin)
	}
}

func Resource(app fiber.Router, group string) {
	resourceRouter := app.Group(group)
	{
		resourceRouter.Get("/img/:name", handler.ProxyToFileBrowser)
	}
}
