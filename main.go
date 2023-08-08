package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		aStr := r.FormValue("a")
		bStr := r.FormValue("b")

		a, _ := strconv.Atoi(aStr)
		b, _ := strconv.Atoi(bStr)

		c := a + b
		result := fmt.Sprintf("Result: %d", c)

		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Addition Result</title>
		</head>
		<body>
			<h1>Addition Result</h1>
			<p>%s</p>
			<form method="post">
				<label for="a">Enter a:</label>
				<input type="text" name="a"><br>
				<label for="b">Enter b:</label>
				<input type="text" name="b"><br>
				<button type="submit">Calculate</button>
			</form>
		</body>
		</html>
		`

		fmt.Fprintf(w, html, result)
	} else {
		http.ServeFile(w, r, "index.html")
	}
}

func main() {
	http.HandleFunc("/", addHandler)
	http.ListenAndServe(":8080", nil)
}
