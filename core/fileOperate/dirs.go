package fileOperate

import "os"

// 判断目录是否存在

// 创建目录
func CreateDirIfNotExist(dirName string, permMode os.FileMode) {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err = os.MkdirAll(dirName, permMode)
		if err != nil {
			panic(err)
		}
	}
}
