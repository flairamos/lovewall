package router

import (
	"shoollove/handler"

	"github.com/gofiber/fiber/v2"
)

func HandleFile(app fiber.Router, group string) {
	baseRouter := app.Group(group)
	{
		baseRouter.Post("upload", handler.UploadFile)
		baseRouter.Get("delete/:name",handler.DeleteFile)
	}
}
