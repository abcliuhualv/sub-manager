package initialize

import (
	"fmt"
	"sub-manager/dao"
	"sub-manager/model"

	"github.com/glebarez/sqlite" // 纯 Go 实现的 SQLite 驱动, 详情参考： https://github.com/glebarez/sqlite
	"gorm.io/gorm"
)

var (
	databasePath string
)

func DbInit() {
	db, err := gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.User{})

	dao.UserDao = dao.NewUserDAO(db)
}
