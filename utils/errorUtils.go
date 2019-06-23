package utils

import "fmt"

func PrintErr(err error) {
	if err == nil {
		return
	}
	fmt.Println(err)
}
