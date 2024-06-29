package main

import (
	"net/http"
	"strconv"
	"time"
)

func autoChunkHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json := generateHeavyJson()
	w.Write(json)
}

func manualChunkHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json := generateHeavyJson()

	manualChunkSize := 10000

	for i := 0; i < len(json); i += manualChunkSize {
		flusher, _ := w.(http.Flusher)
		chunk := json[i : i+manualChunkSize]
		w.Write(chunk)
		time.Sleep(1 * time.Millisecond) // TCPバッファリングを回避するために1ms待つ
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
	http.ListenAndServe(":8080", nil)
}
