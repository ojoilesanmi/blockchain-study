package main

type CaesarInterface interface {
	Encrypt(input string, key int) string
	Decrypt(input string, key int) string
	Crack(input string) string
}

type Caesar struct {
	// empty struct
}

func (c Caesar) Encrypt(input string, key int) string {
	//TODO implement me
	panic("implement me")
}

func (c Caesar) Decrypt(input string, key int) string {
	//TODO implement me
	panic("implement me")
}

func (c Caesar) Crack(input string) string {
	//TODO implement me
	panic("implement me")
}

func NewCaeserCipher() CaesarInterface {
	return &Caesar{}
}

func main() {

}

// Watch 1899 series on Netflix.
