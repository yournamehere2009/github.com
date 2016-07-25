package calc_test

import (
    "testing"
    "calc"
)

func TestAdd(t *testing.T){
    if calc.Add(1, 1) != 2 {
        t.Error("expected 2")
    }
}

func TestSubtract(t *testing.T) {
    if calc.Subtract(2, 1) != 1 {
        t.Error("expecting 1")
    }
}

func TestMultiply(t *testing.T) {
    if calc.Multiply(2, 2) != 4 {
        t.Error("expecting 4")
    }
}

func TestDivide(t *testing.T) {
    if calc.Divide(4, 2) != 2 {
        t.Error("expecting 2")
    }
}

func TestParseFormulaAddition(t *testing.T) {
    formula := calc.ParseFormula("10+2")

    if formula.Operator != "+" {
        t.Error("expecting operator: addition")
    } else if formula.Expression1 != 10 {
        t.Error("expecting expression1: 10")
    } else if formula.Expression2 != 2 {
        t.Error("expecting expression2: 2")
    }
}

func TestParseFormulaSubtraction(t *testing.T) {
    formula := calc.ParseFormula("7-2")

    if formula.Operator != "-" {
        t.Error("expecting operator: subtraction")
    } else if formula.Expression1 != 7 {
        t.Error("expecting expression1: 7")
    } else if formula.Expression2 != 2 {
        t.Error("expecting expression2: 2")
    }
}

func TestParseFormulaDivision(t *testing.T) {
    formula := calc.ParseFormula("15/5")

    if formula.Operator != "/" {
        t.Error("expecting operator: division")
    } else if formula.Expression1 != 15 {
        t.Error("expecting expression1: 15")
    } else if formula.Expression2 != 5 {
        t.Error("expecting expression2: 5")
    }
}

func TestParseFormulaMultiplication(t *testing.T) {
    formula := calc.ParseFormula("20*3")

    if formula.Operator != "*" {
        t.Error("expecting operator: multiplication")
    } else if formula.Expression1 != 20 {
        t.Error("expecting expression1: 20")
    } else if formula.Expression2 != 3 {
        t.Error("expecting expression2: 3")
    }
}
