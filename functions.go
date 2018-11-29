package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x , y int) int {
	return x + y
}

func swap (x,y string)(string,string) {
	return y, x
}

func split(sum int)(x, y int) {
	x = sum * 4/9
	y = sum - x
	return
}

func main() {
	fmt.Println(rand.Intn(10)"\n")
	fmt.Printf("Square root of 81 is %g \n", math.Sqrt(81))
	fmt.Println(" Pi is %g", math.Pi)
	fmt.Println(add(100,99))
	
	a, b := swap("John", "Doe")
	fmt.Println(a,b)
	
	fmt.Println(split(20))
}


