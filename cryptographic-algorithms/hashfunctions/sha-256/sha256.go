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
	// output hash value: 78fe461fd72e68a1c7922b227a462e4b417e58bbccc81c6986d863069d8ae74e%
}
