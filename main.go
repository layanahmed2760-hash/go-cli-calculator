package main

import (
	"bufio"
	"fmt"
	"os"
)

func add(a, b float64) float64 {
	return a + b
}
func subtract(a, b float64) float64 {
	return a - b
}
func multiply(a, b float64) float64 {
	return a * b
}
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		var num1, num2 float64
		var operator string

		fmt.Println("Enter first number: ")
		fmt.Fscanln(reader, &num1)
		fmt.Println("Enter operator (+, -, *, /):")
		fmt.Fscanln(reader, &operator)
		if operator == "exit" {
			fmt.Println("Exiting the calculator.")
			break
		}
		fmt.Println("Enter second number: ")
		fmt.Fscanln(reader, &num2)

		var result float64
		var err error
		switch operator {
		case "+":
			result = add(num1, num2)
		case "-":
			result = subtract(num1, num2)
		case "*":
			result = multiply(num1, num2)
		case "/":
			result, err = divide(num1, num2)
		default:
			fmt.Println("Invalid operator used enter a valid operator.")
			continue
		}
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
		fmt.Println("Result:", result)
	}

}
