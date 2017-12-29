package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

/**
Begins with a capital letter means the function can be accessed by outside class;
otherwise, its only internal usage.
Go syntax is: name-> type, this rule make code be more readble since we can read it more natually-left to right;
therefore, the below function we can illustrate it as: function add take two int and return int
**/
func add(x int, y int) int {
	return x + y
}

func printRand() {
	//For real random value, set seed first and then call the random funaction u want
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("My favorite number is ", rand.Intn(10))
}

/**
Return multiple result for a single invoke
**/
func swap(x, y string) (string, string) {
	return y, x
}

/**
Naked return by name, but should avoid using it for long bloack func
**/
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	//Basic operation
	fmt.Println("Current time is ", time.Now())
	fmt.Println("Now you have %g problems.", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println("Sum for 5 and 6 is ", add(5, 6))
	//call private method
	printRand()
	//swap
	a, b := swap("A", "B")
	fmt.Println("Swap value A and B:", a, "and", b)
	//naked return
	c, d := split(17)
	fmt.Println("split 17 and we get:", c, "and", d)

}
