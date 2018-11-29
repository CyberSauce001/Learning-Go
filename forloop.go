package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	
	sum1 := 1
	for sum1 < 1000; {
		sum1 += sum1
	}
	fmt.Println(sum1)
	
	func sqrt(x float64) string {
		if x < 0 {
			return sqrt(-x) + "i"
		}
		return fmt.Sprint(math.Sqrt(x))
	}

	fmt.Println(sqrt(2), sqrt(-4))
}
