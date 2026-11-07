package minioutil

import (
	"bytes"
	"io"
	"log"
	"shoollove/config"
	"shoollove/pkg/tool"

	"github.com/minio/minio-go"
)

// filetype是文件的后缀
func Upload(by []byte, filetype string, buckets ...string) (string, error) {
	filename := tool.Uuid() + "." + filetype
	var bucket string
	if len(bucket) == 0 {
		bucket = config.Config.Minio.Bucket
	} else {
		bucket = buckets[0]
	}
	_, err := Client.PutObject(bucket, filename, bytes.NewReader(by), int64(len(by)), minio.PutObjectOptions{})
	if err != nil {
		log.Println("minio上传文件失败", err.Error())
		return "", err
	}
	return filename, nil
}

func Download(filename string, buckets ...string) ([]byte, error) {
	var bucket string
	if len(bucket) == 0 {
		bucket = config.Config.Minio.Bucket
	} else {
		bucket = buckets[0]
	}
	obj, err := Client.GetObject(bucket, filename, minio.GetObjectOptions{})
	if err != nil {
		log.Println("minio获取文件失败", err.Error())
		return nil, err
	}
	by, err := io.ReadAll(obj)
	if err != nil {
		log.Println("读取文件流失败", err.Error())
	}
	return by, nil
}
