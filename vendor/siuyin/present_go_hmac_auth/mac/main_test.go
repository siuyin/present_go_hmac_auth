package mac

import (
	"net/url"
	"testing"
	"time"
)

func TestCheckMAC(t *testing.T) {
	key := []byte("ourSecret")
	msg := []byte("Brown Fox")
	mac1 := MAC(msg, key)
	if !CheckMAC(msg, mac1, key) {
		t.Errorf("Authentication failed")
	}

	msg2 := []byte("Brown Fox.")
	if CheckMAC(msg2, mac1, key) {
		t.Errorf("Authentication failed")
	}
}
func TestValues(t *testing.T) {
	key := []byte("ourSecret")
	u := url.Values{}
	u.Add("a", "apple")
	u.Add("b", "boy")
	s := MACbs([]byte(u.Encode()), key)
	v := Values(u, "hmac", key)
	if s1 := v.Get("hmac"); s1 != s {
		t.Error("mac not matched")
	}

	if !CheckValues(v, "hmac", key) {
		t.Error("hmac not matched")
	}
}

func TestTimePeriod(t *testing.T) {
	now := time.Date(2016, 10, 28, 0, 0, 0, 0, time.Local)
	tp := TimePeriod(now, 30*time.Second)
	if tp != 49252800 {
		t.Errorf("bad timePeriod: %v", tp)
	}
}

func TestJunk(t *testing.T) {
	a := []byte("a")
	b := []byte("b")
	a = append(a, b...)
	ab := []byte("ab")
	if string(a) != string(ab) {
		t.Errorf("can't add")
	}
}
