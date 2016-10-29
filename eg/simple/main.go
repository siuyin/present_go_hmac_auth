package main

import (
	"fmt"
	"siuyin/present_go_hmac_auth/mac"
)

// 010 OMIT
func main() {
	key := []byte("ourSecret")
	msg := []byte("Brown Fox")
	fmt.Printf("%x\n", mac.MAC(msg, key))
}

// 020 OMIT
