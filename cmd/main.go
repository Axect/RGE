package main

import (
	"fmt"

	. "github.com/Axect/RGE"
)

func main() {
	mt, xi := 170.85, 50.
	var C Container
	C.SolveRGE(mt, xi)

	fmt.Println(C[0], C[Step-2], C[Step-1])
}
