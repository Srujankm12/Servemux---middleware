package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello World"))
}
func ByeHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Bye World"))
}
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {

	// http.HandleFunc("/bye", ByeHandler)
	// fmt.Println("server started at 8080")
	// http.ListenAndServe(":8080", nil)

	mux := http.NewServeMux()
	fmt.Println("server started at 8080")
	mux.HandleFunc("/", HelloHandler)
	mux.HandleFunc("/bye", ByeHandler)
	loggedmux := loggingMiddleware(mux)

	http.ListenAndServe(":8080", loggedmux)

}
