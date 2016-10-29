package main

import (
	"fmt"
	"siuyin/present_go_hmac_auth/mac"
	"time"
)

// 010 OMIT
func main() {
	key := []byte("ourSecret")
	tp := mac.TimePeriod(time.Now(), time.Second)
	tKeyN := mac.TimeKey(key, tp)   // now
	tKeyB := mac.TimeKey(key, tp-1) // before
	tKeyA := mac.TimeKey(key, tp+1) // after
	msg := []byte("Brown Fox")
	msg2 := []byte("Brown Fox")
	myMAC := mac.MAC(msg, tKeyN)

	fmt.Printf("%x\n", tKeyN)
	fmt.Printf("Authenticated (before): %t\n", mac.CheckMAC(msg2, myMAC, tKeyB))
	fmt.Printf("Authenticated (now)   : %t\n", mac.CheckMAC(msg2, myMAC, tKeyN))
	fmt.Printf("Authenticated (after) : %t\n", mac.CheckMAC(msg2, myMAC, tKeyA))
}

// 020 OMIT
