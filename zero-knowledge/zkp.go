package main

import (
	"fmt"
	// "github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto" // https://pkg.go.dev/github.com/ethereum/go-ethereum/crypto
	"math/big"
	// "crypto/rand"
)

// generateZKProof is the prover function that generates a proof that it knows the secretNumber.
func generateZKProof(secretNumber *big.Int) (*big.Int, *big.Int) {
	// randomNumber generates a cryptographically secure random *big.Int with 256 bits of entropy.
	randomNumber := crypto.RandBigInt()
	//TODO
	// - why is RandBigInt from go-ethereum not working?

	// concatHash computes the hash of the concatenation of secretNumber and randomNumber.
	concatHash := crypto.Keccak256Hash(append(secretNumber.Bytes(), randomNumber.Bytes()...))

	//compute the value: randomNumber - concatHash * secretNumber
	solution := new(big.Int).Sub(randomNumber, new(big.Int).Mul(concatHash, secretNumber))
	//TODO
	// - how to convert the Keccak256Hash output to big.Int, or replace the hashing algorithm.
	return concatHash, solution
}

// verifyZKProof is the verifier function that verifies the proof of knowledge of secretNumber is valid.
func verifyZKProof(secretNumber, concatHash, solution *big.Int) bool {
	// compute the hash of the concatenation of secretNumber and randomNumber.
	verifiedConcatHash := crypto.Keccak256Hash(append(secretNumber.Bytes(), new(big.Int).Add(secretNumber, concatHash).Bytes()...))

	// verify that the computed hash matches the received hash by comparing. (i.e. that verifiedConcatHash = concatHash).
	return verifiedConcatHash.Cmp(concatHash) == 0
	//TODO
	// - why is the Cmp() method that is available to all big.Ints not working?
}

func main() {
	// Assuming our secretNumber is 100.
	secretNumber := big.NewInt(100)
	// generate a proof of knowledge.
	testConcatHash, testSecretNumber := generateZKProof(secretNumber)
	// verify the proof of knowledge.
	if verifyZKProof(secretNumber, testConcatHash, testSecretNumber) {
		fmt.Println("The proof is valid.")
	} else {
		fmt.Println("The proof is invalid.")
	}
}
