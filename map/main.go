package main

import (
	"fmt"
)

func main() {
	// var colors map[string]string

	colors := make(map[string]string)

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#23bf745",
	// }

	colors["white"] = "#ffffff"

	fmt.Println(colors)
}
