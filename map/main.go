package main

import "fmt"

func main() {
	// var colors map[string] string

	colors := map[string]string{
		"red":   "#ff0000",
		"black": "#000000",
		"white": "#ffffff",
	}

	// colors := make(map[string]string)

	// colors["black"] = "#000000"
	// colors["white"] = "#ffffff"

	// delete(colors, "white")

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("The hex code for", color, "is", hex)
	}
}
