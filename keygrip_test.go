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
	foo := New([]string{"06ae66fdc6c2faf5a401b70e0bf885cb"})
	logger.Printf("%s", foo.Sign([]byte("bieberschnitzel")))

	log.Println("-----")
	{
		// foo := "abcdefg"
		// bar := hmac.New(sha256.New, []byte(foo))
		// bar.Write([]byte("I love cupcakes"))
		// log.Println(hex.EncodeToString(bar.Sum(nil)))
	}
}

// const secret = 'abcdefg';
// const hash = crypto.createHmac('sha256', secret)
//                    .update('I love cupcakes')
//                    .digest('hex');
// console.log(hash);
// // Prints:
// //   c0fa1bc00531bd78ef38c628449c5102aeabd49b5dc3a2a516ea6ea959d6658e
