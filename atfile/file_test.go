package atfile

import (
	"testing"
	"fmt"
)

func TestGetCurrentDirectory(t *testing.T) {
	fmt.Println("current file directory = ",GetCurrentDirectory())
	//fmt.Println("current file directory = ",GetCurrentRelDir())
}

func TestGetCurrentRelDir(t *testing.T) {

}