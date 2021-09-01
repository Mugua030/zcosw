package main

import (
	"log"
	"net/http"
)

// handler process all request
type handler struct{}

// ServeHTTP must from the interface
func (m *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("[ServeHTTP] request=> [%s] %s\n", r.Method, r.URL.Path)
}

func main() {
	log.Println("runing...")

	mhandler := &handler{}
	log.Fatal(http.ListenAndServe("127.0.0.1:6888", mhandler))
}
