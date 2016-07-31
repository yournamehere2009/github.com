package calc

import (
    "strings"
    "strconv"
)

//ComputeFormula blah
func ComputeFormula (formula string) float32 {
    return compute(ParseFormula(decompose(formula)))
}

func decompose (formula string) string {
    // If we still have parenthesis in the formula
    if strings.Index(formula, "(") != -1 {
        // Find an opening followed by a closing with no other parenthesis in-between

        //Starting parenthesis
        iOpening := 0
        iClosing := 0

        // Loop over every character
        for i, r := range formula {
            if string(r) == "(" {
                iOpening = i;
            } else if string(r) == ")" {
                iClosing = i;
                break;
            }
        }

        contents := formula[iOpening + 1: iClosing]

        f := ParseFormula(contents)

        result := compute(f)

        formula = strings.Replace(formula, formula[iOpening : iClosing + 1], strconv.FormatFloat(float64(result), 'f', -1, 32), -1)

        formula = decompose(formula)
    }
    return formula
}

func compute (fp *FormulaParts) float32 {
    var result float32
    switch fp.Operator {
        case "+":
            result = Add(fp.Expression1, fp.Expression2)
        case "-":
            result = Subtract(fp.Expression1, fp.Expression2)
        case "*":
            result = Multiply(fp.Expression1, fp.Expression2)
        case "/":
            result = Divide(fp.Expression1, fp.Expression2)
    }
    return result
}