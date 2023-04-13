package csvutil

import (
	"encoding/csv"
	"fmt"
	"os"
)

func FileToSlice(filepath string) ([][]string, error) {
	// 打开文件
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer file.Close()

	// 从文件创建CSV读取器
	reader := csv.NewReader(file)

	// 读取所有记录
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return records, nil
}
