package main

import (
	"fmt"
)

func main() {
	cipherText := "ECFRZKYGLGRMUSDHRXK"
	keyword := "GOLANG"

	var dText string
	for i := 0; i < len(cipherText); i++ {
		dText += string((cipherText[i]-keyword[i%len(keyword)]+26)%26 + 'A')
	}
	fmt.Println(dText)
}
