package utils

import (
	"os"

	"github.com/json-iterator/go"

	"github.com/qiniu/log"
)

// IsExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// 获取测试数据
func GetParseTestData(line string, size int) []string {
	testSlice := make([]string, 0)
	totalSize := 0
	for {
		if totalSize > size {
			return testSlice
		}
		testSlice = append(testSlice, line)
		totalSize += len(line)
	}
}

// 深拷贝
func DeepCopyByJSON(dst, src interface{}) {
	confBytes, err := jsoniter.Marshal(src)
	if err != nil {
		log.Errorf("deepCopyByJSON marshal error %v, use same pointer", err)
		dst = src
		return
	}
	if err = jsoniter.Unmarshal(confBytes, dst); err != nil {
		log.Errorf("deepCopyByJSON unmarshal error %v, use same pointer", err)
		dst = src
		return
	}
}
