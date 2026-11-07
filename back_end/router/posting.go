package router

import (
	"shoollove/handler"

	"github.com/gofiber/fiber/v2"
)

// posting 路由，记得把 posting(app) 添加把 router.go 中
func Posting(app fiber.Router, group string) {
	postingRouter := app.Group(group)
	{
		postingRouter.Post("/add", handler.PostingAdd)
		postingRouter.Post("/edit", handler.PostingEdit)
		postingRouter.Post("/delete", handler.PostingDelete)
		postingRouter.Get("/get", handler.PostingGet)
		postingRouter.Get("/form", handler.PostingForm)
		postingRouter.Post("/praise", handler.PostingPraise)
		postingRouter.Get("/get_comments", handler.GetCommentsById)
	}
}
