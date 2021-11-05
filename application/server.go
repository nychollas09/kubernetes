package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		name := os.Getenv("NAME")
		age := os.Getenv("AGE")
		response.Write([]byte(name + " is " + age + " years old"))
	})

	http.HandleFunc("/secret", func(response http.ResponseWriter, request *http.Request) {
		user := os.Getenv("USER")
		password := os.Getenv("PASSWORD")
		response.Write([]byte("User: " + user + " Password: " + password))
	})

	http.HandleFunc("/health", func(response http.ResponseWriter, request *http.Request) {
		duration := time.Since(startedAt)

		if duration.Seconds() > 15 {
			response.WriteHeader(500)
			response.Write([]byte(fmt.Sprintf("Server is Down %f", duration.Seconds())))
			return
		}

		response.WriteHeader(200)
		response.Write([]byte("Alive"))

	})
	http.ListenAndServe(":8282", nil)
}
