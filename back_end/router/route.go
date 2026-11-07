package router

import (
	"github.com/gofiber/contrib/websocket"
	"shoollove/config"
	"shoollove/handler"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	base := app.Group("/api/wechat")
	// public
	Public(base, "public")
	Account(base, "user")

	Topic(base, "topic")
	Posting(base, "posting")
	//app.Use(JwtAuth)
	Selfie(base, "ablum")
	Goods(base, "goods")
	Follow(base, "follow")
	Message(base, "message")

	HandleFile(base, "images")

	resource := app.Group("/resource")
	Resource(resource, "/")

	// websocket
	app.Get("/ws", websocket.New(handler.MessageSocket))
	app.Get("/push", websocket.New(handler.MessagePushNew))
}

func Init() {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true, // 打印路由信息
	})
	Route(app)
	err := app.Listen(":" + strconv.Itoa(config.Config.Server.Port))
	if err != nil {
		panic(err)
	}
}
