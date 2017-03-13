package gokeygrip

import (
	"crypto"
	"log"
	"os"
	"testing"

	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"
)

var logger *log.Logger

func TestMain(m *testing.M) {
	logger = log.New(os.Stderr, "testing: ", log.LstdFlags|log.Lshortfile)
	os.Exit(m.Run())
}

func Test(t *testing.T) {
	foo := New(crypto.SHA1, BASE64, "06ae66fdc6c2faf5a401b70e0bf885cb")
	hash := foo.Sign("bieberschnitzel")
	logger.Printf("%s", hash)
	if hash != "qcnaleT6EOQqYoxYrness05NAW8" {
		t.Fatal("Hash didn't match - got", hash, "expected qcnaleT6EOQqYoxYrness05NAW8")
	}

	// .index returns the index of the first matching key
	index := foo.Index("bieberschnitzel", hash)
	if index != 0 {
		t.Fatal("Index didn't match - got", index, "expected 0")
	}

	// .verify returns the a boolean indicating a matched key
	matched := foo.Verify("bieberschnitzel", hash)
	if matched != true {
		t.Fatal("matched didn't match - got", matched, "expected true")
	}

	index = foo.Index("bieberschnitzel", "o_O")
	if index != -1 {
		t.Fatal("Index didn't match - got", index, "expected -1")
	}
}
