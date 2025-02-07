package main

import (
	"fmt"
	"math"
)

var (
	a int = 1
)

func main() {
	var a int = 10
	var b int = 2

	println(a + b)
	println(b - a)

	if b != 0 {
		println(a / b)
		println(a % b)
	}

	println(math.Cos(float64(b)))

	println(`He said:"Hello Go!" `)
	println("`")

	println(fmt.Sprintf("hello %d", a))
	//strings.Count()

	var str string = "this is string"
	var bs []byte = []byte(str)
	bs[0] = 'T'
	println(str, bs)
}
