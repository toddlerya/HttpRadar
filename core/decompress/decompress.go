package decompress

import (
	"bytes"
	"fmt"
	"github.com/gpmgo/gopm/modules/cae/zip"
	"github.com/toddlerya/HttpRadar/core/fileOperate"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// 判断是否为zip文件
func isZip(zipFilePath string) bool {
	f, err := os.Open(zipFilePath)
	if err != nil {
		return false
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	buf := make([]byte, 4)
	if n, err := f.Read(buf); err != nil || n < 4 {
		return false
	}
	return bytes.Equal(buf, []byte("PK\x03\x04"))
}

// 解压文件
func unZip(zipFile, dest string) error {
	reader, err := zip.Open(zipFile)
	if err != nil {
		return err
	}
	defer func() {
		if err := reader.Close(); err != nil {
			panic(err)
		}
	}()

	fileOperate.CreateDirIfNotExist(dest, 0755)

	for _, file := range reader.File {
		destPath := filepath.Join(dest, file.Name)
		if file.FileInfo().IsDir() {
			fileOperate.CreateDirIfNotExist(destPath, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := fileReader.Close(); err != nil {
				panic(err)
			}
		}()

		targetFile, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}

		defer func() {
			if err := targetFile.Close(); err != nil {
				panic(err)
			}
		}()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}

// 处理saz文件
func Saz(sazFilePath string) {
	// 创建临时目录
	sazFileName := strings.Split(path.Base(sazFilePath), ".saz")[0]
	tempDirName := fmt.Sprintf("./temp/%s", sazFileName)
	fileOperate.CreateDirIfNotExist(tempDirName, 0755)
	tempZipPath := fmt.Sprintf("%s.zip", tempDirName)
	_, err := fileOperate.Copy(sazFilePath, tempZipPath)
	if isZip(tempZipPath) {
		err = unZip(tempZipPath, tempDirName)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Printf("请检查saz文件是否正确: %s", sazFileName)
	}
}
