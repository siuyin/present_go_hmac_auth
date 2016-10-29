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
	fmt.Printf("random key example: %v\n", mac.RandID(10))
}

// 020 OMIT
