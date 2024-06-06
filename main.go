package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Mondoo Engineer!")
	})

	port := 8080
	port_str, exists := os.LookupEnv("PORT")
	if exists {
		port, err := strconv.Atoi(port_str)
		if err != nil || (port < 0 || port > 65535) {
			log.Fatalf("Invalid port: %s", port_str)
		}
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
