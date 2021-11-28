package main

import "fmt"


func main() {
	message := "Hola Estación Espacial Internacional"
	var decodeMessage string
	for _, m := range message {
		if m >= 'a' && m <= 'z' {
			decodeMessage += string(((m + 13 - 'a') % 26) + 'a')
			continue
		}

		if m >= 'A' && m <= 'Z' {
			decodeMessage += string(((m + 13 - 'A') % 26) + 'A')
			continue
		}

		decodeMessage += string(m)
	}

	fmt.Println(decodeMessage)
}