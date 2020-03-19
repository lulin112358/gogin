package main

import (
	"gogin/migration"
	"gogin/model"
	"gogin/router"
)

func main() {
	// 连接数据库
	model.Init()

	// 开启数据迁移
	migration.Migration()

	// 加载路由
	r := router.Init()

	// 运行程序
	_ = r.Run(":9090")
}