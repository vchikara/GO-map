package main

import "fmt"

func main() {

	//var colors map[string]string
	colors := map[string]string{
		"red":    "#ff0000",
		"greeen": "#fffff",
	}
	//colors := make(map[string]string)
	//colors["white"] = "#fffff"
	//fmt.Println(colors)
	printMap(colors)
}
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex code for", color, "is", hex)
	}
}
