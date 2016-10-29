package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"siuyin/present_go_hmac_auth/mac"
)

//010 OMIT
func main() {
	key := mac.TimeKey([]byte("ourSecret"), mac.TimePeriod(time.Now(), time.Second)) // HL
	path := "/a/b"
	qry := url.Values{}
	qry.Add("c", "cat")
	qry.Add("d", "dog")
	qry = mac.Values(qry, "hmac", key) // HL

	res, err := http.Get("http://localhost:8123" + path + "?" + qry.Encode())
	if err != nil {
		log.Fatalf("bad get: %v", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("bad read: %v", err)
	}
	fmt.Printf("%s", body)
}

//020 OMIT
