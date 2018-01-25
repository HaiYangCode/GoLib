package aterr

import "fmt"

func CheckErr(err error) {
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
}
