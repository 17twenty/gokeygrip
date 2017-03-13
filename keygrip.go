package gokeygrip

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/subtle"
	"encoding/base64"
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

func (kg *Keygrip) Index(data []byte, digest string) int {
	for i, key := range kg.keys {
		if subtle.ConstantTimeCompare([]byte(digest), []byte(kg.SignWithKey(data, key))) == 1 {
			return i
		}
	}
	return -1
}

func (kg *Keygrip) SignWithKey(data []byte, key string) string {
	mac := hmac.New(sha1.New, []byte(key))
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

func (kg *Keygrip) Sign(data []byte) string {
	return kg.SignWithKey(data, kg.keys[0])
}

func (kg *Keygrip) Verify(data []byte, digest string) bool {
	return kg.Index(data, digest) > -1
}
