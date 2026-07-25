package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
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

type Calculation struct {
	num1     float64
	num2     float64
	operator string
	result   float64
}
func tokenize(expr string) []string {
	return strings.Fields(expr)
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	var history []Calculation


	for {
		fmt.Println("\n--- Menu ---")
		fmt.Println("1. Calculate")
		fmt.Println("2. View History")
		fmt.Println("3. Clear History")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		var choice string
		fmt.Fscanln(reader, &choice)

		switch choice {
		case "1":
			var num1, num2 float64
			var operator string

			fmt.Println("Enter first number: ")
			fmt.Fscanln(reader, &num1)
			fmt.Println("Enter operator (+, -, *, /):")
			fmt.Fscanln(reader, &operator)
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
				fmt.Println("Invalid operator used, enter a valid operator.")
				continue
			}

			if err != nil {
				fmt.Println("error:", err)
				continue
			}

			fmt.Println("Result:", result)
			history = append(history, Calculation{num1, num2, operator, result})

		case "2":
			if len(history) == 0 {
				fmt.Println("No history yet.")
				continue
			}
			for i, calc := range history {
				fmt.Printf("%d: %.2f %s %.2f = %.2f\n", i+1, calc.num1, calc.operator, calc.num2, calc.result)
			}

		case "3":
			history = []Calculation{}
			fmt.Println("History cleared.")

		case "4":
			fmt.Println("Exiting the calculator.")
			return

		default:
			fmt.Println("Invalid menu option. Choose 1-4.")
		}
	}
}
