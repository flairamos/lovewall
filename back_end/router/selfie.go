package router

import (
	"shoollove/handler"

	"github.com/gofiber/fiber/v2"
)

// selfie 路由，记得把 selfie(app) 添加把 router.go 中
func Selfie(app fiber.Router, group string) {

	selfieRouter := app.Group(group)
	{
		selfieRouter.Post("/add", handler.SelfieAdd)
		// selfieRouter.Post("/edit", handler.SelfieEdit)
		// selfieRouter.Post("/delete", handler.SelfieDelete)
		selfieRouter.Get("/get", handler.SelfieGet)
		selfieRouter.Get("/praise", handler.SelfiePraise)
		// selfieRouter.Get("/form", handler.SelfieForm)
	}
}
