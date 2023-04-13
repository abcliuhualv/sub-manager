package controller

import (
	"net/http"
	"path/filepath"
	"sub-manager/utils/linkutil"

	"github.com/gin-gonic/gin"
)

func CreateSub(c *gin.Context) {
	u, _ := c.Get("username")
	username, _ := u.(string)
	userDir := filepath.Join("./files/", username)

	linkutil.CreateSubAndUpload(userDir)
	c.HTML(http.StatusOK, "createsub_success.tmpl", gin.H{
		"msg": "订阅生成并同步成功,3s后自动跳转至主页面...",
	})
}
