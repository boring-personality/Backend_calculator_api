package main

import (
	"fmt"
	"net/http"
)

type Config struct{}

const webPort = 3000

func main() {
	fmt.Println("Starting the calculator app")
	app := Config{}
	mux := http.NewServeMux()
	app.indexHandler(mux)
	app.addHandler(mux)
	app.subtractHandler(mux)
	app.multiplyHandler(mux)
	app.divideHandler(mux)
	url := fmt.Sprintf("localhost:%d", webPort)

	if err := http.ListenAndServe(url, mux); err != nil {
		fmt.Println(err.Error())
	}
}
