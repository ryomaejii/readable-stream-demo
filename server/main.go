package main

import (
	"io/ioutil"
	"net/http"
)

func autoChunkHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	message := getMessageFromFile("message.txt")
	w.Write(message)
}

func manualChunkHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	message := getMessageFromFile("message.txt")

	manualChunkSize := 1024 // 1KB

	for i := 0; i < len(message); i += manualChunkSize {
		flusher, _ := w.(http.Flusher)
		end := i + manualChunkSize
		if end > len(message) {
			end = len(message)
		}
		chunk := message[i:end]
		w.Write(chunk)
		flusher.Flush()
	}
}

func getMessageFromFile(filename string) []byte { // getMessageFromFile を []byte 型に変更
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte("Failed to read message")
	}
	return content
}

func main() {
	http.HandleFunc("/auto-chunk", autoChunkHandler)
	http.HandleFunc("/manual-chunk", manualChunkHandler)

	server := &http.Server{
		Addr:      ":8080",
		TLSConfig: nil, // Omit for automatic configuration
	}

	// Start server with HTTP/2.0 support
	err := server.ListenAndServeTLS("cert.pem", "key.pem")
	if err != nil {
		panic(err)
	}
}
