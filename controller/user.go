package controller

import (
	"fmt"
	"net/http"
	"os"
	"sub-manager/dao"
	"sub-manager/initialize"
	"sub-manager/middleware"
	"sub-manager/model"

	"github.com/gin-gonic/gin"
)

func GoLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", nil)
}

func Login(c *gin.Context) {
	username, ok := c.GetPostForm("username")
	if !ok || username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请输入用户名"})
		return
	}
	password, ok := c.GetPostForm("password")
	if !ok || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请输入密码"})
		return
	}

	user, err := dao.UserDao.FindByName(username)
	if err != nil || user == nil || password != user.Password {
		fmt.Printf("账号或密码错误")
		c.HTML(http.StatusBadRequest, "login.tmpl", gin.H{"error": "账号或密码错误"})
		return
	}

	tokenString, err := middleware.CreateToken(user.Username)
	if err != nil {
		fmt.Printf("Login err1: %v\n", err)
		c.HTML(http.StatusBadRequest, "login.tmpl", gin.H{"error": "token创建失败,请查看运行日志并检查代码"})
		return
	}

	c.SetCookie("token", tokenString, middleware.TokenMaxAge, "/", c.GetHeader("Host"), false, true)
	c.Redirect(http.StatusFound, "/")
}

func LogOut(c *gin.Context) {
	c.SetCookie("token", "", 0, "/", c.GetHeader("Host"), false, true)
	c.HTML(http.StatusOK, "logout_success.tmpl", nil)
}

func GoRegist(c *gin.Context) {
	if !initialize.OpenRegist {
		c.HTML(http.StatusBadRequest, "regist.tmpl", gin.H{"error": "注册通道已关闭"})
		return
	}
	c.HTML(http.StatusOK, "regist.tmpl", nil)
}

func Regist(c *gin.Context) {
	if !initialize.OpenRegist {
		c.HTML(http.StatusBadRequest, "regist.tmpl", gin.H{"error": "注册通道已关闭"})
		return
	}

	username, ok := c.GetPostForm("username")
	if !ok || username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请输入用户名"})
		return
	}
	password, ok := c.GetPostForm("password")
	if !ok || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请输入密码"})
		return
	}
	user := model.User{
		Username: username,
		Password: password,
	}
	err := dao.UserDao.Add(&user)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		c.HTML(http.StatusBadRequest, "regist.tmpl", gin.H{"error": "注册失败,请查看运行日志并检查代码"})
		return
	}

	err = os.MkdirAll("./files/"+username+"/origin", 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.MkdirAll("./files/"+username+"/sub", 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.HTML(http.StatusOK, "login.tmpl", gin.H{"error": "注册成功,请登录"})
}
