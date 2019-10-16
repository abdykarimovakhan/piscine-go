package main

import "github.com/01-edu/z01"

func IsNegative(nb int) {
	int := 1
	if nb < 0 {
		z01.PrintRune('T')
	} else if int <= 0 {
		z01.PrintRune('F')
	}
}