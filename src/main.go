package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello Gophers!")
	port := 3000
	_, err := startWebServer(port, 2)
	fmt.Println(err)
}

func startWebServer(port int, numberOfRetries int) (int, error) {
	fmt.Println("Starting server...")
	// do important things
	fmt.Println("Server started  on port", port)
	fmt.Println("Number of retries", numberOfRetries)

	return port, errors.New("Something went wrong")
}
