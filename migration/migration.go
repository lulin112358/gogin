package migration

import "gogin/model"

// 自动迁移
func Migration() {
	model.Db.AutoMigrate(&model.Users{})
}
