package main

import (
	"io"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// получаем URL из первоначального запроса
	targetURL := r.URL.Query().Get("url")
	if targetURL == "" {
		http.Error(w, "URL parameter 'url' is required", http.StatusBadRequest)
		return
	}

	// формируем новый запрос к целевому URL
	req, err := http.NewRequest(r.Method, targetURL, r.Body)
	if err != nil {
		http.Error(w, "Error creating request to target URL", http.StatusInternalServerError)
		return
	}

	// копируем заголовки из первоначального запроса
	for header, values := range r.Header {
		for _, value := range values {
			req.Header.Add(header, value)
		}
	}

	// отправляем запрос к целевому URL
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error sending request to target URL", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// копируем заголовки из ответа целевого URL в ответ первоначального запроса
	for header, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(header, value)
		}
	}

	// устанавливаем статус-код ответа
	w.WriteHeader(resp.StatusCode)

	// копируем тело ответа целевого URL в ответ первоначального запроса
	io.Copy(w, resp.Body)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServeTLS(":4443", "server.crt", "server.key", nil))
}
