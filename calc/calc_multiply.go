package calc

import (
    "log"
)

// Multiply takes two numbers and adds them together
func Multiply(a float64, b float64) float64 {

    log.Printf("Multiply(), a=%f, b=%f, result=%f", a, b, a*b);

    return a*b;
}