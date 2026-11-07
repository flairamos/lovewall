package minioutil

import (
	"log"
	"shoollove/config"

	"github.com/minio/minio-go"
)

var (
	Client *minio.Client
)

func Init() {
	conf := config.Config.Minio
	client, err := minio.New(conf.Endpoint, conf.AccessKey, conf.AccessSecret, conf.UseSSL)
	if err != nil {
		log.Println("minio client create err: ", err)
		panic(err)
	}
	// 判断认证信息是否有误
	_, err = client.ListBuckets()
	if err != nil {
		log.Println("minio client list buckets err: ", err)
		panic(err)
	}
	log.Println("minio client created")
	Client = client
	return
}

func MakeBucket(bucket string) error {
	exists, err := Client.BucketExists(bucket)
	if err != nil {
		log.Println("桶验证错误", err.Error())
		return err
	}
	if !exists {
		err := Client.MakeBucket(bucket, "us-east-1")
		if err != nil {
			return err
		}
	}
	return nil
}
