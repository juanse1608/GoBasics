package main

// Import libraries and packages
import "fmt"

func square_area(base int) int {
	return base * base
}

func main() {
	// Constant definition
	const euler_0 float64 = 2.72
	const euler_1 = 2.71828

	// Print some lines
	fmt.Println("This is my first line in go! Yuuuuup!")
	fmt.Println("Euler Type 0:", euler_0)
	fmt.Println("Euler Type 1:", euler_1)

	// Variable definition
	age := 27
	var high int = 185
	var weight int

	fmt.Println("Vars:", age, high, weight)

	base_square := 10
	fmt.Println("Area of square with base", base_square, "is", square_area(base_square))

}
