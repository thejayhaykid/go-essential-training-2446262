package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`

	text = strings.ToLower(text)
	textSlice := strings.Fields(text)
	textMap := map[string]int{}

	for _, value := range textSlice {
		_, ok := textMap[value]
		if !ok {
			textMap[value] = 1
		} else {
			textMap[value] = textMap[value] + 1
		}
	}

	fmt.Println(textMap)
}
