package main

type str string

func (s str) print() {
	println(s)
}

func main() {
	val := str("Hello")
	val.print()
}
