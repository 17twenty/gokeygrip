package gokeygrip

import (
	"log"
	"os"
	"testing"
)

var logger *log.Logger

func TestMain(m *testing.M) {
	logger = log.New(os.Stderr, "gotifications: ", log.LstdFlags|log.Lshortfile)
	os.Exit(m.Run())
}

func Test(t *testing.T) {
	foo := New([]string{"06ae66fdc6c2faf5a401b70e0bf885cb"}, "", "")
	hash := foo.Sign([]byte("bieberschnitzel"))
	logger.Printf("%s", hash)

	log.Println("-----")
	{
		// foo := "abcdefg"
		// bar := hmac.New(sha256.New, []byte(foo))
		// bar.Write([]byte("I love cupcakes"))
		// log.Println(hex.EncodeToString(bar.Sum(nil)))

		// .index returns the index of the first matching key
		index := foo.Index([]byte("bieberschnitzel"), hash)
		log.Println(index, 0)

		// .verify returns the a boolean indicating a matched key
		matched := foo.Verify([]byte("bieberschnitzel"), hash)
		log.Println(matched)

		index = foo.Index([]byte("bieberschnitzel"), "o_O")
		log.Println(index, -1)
	}
}

// const secret = 'abcdefg';
// const hash = crypto.createHmac('sha256', secret)
//                    .update('I love cupcakes')
//                    .digest('hex');
// console.log(hash);
// // Prints:
// //   c0fa1bc00531bd78ef38c628449c5102aeabd49b5dc3a2a516ea6ea959d6658e
