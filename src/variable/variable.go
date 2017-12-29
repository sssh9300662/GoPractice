package main

import (
	"fmt"
)

/**
Outside a function, every statement begins with a keyword,
the := construct is not available.
**/
var x, y, z int    //int var, default value is zero
var c, python bool //default value is false
var java = "java"  //default value for string is "" and it's immutable

func main() {
	fmt.Println(x, y, z, c, python, java)
	x, y, z = 1, 2, 3
	//:= short assignment statement can be used in place of a var declaration with implicit type.
	c, python, java := true, false, "no!"
	fmt.Println(x, y, z, c, python, java)
}
