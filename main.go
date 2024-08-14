// main.go
package main

import (
    "fmt"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    hostname, err := os.Hostname()
    if err != nil {
        http.Error(w, "Could not get hostname", http.StatusInternalServerError)
        return
    }
    fmt.Fprintf(w, "Hello, World! Running on container: %s", hostname)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Failed to start server:", err)
    }
}

