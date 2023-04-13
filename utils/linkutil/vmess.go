package linkutil

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

func ParseVmessLink(link string) (map[string]interface{}, error) {
	vmessData, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(link, "vmess://"))
	if err != nil {
		return nil, err
	}

	var paramMap map[string]interface{}
	err = json.Unmarshal(vmessData, &paramMap)
	if err != nil {
		return nil, err
	}
	return paramMap, nil
}

func BuildVmessLink(paramMap map[string]interface{}) (string, error) {
	vmessData, err := json.Marshal(&paramMap)
	if err != nil {
		// 编码错误
		return "", err
	}

	vmessLink := "vmess://" + base64.StdEncoding.EncodeToString(vmessData)
	return vmessLink, nil
}
