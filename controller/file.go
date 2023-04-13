package controller

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func GoNewFile(c *gin.Context) {
	c.HTML(http.StatusOK, "new.tmpl", nil)
}

func NewFile(c *gin.Context) {
	filename := c.PostForm("filename")
	content := c.PostForm("content")
	u, _ := c.Get("username")
	username, _ := u.(string)
	filePath := filepath.Join("./files/", username, "/origin/", filename)
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		c.HTML(http.StatusBadRequest, "new.tmpl", gin.H{"error": "文件已存在"})
		return
	}
	err := ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}
	c.Redirect(http.StatusFound, "/")
}

func GoEditFile(c *gin.Context) {
	filename := c.Query("filename")
	u, _ := c.Get("username")
	username, _ := u.(string)
	filePath := filepath.Join("./files/", username, "/origin/", filename)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	c.HTML(http.StatusOK, "edit.tmpl", gin.H{
		"filename": filename,
		"content":  string(content),
	})
}

func EditFile(c *gin.Context) {
	filename := c.PostForm("filename")
	content := c.PostForm("content")
	u, _ := c.Get("username")
	username, _ := u.(string)
	filePath := filepath.Join("./files/", username, "/origin/", filename)
	err := ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}
	c.Redirect(http.StatusFound, "/")
}

func DeleteFile(c *gin.Context) {
	filename := c.Query("filename")
	u, _ := c.Get("username")
	username, _ := u.(string)
	filePath := filepath.Join("./files/", username, "/origin/", filename)
	err := os.Remove(filePath)
	if err != nil {
		log.Fatal(err)
	}
	c.Redirect(http.StatusFound, "/")
}
