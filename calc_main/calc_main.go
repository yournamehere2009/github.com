package main

import (
    "fmt"
    "github.com/calc"
    "bufio"
    "os"
    "strings"
)

func main () {
    reader := bufio.NewReader(os.Stdin)

    fmt.Println("Welcome to Mike's first Go calculator!")
    fmt.Println("This basic calculator supports addition, subtraction, multiplication, and division.")

    for {
        fmt.Print("\nWhat can I compute for you? (type 'exit' to stop): ")
        formula, _ := reader.ReadString('\n')

        if (strings.TrimSpace(strings.ToLower(formula)) == "exit") {
            break;
        }

        fmt.Printf("The answer is %g\n", calc.ComputeFormula(formula))
    }
}