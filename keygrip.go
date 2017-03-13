package gokeygrip

import (
	"crypto"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"hash"
	"log"
	"strings"
)

// Keygrip is our internal struct
type Keygrip struct {
	keys     []string
	hashFunc func() hash.Hash
	encoding Encoding
}

// Encoding is the preferred signing result format
type Encoding int

const (
	// HEX allows for hex encoded signing results
	HEX Encoding = iota
	// BASE64 allows for base64 encodded signing results
	BASE64
)

// New Keygrip based on the algorithm (SHA1, SHA256...), Encoding and provided keylist
func New(algorithm crypto.Hash, enc Encoding, keys ...string) *Keygrip {

	if !algorithm.Available() {
		log.Println("Binary wasn't compiled with algorithm")
		return nil
	}

	return &Keygrip{
		keys:     keys,
		hashFunc: algorithm.New,
		encoding: enc,
	}
}

// NewDefault returns a keygrip with and base64 encoding and provided keylist
func NewDefault(keys ...string) *Keygrip {

	return &Keygrip{
		keys:     keys,
		hashFunc: sha1.New,
		encoding: BASE64,
	}
}

// Index returns the index of the first matching key
func (kg *Keygrip) Index(data, digest string) int {
	if kg == nil {
		log.Fatalln("Keygrip not initialised")
	}
	for i, key := range kg.keys {
		if subtle.ConstantTimeCompare([]byte(digest), []byte(kg.SignWithKey(data, key))) == 1 {
			return i
		}
	}
	return -1
}

// SignWithKey returns the hash for the given key
// - default hashes are SHA1 HMACs in url-safe base64
func (kg *Keygrip) SignWithKey(data, key string) string {
	if kg == nil {
		log.Fatalln("Keygrip not initialised")
	}

	mac := hmac.New(kg.hashFunc, []byte(key))
	mac.Write([]byte(data))
	digest := mac.Sum(nil)
	cleanDigest := ""
	switch kg.encoding {
	case HEX:
		cleanDigest = hex.EncodeToString(digest)
	case BASE64:
		fallthrough
	default:
		cleanDigest = base64.URLEncoding.EncodeToString(digest)
	}

	keygripSafeFunc := func(r rune) rune {
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

// Sign returns the hash for the first key
// - default hashes are SHA1 HMACs in url-safe base64
func (kg *Keygrip) Sign(data string) string {
	return kg.SignWithKey(data, kg.keys[0])
}

// Verify returns the a boolean indicating a matched key
func (kg *Keygrip) Verify(data, digest string) bool {
	return kg.Index(data, digest) > -1
}
