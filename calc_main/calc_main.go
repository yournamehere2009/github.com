package main

import (
    "fmt"
    "github.com/calc"
    // "bufio"
    // "os"
    // "strings"
    "net/http"
	"strconv"
)

func handler (w http.ResponseWriter, r *http.Request){
    var html string

    html = "<html><head><title>Mike's Go Calculator - Result</title></head>";

    html += "<form action\"/calculate/\" method=\"POST\">" +
        "<input type=\"text\" name=\"formula\" >" +
        "<input type=\"submit\" value=\"Go!\" >" +
        "</form>";
    
    if r.FormValue("formula") != "" {
       if  result, err := calc.ComputeFormula(r.FormValue("formula")); err == nil {
           html += "<div>Answer: " + strconv.FormatFloat(float64(result), 'f', -1, 64) + "</div";
       }
    }

    html += "</body</html>";

    fmt.Fprintf(w, html);
}

func main () {
    http.HandleFunc("/", handler);
    http.ListenAndServe(":8080", nil);


    // reader := bufio.NewReader(os.Stdin)

    // fmt.Println("Welcome to Mike's first Go calculator!")
    // fmt.Println("This basic calculator supports addition, subtraction, multiplication, and division.")

    // for {
    //     fmt.Print("\nWhat can I compute for you? (type 'exit' to stop): ")
    //     formula, _ := reader.ReadString('\n')

    //     if (strings.TrimSpace(strings.ToLower(formula)) == "exit") {
    //         break;
    //     }

    //     result, err := calc.ComputeFormula(formula)

    //     if err != nil {
    //         fmt.Printf("%v\n", err);
    //     } else {
    //         fmt.Printf("The answer is %g\n", result)
    //     }
    // }
}