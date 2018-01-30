package atfile

import (
	"os"
	"log"
	"path/filepath"
	"strings"
)

func CreateFile(logFilePath, fileName string) (*os.File, error) {
	file, err := os.Create(logFilePath + fileName)
	if err != nil {
		log.Fatal("create log file failed", err)
		return nil, err
	}
	return file, nil

}

/**
获取当前的相对路径
*/
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(""))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
