package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint("Error dumping request:", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(requestDump)

	if err != nil {
		log.Println("Error writing response:", err)
		return
	}
}

func main() {
	http.HandleFunc("/", echoHandler)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	log.Printf("Server listening on %s", addr)
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
