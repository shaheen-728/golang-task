package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the number for fibonacci")
	fmt.Scanf("%d", &n)
	fmt.Println("The number is ", n)
	last_digit := fibonacci_last_digit(n) //call fibonacci_last_digit function
	fmt.Println("the last digit of nth fibonacci number is ", last_digit)

}

//calculate fibonacci_last_digit function to find the last digit of nth fibonacci number
func fibonacci_last_digit(n int) int {
	first, second := 0, 1
	var res int
	for i := 2; i <= n; i++ {
		res = (first + second) % 10
		first = second
		second = res
	}
	return res
}
