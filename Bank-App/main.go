package main

import (
	"fmt"
	"net/http"
)

// Handler function
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	server := NewAPIServer(":8000")
	fmt.Println("Server running on 8000", server)
	server.Run()
}
