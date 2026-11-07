package router

import (
	"shoollove/handler"

	"github.com/gofiber/fiber/v2"
)

// account 路由，记得把 account(app) 添加把 router.go 中
func Account(app fiber.Router, group string) {

	accountRouter := app.Group(group)
	{
		// accountRouter.Post("/add", handler.AccountAdd)
		accountRouter.Post("/edit", handler.AccountEdit)
		accountRouter.Post("/delete", handler.AccountDelete)
		accountRouter.Get("/get", handler.AccountGet)
		accountRouter.Get("/form", handler.AccountForm)
		accountRouter.Get("/getbyid", handler.AccountById)
	}
}
