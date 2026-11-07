package main

import (
	"fmt"
	"os"
	"shoollove/config"
	"shoollove/pkg/db"
	redisutil "shoollove/pkg/redis_util"
	"shoollove/router"
)

func main() {
	// 获取文件地址
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("----获取文件当前目录失败---", err.Error())
		return
	}
	// 初始化配置
	config.Init(path)

	// 初始化minio
	// minioutil.Init()

	// 初始化数据库
	redisutil.Init()

	// 初始化数据库客户端
	db.Init()

	// 初始话路由
	router.Init()
}
