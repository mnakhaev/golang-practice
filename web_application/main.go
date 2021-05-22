package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html") // needed for correct display of HTML code below
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my site!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please send email to <a href=\"mailto:test@example.com\">test@example.com</a>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>We couldn't find the page you were looking for</h1>"+
			"<p>Please email us</p>")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8080", nil)
}
