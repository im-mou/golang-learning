package main

import "fmt"

const PI = 3.14

const (
	Big = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	const MyConstant = "This is my contant"

	fmt.Println("hey there", MyConstant)
	fmt.Println("this value of the PI is:", PI)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}