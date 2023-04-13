package initialize

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"sub-manager/middleware"

	"gopkg.in/ini.v1"
)

var (
	RcloneRemotePath string
	OpenRegist       bool
	ListenPort       string
)

/*
jwt配置
数据库配置
是否开启注册的配置
rclone配置
*/
func ReadOrCreateConf() {
	file, err := ini.Load("./config.ini")
	if err != nil {
		fmt.Println("配置文件不存在,创建默认配置文件")
		b := make([]byte, 15)
		_, err := rand.Read(b)
		if err != nil {
			log.Fatal(err)
		}
		randStr := base64.URLEncoding.EncodeToString(b)
		fmt.Printf("randStr: %v\n", randStr)
		cfg := ini.Empty()
		cfg.NewSections("listen", "jwt", "database", "regist", "rclone")
		cfg.Section("jwt").NewKey("JwtKey", randStr)
		cfg.Section("jwt").NewKey("TokenMaxAge", "7200")
		cfg.Section("database").NewKey("type", "sqlite")
		cfg.Section("database").NewKey("addr", "./sub-manager.db")
		cfg.Section("regist").NewKey("openRegist", "false")
		cfg.Section("rclone").NewKey("remotePath", "rclone_onedrive:cloudreve/1/sub")
		cfg.Section("listen").NewKey("port", "16666")

		err = cfg.SaveTo("./config.ini")
		if err != nil {
			log.Fatal(err)
		}
		file = cfg
	}
	middleware.JwtKey = []byte(file.Section("jwt").Key("JwtKey").String())
	middleware.TokenMaxAge, _ = file.Section("jwt").Key("TokenMaxAge").Int()
	databasePath = file.Section("database").Key("addr").String()
	OpenRegist, _ = file.Section("regist").Key("openRegist").Bool()
	RcloneRemotePath = file.Section("rclone").Key("remotePath").String()
	ListenPort = file.Section("listen").Key("port").String()
}
