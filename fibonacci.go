package piscine

func Fibonacci(index int) int {
	if index < -25 || index > 25 {
		return 0
	}
	if index == 0 {
		return 0
	}
	if index == 1 {
		return 1
	}
	if index > 1 {
		return Fibonacci(index-2) + Fibonacci(index-1)
	}
	return -1
}
