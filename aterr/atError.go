package aterr

import "fmt"

func CheckErr(msg interface{}) {
	fmt.Println(msg)
	if msg != nil {
		panic(msg)
	}
}

func TestM()  {
	fmt.Println("..........")
}
