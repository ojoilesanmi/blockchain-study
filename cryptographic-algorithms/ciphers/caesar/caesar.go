package main

import (
	"bytes"
	"fmt"
)

type CaesarCipherInterface interface {
	Encrypt(input string, rotKey int) string
	Decrypt(input string, rotKey int) string
	Crack(input string) string
}

type CaesarCipher struct {
	// an empty struct
}

/* NOTE:
-> ASCII CODE:
-> 97-122 are lowercase letters
-> 65-90 are uppercase letters
-> a cipher letter x = (originalCharacter + rotKey) % 26
-> for example, letter H = (8+4)%26
*/

/* STUDY THESE:
 	-> runes
	-> bytes
	-> bitwise operations with Go
	-> interfaces, methods, structs, types

*/

func getCharacterString(alphabetPosition int, asciiValue int32) string {
	var isLowerCase = asciiValue >= 97
	if isLowerCase {
		return string('a' - 1 + alphabetPosition)
	}
	return string('A' - 1 + alphabetPosition)
}

func getCharacterPosition(asciiValue int32) int {
	asciiValue |= 32 // bitwise OR operation
	return int(asciiValue) - 96
}

func (c *CaesarCipher) Encrypt(input string, rotKey int) string {
	var cipherText bytes.Buffer
	for _, character := range input {
		if character == 32 {
			cipherText.WriteByte(byte(character))
			continue
		}
		var characterPosition = getCharacterPosition(character)
		var cipherCharacter = getCharacterString((characterPosition+rotKey)%26, character)
		cipherText.WriteString(cipherCharacter)
	}
	return cipherText.String()
}

func (c *CaesarCipher) Decrypt(input string, rotKey int) string {
	var cipherText bytes.Buffer

	for _, character := range input {
		if character == 32 {
			cipherText.WriteByte(byte(character))
			continue
		}
		var characterPosition = getCharacterPosition(character)
		var modulusOf26 = (characterPosition - rotKey) % 26
		if modulusOf26 < 0 {
			modulusOf26 += 26
		}
		var originalCharacter = getCharacterString(modulusOf26, character)
		cipherText.WriteString(originalCharacter)
	}
	return cipherText.String()
}

func (c *CaesarCipher) Crack(input string) string {
	var characterFrequency = []float32{.0820, .0150, .0280, .0430, .1270, .0220, .0200, .0610, .0700, .0015, .0080, .0400, .0240,
		.0670, .0750, .0190, .0010, .0600, .0630, .0910, .0280, .0010, .0240, .0015, .0200, .0007,
	}
	var characterSums [26]float32

	for _, character := range input {
		if character == 32 {
			continue
		}
		var characterPosition = getCharacterPosition(character)
		characterSums[characterPosition-1] += characterFrequency[characterPosition-1]
	}
	return "not implemented..."
}

func newCaesarCipher() CaesarCipherInterface {
	return &CaesarCipher{}
}

func main() {
	test1 := newCaesarCipher().Encrypt("Apple", 3)
	fmt.Printf("The encrypted form of the word 'Apple' using Rotation Key 3 is: %s", test1)
}
