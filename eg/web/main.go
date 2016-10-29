package main

import (
	"fmt"
	"log"
	"net/http"

	"siuyin/present_go_hmac_auth/mac"
)

//010 OMIT
func main() {
	key := []byte("ourSecret")
	log.Println("Server starting.")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "URL: Path: %q Query: %q", r.URL.Path, r.URL.Query())

		auth := mac.CheckValues(r.URL.Query(), "hmac", key)
		log.Printf("auth check %v: %t", r.URL.Query(), auth)
		fmt.Fprintf(w, "<p>Authenticated: %t", auth)
	})
	log.Fatal(http.ListenAndServe(":8123", nil))
}

//020 OMIT
