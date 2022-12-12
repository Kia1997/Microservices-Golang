package main

import (
	"errors"
	"fmt"
)

func main() {
	port := 3000
	// saljemo ARGUMENTE
	data, err := startWebServer(port, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

	_, err = startWebServer(port, 3)
}

// primamo PARAMETAR
func startWebServer(port int, numberOfRetries int) (int, error) {
	fmt.Println("Starting server...")
	// Logic
	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)
	return 17, errors.New("Something went wrong")
}
