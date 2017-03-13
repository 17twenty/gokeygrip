package gokeygrip

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
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

// SetHash sets the hash type
func (kg *Keygrip) SetHash() {

}

// SetHash sets the hash type
func (kg *Keygrip) GetHash() {

}

// SetCipher sets the cipher type
func (kg *Keygrip) SetCipher() {

}

// GetCipher gets the cipher type
func (kg *Keygrip) GetCipher() {

}

func (kg *Keygrip) Encrypt() []byte {
	return []byte{}
}

func (kg *Keygrip) Decrypt([]byte) (error, []byte) {
	return errors.New("Couldn't decrypt"), []byte{}
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
