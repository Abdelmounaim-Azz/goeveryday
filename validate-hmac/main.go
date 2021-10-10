package main

import (
	"fmt"

	"github.com/alexellis/hmac"
)

func main() {
	input := []byte(`input message from API`)
	secret := []byte(`so secret`)
	digest := hmac.Sign(input, secret)
	fmt.Printf("Digest: %x\n", digest)
	err := hmac.Validate(input, fmt.Sprintf("sha1=%x", digest), string(secret))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Digest validated.\n")
}
