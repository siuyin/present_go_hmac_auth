package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"siuyin/present_go_hmac_auth/mac"
)

//010 OMIT
func main() {
	log.Println("Server starting.")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tp := mac.TimePeriod(time.Now(), time.Second)
		keyB := mac.TimeKey([]byte("ourSecret"), tp-1)
		keyN := mac.TimeKey([]byte("ourSecret"), tp) // HL
		keyA := mac.TimeKey([]byte("ourSecret"), tp+1)
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "URL: Path: %q Query: %q", r.URL.Path, r.URL.Query())

		authB := mac.CheckValues(r.URL.Query(), "hmac", keyB)
		authN := mac.CheckValues(r.URL.Query(), "hmac", keyN)
		authA := mac.CheckValues(r.URL.Query(), "hmac", keyA)
		log.Printf("auth check %v: %t", r.URL.Query(), authB || authN || authA)
		fmt.Fprintf(w, "<p>Authenticated: %t", authB || authN || authA)
	})
	log.Fatal(http.ListenAndServe(":8123", nil))
}

//020 OMIT
