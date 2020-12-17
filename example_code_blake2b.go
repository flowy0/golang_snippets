package main
// from https://github.com/golang/go/issues/21644
import (
	"crypto"
	"crypto/hmac"
	"fmt"
	_ "golang.org/x/crypto/blake2b"
)

func main() {
	key := make([]byte, 32)
	msg := []byte("hello world")
	hash := hmac.New(crypto.BLAKE2b_256.New, key)
	hash.Write(msg)
	fmt.Println(hash.Sum(Nil))
}