package main

import (
	"fmt"
)

// add operation 
func add(x int, y int) int {
	return x + y
}

// substract function with fancy params
func sub(x,y int) int {
	return x - y
}

// function to return multiple values
func swap(s1,s2 string) (string, string) {
	return s2, s1
}

// naked return
func split(sum int) (x,y int) {
	x = sum * 4/9
	y = sum - x
	return
}

func main(){
	var x,y int = 5,6
	fmt.Println("sum of",x,"and",y,"is:", add(x,y))
	fmt.Println("substraction of",x,"and",y,"is:", sub(x,y))
	fmt.Println(swap("Swapped","value"))
	x,y = split(3)
	fmt.Println(y)

}
