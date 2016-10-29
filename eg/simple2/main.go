package main

import (
	"fmt"
	"siuyin/present_go_hmac_auth/mac"
)

// 010 OMIT
func main() {
	key := []byte("ourSecret")
	msg := []byte("Brown Fox")
	msg2 := []byte("Brown Fox")
	myMAC := mac.MAC(msg, key)

	fmt.Printf("Authenticated: %t\n", mac.CheckMAC(msg2, myMAC, key))
}

// 020 OMIT
