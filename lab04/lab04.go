package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

// TODO: Create a struct to hold the data sent to the template

type Page struct {
	Expression string
	Result     string
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func getTemplatePath(filename string) string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir + "/" + filename
}

func Calculator(w http.ResponseWriter, r *http.Request) {
	// TODO: Finish this function
	tmpl := template.Must(template.ParseFiles(getTemplatePath("index.html")))
	errtmpl := template.Must(template.ParseFiles(getTemplatePath("error.html")))
	op := r.URL.Query().Get("op")
	num1 := r.URL.Query().Get("num1")
	num2 := r.URL.Query().Get("num2")
	var ans int
	ans_str := ""
	n1, err := strconv.Atoi(num1)
	if err != nil {
		errtmpl.Execute(w, nil)
		return
	}
	n2, err := strconv.Atoi(num2)
	if err != nil {
		errtmpl.Execute(w, nil)
		return
	}
	if op == "add" {
		ans = n2 + n1
		ans_str += num1
		ans_str += " + "
		ans_str += num2
	} else if op == "sub" {
		ans = n1 - n2
		ans_str += num1
		ans_str += " - "
		ans_str += num2
	} else if op == "mul" {
		ans = n1 * n2
		ans_str += num1
		ans_str += " * "
		ans_str += num2
	} else if op == "div" {
		if n2 == 0 {
			errtmpl.Execute(w, nil)
			return
		}
		ans = n1 / n2
		ans_str += num1
		ans_str += " / "
		ans_str += num2
	} else if op == "gcd" {
		ans = gcd(n1, n2)
		ans_str += "GCD("
		ans_str += num1
		ans_str += ", "
		ans_str += num2
		ans_str += ")"
	} else if op == "lcm" {
		ans = lcm(n1, n2)
		ans_str += "LCM("
		ans_str += num1
		ans_str += ", "
		ans_str += num2
		ans_str += ")"
	} else {
		errtmpl.Execute(w, nil)
		return
	}
	data := Page{
		Expression: ans_str,
		Result:     strconv.Itoa(ans),
	}
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", Calculator)
	log.Fatal(http.ListenAndServe(":8084", nil))
}
