package main

import "bytes"

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
	//TODO implement me
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
	//TODO implement me
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
	//TODO implement me
	panic("implement me")
}

func newCaesarCipher() CaesarCipherInterface {
	return &CaesarCipher{}
}

// ANOTHER IMPL
//func cipher(input string, key int) string {
//	rotKey, alphabetSize := rune(3), rune(26)
//
//	runes := []rune(input)
//
//	for index, char := range runes {
//		switch key {
//		case -1:
//			if char >= 'a'+rotKey && char <= 'z' || char >= 'A'+rotKey && char <= 'Z' {
//				char = char - rotKey
//			}
//		}
//	}
//}
