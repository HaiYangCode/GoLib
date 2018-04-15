package aterr

import "testing"

func TestCheckErr(t *testing.T) {
	err := &CustomError{}
	CheckErr(err)
}

type CustomError struct {
}

func (c *CustomError) Error() string {

	return "xxxxx"
}

