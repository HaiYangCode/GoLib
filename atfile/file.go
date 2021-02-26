package atfile

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
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
获取当前的绝对路径
*/
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(""))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

/**
获取当前的相对路径
*/
func GetCurrentRelDir() string {
	dir, err := filepath.Rel("./", GetCurrentDirectory())
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

var ostype = runtime.GOOS
/**
	读取配置
 */

func ReadConfig(currentPath, fileName string) (data []byte, err error) {

	fmt.Println("操作系统类型 ", ostype)
	if "windows" == ostype {
		data, err = ioutil.ReadFile(currentPath + "\\" + fileName)
		fmt.Println("windows 读取配置")
	} else {
		data, err = ioutil.ReadFile(currentPath + "/" + fileName)
		fmt.Println("Mac 读取配置")
	}

	if err != nil {
		fmt.Println("err = ", err)
	}
	return data, err
}
