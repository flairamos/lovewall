package router

import (
	"github.com/gofiber/fiber/v2"
	"shoollove/handler"
)

// goods 路由，记得把 goods(app) 添加把 router.go 中
func Goods(app fiber.Router, group string) {

	goodsRouter := app.Group(group)
	{
		goodsRouter.Post("/add", handler.GoodsAdd)
		goodsRouter.Post("/edit", handler.GoodsEdit)
		goodsRouter.Post("/delete", handler.GoodsDelete)
		goodsRouter.Get("/get", handler.GoodsGet)
		goodsRouter.Get("/form", handler.GoodsForm)
	}
}
