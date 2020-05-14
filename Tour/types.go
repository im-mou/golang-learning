package main

import(
	"fmt"
	"math"
	"math/cmplx"
)

// declare var types
var (
	boolean bool 	= false
	maxInt uint64 	= 1<<64 - 1
	z complex128 	= cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Println("\n========= types =========")
	fmt.Printf("type: %T value: %v\n", boolean, boolean)
	fmt.Printf("type: %T value: %v\n", maxInt, maxInt)
	fmt.Printf("type: %T value: %v\n", z,z)

	/* user input

	var str string
	fmt.Scanln(&str)
	fmt.Printf("input from the user %v eith the type %T\n", str, str)
	
	*/

	// if values are not initialized they're set to 0, false, ""
	fmt.Println("\n========= initialization =========")
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// type conversion
	fmt.Println("\n========= type conversion =========")
	var x,y int = 3,4
	f = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x,y,z)

	// type inference
	fmt.Println("\n========= type inference =========")
	value1 := false
	value2 := "this is a string"
	value3 := 45
	value4 := 4.0
	fmt.Printf("value type is %T\n", value1)
	fmt.Printf("value type is %T\n", value2)
	fmt.Printf("value type is %T\n", value3)
	fmt.Printf("value type is %T\n", value4)

}