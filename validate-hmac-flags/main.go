package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/alexellis/hmac"
)

func main() {
	var input string
	var secret string
	flag.StringVar(&input, "message", "", "message to create a digest from")
	flag.StringVar(&secret, "secret", "", "secret for the digest")
	flag.Parse()
	if len(strings.TrimSpace(secret)) == 0 {
		panic("--secret is required")
	}
	fmt.Printf("Computing hash for: %q\nSecret: %q\n", input, secret)
	digest := hmac.Sign([]byte(input), []byte(secret))
	fmt.Printf("Digest: %x\n", digest)
}
