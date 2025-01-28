package main

import (
	"fmt"
	"net/http"
)

const webPort = "3000"

func main() {
	fmt.Println("Starting the calculator app")
	mux := http.NewServeMux()
	add(mux)
	subtract(mux)
	multiply(mux)
	divide(mux)
	url := "localhost:" + webPort

	if err := http.ListenAndServe(url, mux); err != nil {
		fmt.Println(err.Error())
	}
}
