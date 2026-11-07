package router

import (
	"github.com/gofiber/fiber/v2"
	"shoollove/handler"
)

func Follow(app fiber.Router, group string) {

	followRouter := app.Group(group)
	{
		followRouter.Post("/add", handler.FollowAdd)
		followRouter.Post("/delete", handler.FollowDelete)
		followRouter.Get("/get", handler.FollowGet)
		followRouter.Get("/info", handler.FollowInfo)
		followRouter.Get("/follow_list", handler.FindFollowList)
	}
}
