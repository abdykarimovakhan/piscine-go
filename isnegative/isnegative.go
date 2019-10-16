package main

import "github.com/01-edu/z01"

func main() {

	int := 1

	if int >= 0 {
		z01.Println("true")
	} else if int <= 0 {
		z01.Println("false")
	}
}