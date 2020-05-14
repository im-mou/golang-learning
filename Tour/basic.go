package main

import (
	"fmt"
	"time"
	"math"
	"math/rand"
)

func main(){
	fmt.Println("Basic Script")

	// print time
	fmt.Println("print the time: it's:", time.Now())

	// generate a random number + seed (unix timestamp)
	rand.Seed(time.Now().Unix())
	fmt.Println("Random number: ", rand.Intn(100))

	// square root from the math package
	var x float64 = 12
	fmt.Println("square root of",x ,"is", math.Sqrt(x))

	// exporting names from packages
	fmt.Println("π = ",math.Pi)
}