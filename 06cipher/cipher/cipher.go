package main

import (
	"fmt"
	"strings"
)

func main() {
	plainText := "your message goes here"
	keyword := "GOLANG"

	plainText = strings.Replace(plainText, " ", "", -1)
	plainText = strings.ToUpper(plainText)

	var cipherText string
	for i := 0; i < len(plainText); i++ {
		cipherText += string((plainText[i] + keyword[i % len(keyword)]) % 26 + 'A')
	}
	fmt.Println(cipherText)
}
