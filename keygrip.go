package gokeygrip

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"log"
	"strings"
)

// Keygrip is our internal struct
type Keygrip struct {
	keys   []string
	hash   string
	cipher string
}

// New Keygrip based on the provided keylist.
func New(keys []string, algorithm string, encoding string) *Keygrip {
	switch strings.ToLower(algorithm) {
	case "sha256":

	case "sha1":

	}
	return &Keygrip{
		keys:   keys,
		hash:   "sha256",
		cipher: "aes-256-cbd",
	}
}

func (kg *Keygrip) Index() {
	log.Fatal("Not Implemented")
}

func (kg *Keygrip) Sign(data []byte) string {
	mac := hmac.New(sha1.New, []byte(kg.keys[0]))
	mac.Write(data)
	digest := mac.Sum(nil)
	cleanDigest := base64.URLEncoding.EncodeToString(digest)

	keygripSafeFunc := func(r rune) rune {
		// log.Printf("%c", r)
		switch {
		case r == '/':
			return '_'
		case r == '+':
			return '-'
		case r == '=':
			return -1
		}
		return r
	}
	return strings.Map(keygripSafeFunc, cleanDigest)
}

// func (kg *Keygrip) Verify() {
// 	mac := hmac.New(sha256.New, key)
// 	mac.Write(message)
// 	expectedMAC := mac.Sum(nil)
// 	return hmac.Equal(messageMAC, expectedMAC)
// }
