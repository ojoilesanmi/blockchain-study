package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	inputMessage := "sample input"
	messageDigest := sha256.New()
	messageDigest.Write([]byte(inputMessage))
	fmt.Printf("%x", messageDigest.Sum(nil))
}
