package fileutil

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func TestAndCreateDir(dirPath string) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// 如果不存在，创建文件夹
		err = os.MkdirAll(dirPath, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func WriteStringToFile(content string, outFile string) {
	err := ioutil.WriteFile(outFile, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("File written successfully")
}

func Upload(sourcePath string, destPath string) {
	cmd := exec.Command("rclone", "sync", sourcePath, destPath)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("上传出错:", err)
		return
	}
}
