package main

import (
	"fmt"
	"net/http"
)

// handler is an HTTP handler that serves a predefined string.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	// Serve on local port 8080 for demonstration
	http.HandleFunc("/", handler)
	fmt.Println("Server is listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
