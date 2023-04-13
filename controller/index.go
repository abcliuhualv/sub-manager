package controller

import (
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	u, _ := c.Get("username")
	username, _ := u.(string)
	filePath := filepath.Join("./files/", username, "/origin/")
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		log.Fatal(err)
	}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"files": files,
	})
}
