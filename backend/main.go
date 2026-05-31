package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from PlatformPilot!")
}

func version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "v2")
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "healthy")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/version", version)
	http.HandleFunc("/health", health)

	fmt.Println("Server running on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
