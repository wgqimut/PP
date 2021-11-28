package main

import "fmt"

func main() {
	message := "L fdph, L vdz, L frqtxhuhg."

	var decodeMessage string
	for _, m := range message {
		if m >= 'a' && m <= 'z' {
			decodeMessage += string(((m - 'a' + 26 -3) %  26) + 'a')
			continue
		}

		if m >= 'A' && m <= 'Z' {
			decodeMessage += string(((m - 'A' + 26 -3) %  26) + 'A')
			continue
		}

		decodeMessage += string(m)
	}
	fmt.Println(decodeMessage)
}