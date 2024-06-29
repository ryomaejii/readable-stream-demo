package main

import (
	"net/http"
	"strconv"
)

func autoChunkHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json := generateHeavyJson()
	w.Write(json)
}

func manualChunkHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json := generateHeavyJson()

	manualChunkSize := 1024 // 1KB

	for i := 0; i < len(json); i += manualChunkSize {
		flusher, _ := w.(http.Flusher)
		chunk := json[i : i+manualChunkSize]
		w.Write(chunk)
		flusher.Flush()
	}
}

func generateHeavyJson() []byte {
	json := []byte("[")
	for i := 0; i < 20000; i++ {
		json = append(json, []byte(`{"id":`)...)
		json = append(json, []byte(`"`)...)
		json = append(json, []byte(strconv.Itoa(i))...)
		json = append(json, []byte(`"}`)...)
		if i != 9999 {
			json = append(json, []byte(",")...)
		}
	}
	json = append(json, []byte("]")...)

	return json
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
