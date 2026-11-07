package router

import (
	"github.com/gofiber/fiber/v2"
	"shoollove/handler"
)

// topic 路由，记得把 topic(app) 添加把 router.go 中
func Topic(app fiber.Router, group string) {
	topicRouter := app.Group(group)
	{
		topicRouter.Post("/add", handler.TopicAdd)
		topicRouter.Post("/edit", handler.TopicEdit)
		topicRouter.Post("/delete", handler.TopicDelete)
		topicRouter.Get("/get", handler.TopicGet)
		topicRouter.Get("/form", handler.TopicForm)
		topicRouter.Get("/praise", handler.TopicPraise)
	}
}
