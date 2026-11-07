package handler

import (
	"errors"
	"fmt"
	"io"
	"shoollove/config"
	"shoollove/pkg/filebrowser"
	"shoollove/pkg/respond"
	"shoollove/pkg/tool"

	"github.com/gofiber/fiber/v2"
)

func UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return respond.Fail(c, "读取文件失败", err)
	}
	open, err := file.Open()
	if err != nil {
		return respond.Fail(c, "打开文件失败", err)
	}
	defer open.Close()
	fileBytes, err := io.ReadAll(open)
	if err != nil {
		return respond.Fail(c, "文件转字节失败", err)
	}
	// 获取文件类型
	fileType := c.Get("file_type")
	fileName := tool.Uuid() + "." + fileType
	str, err := filebrowser.UploadFileStream(filebrowser.FileBrowserConfig{
		BaseURL:  config.Config.Filebrowser.Url,
		Username: config.Config.Filebrowser.Username,
		Password: config.Config.Filebrowser.Password,
		FilePath: fmt.Sprintf("/%s/%s", config.Config.Filebrowser.Bucket, fileName),
		Override: false,
	}, fileBytes)
	if err != nil {
		return respond.Fail(c, "fail", errors.New(str))
	}
	c.Set("Content-Type", "application/json")
	return respond.Ok(c, map[string]string{"fileName": fileName}, "success")
}

func DeleteFile (c *fiber.Ctx) error {
	name := c.Params("name")
	if err := filebrowser.DeleteFile(fmt.Sprintf("%s", name)); err != nil {
		return respond.Fail(c,"删除失败",err)
	}
	return respond.Ok(c,"删除成功")
}
