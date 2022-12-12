package main

import (
	"fmt"
	"net/http"

	"github.com/vsjadeja/middleware/middleware"
)

func main() {
	http.Handle(`/`,middleware.LoggerMiddleware(middleware.LimiterMiddleware(http.HandlerFunc(handler))))
	err := http.ListenAndServe(":1111", nil)
	if err != nil {
		fmt.Println("Error: setting up server")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside the handler")
	_, err := w.Write([]byte("OK - executing handler"))
	if err != nil {
		fmt.Println("Error writing response")
	}
}
