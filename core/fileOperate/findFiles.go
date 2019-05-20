package fileOperate

import (
	"fmt"
	"os"
	"path/filepath"
)

// 获取给定目录下的文件
func GetFiles(pathName string) []string {
	var fileArray []string
	err := filepath.Walk(pathName, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if err != nil {
			return err
		}
		fileArray = append(fileArray, path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	return convertPathRel2Abs(fileArray)
}


// 获取指定的正则匹配模式
func GetFilesByRegex(pattern string) []string {
	fileArray, err := filepath.Glob(pattern)
	if err != nil {
		panic(err)
	}
	return convertPathRel2Abs(fileArray)
}


// 将列表里的目录由相对目录转为绝对目录
func convertPathRel2Abs(relArray []string) []string {
	var absArray []string
	for _, each := range relArray {
		absPath, err := filepath.Abs(each)
		if err != nil {
			panic(err)
		}
		absArray = append(absArray, absPath)
	}
	return absArray
}