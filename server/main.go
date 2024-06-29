package main

import (
	"net/http"
	"strconv"
)

func heavyJsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json := generateHeavyJson()
	w.Write(json)
}

func generateHeavyJson() []byte {
	json := []byte("[")
	for i := 0; i < 100000; i++ {
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
	http.HandleFunc("/", heavyJsonHandler)
	http.ListenAndServe(":8080", nil)
}
