package main

const text = "hello world"

var (
	b bool = true
	c int
	d string = "Eric"
	e float64
)

func main() {
	// var a string = "X" -> Nao permite variaveis unused
	a := "X"

	println(a)
}

func x() {
}
