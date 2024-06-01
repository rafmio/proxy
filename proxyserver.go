package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	redirReq, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Println("Error processing request:", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("DefaultClient.Do(): An internal server error occurred"))
		return
	}
	defer redirReq.Body.Close()

	byteBody, err := io.ReadAll(redirReq.Body)
	if err != nil {
		log.Println("Error processing request:", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("An internal server error occurred"))
		return
	}

	w.Header().Set("Status", "Ok")
	w.Header().Set("Done", "True")
	w.Write(byteBody)
}

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Usage: go run main.go <url>")
		return
	}
	url := os.Args[1]

	http.HandleFunc("/", handleRequest)
	fmt.Println("Listening...")
	log.Fatal("error listening and serve:", http.ListenAndServe(":8080", nil))
}
