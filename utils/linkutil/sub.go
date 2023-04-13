package linkutil

import (
	"encoding/base64"
	"fmt"
	"path/filepath"
	"strings"
	"sub-manager/initialize"
	"sub-manager/utils/csvutil"
	"sub-manager/utils/fileutil"
)

func LinksToSub(links []string) string {
	subLink := base64.StdEncoding.EncodeToString([]byte(strings.Join(links, "\r\n")))

	return subLink
}

func CreateSubToFile(record []string, userDir string) {
	remarksPref := record[0]
	subDir := filepath.Join(userDir, "sub")
	originDir := filepath.Join(userDir, "origin")
	outFile := filepath.Join(subDir, record[1])
	ipFile := filepath.Join(originDir, record[2])
	originLink := record[3]

	fileutil.TestAndCreateDir(subDir)
	fileutil.TestAndCreateDir(originDir)

	var links []string
	reverseProxies, err := csvutil.FileToSlice(ipFile)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	if strings.HasPrefix(originLink, "vmess://") {
		paramMap, _ := ParseVmessLink(originLink)
		for _, reverseProxy := range reverseProxies {
			paramMap["add"] = reverseProxy[0]
			paramMap["port"] = reverseProxy[1]
			paramMap["ps"] = remarksPref + "+" + reverseProxy[2]
			newLink, _ := BuildVmessLink(paramMap)
			links = append(links, newLink)
		}
	} else if strings.HasPrefix(originLink, "vless://") {
		paramMap, _ := ParseVlessLink(originLink)
		for _, reverseProxy := range reverseProxies {
			paramMap["address"] = reverseProxy[0]
			paramMap["port"] = reverseProxy[1]
			paramMap["remark"] = remarksPref + "+" + reverseProxy[2]
			newLink, _ := BuildVlessLink(paramMap)
			links = append(links, newLink)
		}
	}

	subLink := LinksToSub(links)
	fileutil.WriteStringToFile(subLink, outFile)
}

func CreateSubAndUpload(userDir string) {
	subDir := filepath.Join(userDir, "sub")
	originDir := filepath.Join(userDir, "origin")
	fileutil.TestAndCreateDir(subDir)
	fileutil.TestAndCreateDir(originDir)

	recordsFile := filepath.Join(originDir, "map.csv")
	records, err := csvutil.FileToSlice(recordsFile)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	for _, record := range records {
		CreateSubToFile(record, userDir)
	}

	username := filepath.Base(userDir)
	fileutil.Upload(subDir, filepath.Join(initialize.RcloneRemotePath, username))
}
