package main

import (
	"fmt"
)

var a int = 11

/* func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	years := 10.0
	expectedReturnRate := 5.5
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	var futureValue = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
	fmt.Printf("Future value is: %.1f\n\nFuture real value is: %v\n\n", futureValue, futureRealValue)
	fmt.Print(a)
} */

func main() {
	a := 10 // Declare an integer variable 'a'
	p := &a // Declare a pointer 'p' that stores the address of 'a'

	fmt.Println("a (value):", a)                // Prints the value of 'a'
	fmt.Println("p (address):", p)              // Prints the memory address stored in 'p'
	fmt.Println("*p (value at address p):", *p) // Dereferences 'p' to get the value at the address stored in 'p'

	*p = 20
	numbers := []int{1, 2, 3, 4, 5, 6}
	for _, v := range numbers {
		fmt.Println(v)
		// Dereference 'p' and update the value of 'a' indirectly
	}
}
