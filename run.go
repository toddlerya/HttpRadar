package main

import (
	"fmt"
	"github.com/toddlerya/HttpRadar/core/fileOperate"
)

const sazFilesPath = "./sazFiles/*.saz"

func main() {
	fileArray := fileOperate.GetFilesByRegex(sazFilesPath)
	for _, each := range fileArray {
		fmt.Println(each)
	}
}
