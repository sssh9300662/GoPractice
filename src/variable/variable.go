package main

import (
	"fmt"
)

var x, y, z int //int var
var c, python bool
var java= "java"

func main() {
	fmt.Println(x, y, z, c, python, java)
	x, y, z = 1, 2, 3
	c, python, java := true, false, "no!"
	fmt.Println(x, y, z, c, python, java)
}
