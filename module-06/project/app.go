package main

func main() {
	result := add(1, 2)
	result += 1
}

func add[T string | int | float64](a, b T) T {
	return a + b
}
