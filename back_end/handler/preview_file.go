package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"shoollove/config"
)

func ProxyToFileBrowser(c *fiber.Ctx) error {
	name := c.Params("name")
	shareID := config.Config.Filebrowser.ShareID
	if shareID == "" {
		return fiber.NewError(fiber.StatusInternalServerError, "filebrowser.share_id is not configured")
	}
	url := config.Config.Proxy.Url + "/api/public/dl/" + shareID + "/" + name + "?inline=true"
	err := proxy.Do(c, url)
	if err != nil {
		return err
	}
	return nil
}
