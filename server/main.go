package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
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

func eventStreamHandler(w http.ResponseWriter, r *http.Request) {
	// Set the appropriate headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Send an initial message indicating connection is established
	fmt.Fprintf(w, "data: Connected to event stream\n\n")
	flusher, _ := w.(http.Flusher)
	flusher.Flush()

	// Simulate sending events periodically
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Send a new event
			eventData := fmt.Sprintf("data: Event occurred at %v\n\n", time.Now())
			fmt.Fprintf(w, eventData)
			flusher.Flush()
		case <-r.Context().Done():
			fmt.Println("Client closed connection")
			return
		}
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
	http.HandleFunc("/server-sent-events", eventStreamHandler)

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
