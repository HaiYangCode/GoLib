package atfile

import (
	"testing"
	"fmt"
)

func TestGetCurrentDirectory(t *testing.T) {
	fmt.Println("current file directory = ",GetCurrentDirectory())
}