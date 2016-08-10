package main

import (
	"fmt"

	"net/http"
	"strconv"

	"github.com/yournamehere2009/calc"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var html string

	html = "<html><head><title>Mike's Go Calculator - Result</title></head>"

	html += "<form action\"/calculate/\" method=\"POST\">" +
		"<input type=\"text\" name=\"formula\" >" +
		"<input type=\"submit\" value=\"Go!\" >" +
		"</form>"

	if r.FormValue("formula") != "" {
		if result, err := calc.ComputeFormula(r.FormValue("formula")); err == nil {
			html += "<div>Answer: " + strconv.FormatFloat(float64(result), 'f', -1, 64) + "</div"
		}
	}

	html += "</body</html>"

	fmt.Fprintf(w, html)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
