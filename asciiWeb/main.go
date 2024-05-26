package main

import (
	"fmt"
	"net/http"
	"server/func"
	"strings"
	"text/template"
)

func main() {
	port := ":8080"
	http.HandleFunc("/", index)
	fmt.Printf("the server is listening on http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseGlob("templets/*.html")
	banner := r.FormValue("banner")
	input := r.FormValue("input")
	if strings.Contains(input, "\r\n") {
		input = strings.ReplaceAll(input,"\r\n", "\\n")
	}
	if err != nil {
		http.Error(w, "Internal Server Error ", http.StatusInternalServerError)
	}

	if r.URL.Path != "/" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}
	tmp.Execute(w, server.Printascii(input, banner))
}
