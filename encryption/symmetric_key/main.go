package main

// todo: unfinished

import (
	"fmt"

	"golang.org/x/crypto/blowfish"
)

func main() {
	key := []byte("my key")
	cipher, err := blowfish.NewCipher(key)
	if err != nil {
		fmt.Println(err.Error())
	}
	src := []byte("hello\n\n")
	var enc [512]byte
	cipher.Encrypt(enc[0:], src)
	fmt.Println(enc)
	var decrypt [8]byte
	cipher.Decrypt(decrypt[0:], enc[0:])
	fmt.Println(decrypt)
}
