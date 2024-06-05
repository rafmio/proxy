package main

import (
	"net/http"
)

var customTransport = http.DefaultTransport

// func init() {
// 	customTransport = &http.Transport{
// 		Proxy: http.ProxyFromEnvironment,
// 		DialContext: (&net.Dialer{
// 			Timeout:   30 * time.Second,
// 			KeepAlive: 30 * time.Second,
// 		}).DialContext,
// 		MaxIdleConns:          100,
// 		IdleConnTimeout:       90 * time.Second,
// 		TLSHandshakeTimeout:   10 * time.Second,
// 		ExpectContinueTimeout: 1 * time.Second,
// 	}
// }

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// create a new HTTP request with the same method, URL, and body as the original request
	targetURL = r.URL
	proxyReq, err := http.NewRequest(r.Method, targetURL.String(), r.Body)
	if err != nil {
		http.Error(w, "Error creating proxy request", http.StatusInternalServerError)
		return
	}

	// copy the headers from the original request to the proxy request
	for name, values := range r.Header {
		for _, value := range values {
			proxyReq.Header.Add(name, value)
		}
	}
}
