package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, "https://google.com", 302)
	})

	err := http.ListenAndServe(":80", nil)
	fmt.Println(err.Error())
}
