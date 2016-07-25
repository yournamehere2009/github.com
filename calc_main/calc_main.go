package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/calc"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	result := 0

	fmt.Println("Welcome to Mike's first Go calculator!")
	fmt.Println("This basic calculator supports addition, subtraction, multiplication, and division.")

	for {
		fmt.Print("\nWhat can I compute for you? (type 'exit' to stop): ")
		formula, _ := reader.ReadString('\n')

		if strings.TrimSpace(strings.ToLower(formula)) == "exit" {
			break
		}

		formulaParts := calc.ParseFormula(formula)

		switch formulaParts.Operator {
		case "+":
			result = calc.Add(formulaParts.Expression1, formulaParts.Expression2)
		case "-":
			result = calc.Subtract(formulaParts.Expression1, formulaParts.Expression2)
		case "*":
			result = calc.Multiply(formulaParts.Expression1, formulaParts.Expression2)
		case "/":
			result = calc.Divide(formulaParts.Expression1, formulaParts.Expression2)
		}

		fmt.Printf("The answer is %d\n", result)
	}
}
