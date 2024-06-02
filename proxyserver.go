package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {

	fmt.Println("=======Received request:=======")
	fmt.Println("Method:", r.Method)
	fmt.Printf("URL: %v\n", r.URL)
	fmt.Println("Proto:", r.Proto)
	for head, val := range r.Header {
		fmt.Println(head, ":", val)
	}
	fmt.Println("Host:", r.Host)
	fmt.Println("===============================")

	// создаем новый запрос, используя информацию из исходного запроса
	newReq, err := http.NewRequest(r.Method, r.URL.String(), nil)
	if err != nil {
		log.Println("Error processing request:", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Error creating new request"))
		return
	}

	// копируем заголовки из исходного запроса
	for key, values := range r.Header {
		for _, value := range values {
			newReq.Header.Add(key, value)
		}
	}

	redirReq, err := http.DefaultClient.Do(newReq)
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
	http.HandleFunc("/", handleRequest)
	fmt.Println("Listening...")
	log.Fatal("error listening and serve:", http.ListenAndServe(":8080", nil))
}
