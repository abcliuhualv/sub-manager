package main

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"sub-manager/controller"
	"sub-manager/initialize"
	"sub-manager/middleware"

	"github.com/gin-gonic/gin"
)

var (
	//go:embed assets/*
	assetsFS embed.FS

	//go:embed templates
	templatesFS embed.FS
)

// sub 函数返回一个 FS，该 FS 只包含根目录下的子目录和文件
func sub(fsys fs.FS, name string) fs.FS {
	subfs, _ := fs.Sub(fsys, name)
	return subfs
}

func init() {
	initialize.ReadOrCreateConf()
	initialize.DbInit()
}

func main() {
	r := gin.Default()

	// 初始化默认静态资源
	r.StaticFS("/assets", http.FS(sub(assetsFS, "assets")))

	// 设置模板资源
	r.SetHTMLTemplate(template.Must(template.New("").ParseFS(templatesFS, "templates/*")))

	// // 初始化默认静态资源
	// r.Static("/assets", "./assets")

	// // 设置模板资源
	// r.LoadHTMLGlob("templates/*.tmpl")

	// 登录页面
	r.GET("/login", controller.GoLogin)

	// 登录功能
	r.POST("/login", controller.Login)

	// 退出登录
	r.GET("/logout", controller.LogOut)

	// 注册页面
	r.GET("/regist", controller.GoRegist)

	// 注册功能
	r.POST("/regist", controller.Regist)

	authMiddleware := middleware.AuthMiddleWare()
	// 主页
	r.GET("/", authMiddleware, controller.Index)

	// 新建文件
	r.GET("/new", authMiddleware, controller.GoNewFile)

	// 保存新文件
	r.POST("/new", authMiddleware, controller.NewFile)

	// 编辑文件
	r.GET("/edit", authMiddleware, controller.GoEditFile)

	// 保存编辑的文件
	r.POST("/edit", authMiddleware, controller.EditFile)

	// 删除文件
	r.GET("/delete", authMiddleware, controller.DeleteFile)

	//生成订阅并同步至cloudreve
	r.GET("/createsub", authMiddleware, controller.CreateSub)

	// 启动服务器
	r.Run(":" + initialize.ListenPort)
	// if err := r.Run(":8080"); err != nil {
	// 	log.Fatal(err)
	// }
}
