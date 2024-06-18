package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello World"))
	})
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		panic(err)
	}
}
