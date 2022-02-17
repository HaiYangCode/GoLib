package GoLib

import "testing"

func TestCheckErr1(t *testing.T) {

}

func TestCheckErr(t *testing.T) {
	err := &CustomError{}
	CheckErr(err)
}

type CustomError struct {
}

func (c *CustomError) Error() string {

	return "xxxxx"
}
