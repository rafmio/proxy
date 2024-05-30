package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	redirReq, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Println("Error processing request:", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("An internal server error occurred"))
		return
	}
	defer redirReq.Body.Close()

}

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("Listening...")
	log.Fatal("error listening and serve:", http.ListenAndServe(":8080", nil))
}
