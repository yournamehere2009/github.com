package calc

import (
    "testing"
)

var computeTests = []struct {
  n        string // input
  expected float64 // expected result
}{
  {"2+2", 4},
  {"10*5", 50},
  {"4", 4},
  {"-3*2", -6},
}

var decomposeTests = []struct {
  n        string // input
  expected string // expected result
}{
  {"(2+2)", "4"},
  {"2 * (3+5)", "2 * 8"},
  {"(2*(3+5))", "16"},
  {"(2+5) * (3+5)", "7 * 8"},
  {"((10*2)/5)", "4"},
}

func TestCompute (t *testing.T){
    for _ , tt := range computeTests {
        contents := tt.n;

        f, _ := ParseFormula(contents);
        result, _ := compute(f);

        if result != tt.expected {
            t.Errorf("compute(%v): expected %f, actual %f", tt.n, tt.expected, result);
        }
    }
}

func TestDecompose (t *testing.T) {
    for _ , tt := range decomposeTests {
        formula := tt.n;

        formula, _ = decompose(formula);

        if formula != tt.expected {
            t.Errorf("decompose(%v): expected %v, actual %v", tt.n, tt.expected, formula);
        }
    }
}