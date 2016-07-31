package calc

import (
    "strings"
    "strconv"
)

// FormulaParts blah
type FormulaParts struct {
    Operator string
    Expression1 float32
    Expression2 float32
}

//ParseFormula blah
func ParseFormula(formula string) *FormulaParts {
    f := new(FormulaParts)
    oIndex := -1
    operators := [4]string{"+", "-", "*", "/"}

    // First get the operator
    for i := 0; i < len(operators); i++ {
        oIndex = strings.Index(formula, operators[i])

        if oIndex != -1 {
            f.Operator = formula[oIndex : oIndex + 1]
            break
        }
    }

    // Next, get the first expression
    e1, _ := strconv.ParseFloat(strings.TrimSpace(formula[:oIndex]), 32)
    f.Expression1 = float32(e1)

    // Finally, get the second expression
    e2, _ := strconv.ParseFloat(strings.TrimSpace(formula[oIndex + 1 : len(formula)]), 32)
    f.Expression2 = float32(e2)

    return f
}