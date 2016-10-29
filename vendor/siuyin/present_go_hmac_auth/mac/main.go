//005 OMIT

// Package mac implements HMAC with sha256 hash.
package mac

//006 OMIT

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"encoding/base64"
	"encoding/binary"
	"net/url"
	"time"
)

var key = []byte("ourSecret")

//010 OMIT

// MAC generate a Message Authentication Code from the message and key.
func MAC(message, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	return mac.Sum(nil)
}

//020 OMIT
//030 OMIT

// CheckMAC reports whether messageMAC is a valid HMAC tag for message.
func CheckMAC(message, messageMAC, key []byte) bool {
	expectedMAC := MAC(message, key)
	return hmac.Equal(messageMAC, expectedMAC)
}

//040 OMIT

//050 OMIT

// MACbs returns the base64 encoded MAC as a string.
func MACbs(message, key []byte) string {
	return base64.URLEncoding.EncodeToString(MAC(message, key))
}

//060 OMIT

//070 OMIT

//MACsb converts a bas64 encoded MAC as a []byte.
func MACsb(s string) ([]byte, error) {
	b, err := base64.URLEncoding.DecodeString(s)
	return b, err
}

//080 OMIT

//090 OMIT

// Values add key called name with HMAC value hex encoded.
func Values(u url.Values, name string, key []byte) url.Values {
	s := MACbs([]byte(u.Encode()), key)
	u.Set(name, s)
	return u
}

//100 OMIT

//110 OMIT

// CheckValues checks if hmac key named name matches.
func CheckValues(u url.Values, name string, key []byte) bool {
	s := u.Get(name)
	u.Del(name)
	s2 := MACbs([]byte(u.Encode()), key)
	return s == s2
}

//120 OMIT

// TimePeriod returns the number of time periods since Unix start of time (epoch).
func TimePeriod(now time.Time, step time.Duration) int64 {
	return now.UnixNano() / int64(step)
}
func toBytes(i int64) []byte {
	buf := &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, uint64(i))
	return buf.Bytes()
}

// TimeKey return a time-based key
func TimeKey(key []byte, tp int64) []byte {
	t := toBytes(tp)
	dst := make([]byte, 0, len(key)+len(t))
	dst = append(key, t...)
	return dst
}

// RandID returns a base64 encoded string of len random bytes.
func RandID(len int) string {
	buf := make([]byte, len, len)
	rand.Read(buf)
	return base32.StdEncoding.EncodeToString(buf)
}
