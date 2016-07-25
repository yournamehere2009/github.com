package calc

import (
    "strings"
    "strconv"
)

// FormulaParts blah
type FormulaParts struct {
    Operator string
    Expression1 int
    Expression2 int
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
    f.Expression1, _ = strconv.Atoi(strings.TrimSpace(formula[:oIndex]))

    // Finally, get the second expression
    f.Expression2, _ = strconv.Atoi(strings.TrimSpace(formula[oIndex + 1 : len(formula)]))

    return f
}