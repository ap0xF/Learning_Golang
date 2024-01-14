package main

import (
	"fmt"
)

func main(){
	var x int
	fmt.Println("Enter first number: ")
	fmt.Scanln(&x)

	var y int
	fmt.Println("Enter second number: ")
	fmt.Scanln(&y)

	var z string
	fmt.Println("Enter the operator: ")
	fmt.Scanln(&z)

	switch z {
	case "+":
		fmt.Printf("Result: %d\n",x+y);
	case "-":
		fmt.Printf("Result: %d\n",x-y);
	case "*":
		fmt.Printf("Result: %d\n",x*y);			
	case "/":
		fmt.Printf("Result: %d\n",x/y);
	default:
		fmt.Printf("Invalid operator");
	}
}