package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	crand "crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"time"
)

//available characters for keys
var availableCharacters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

//GenerateKey is the main function for generating a new key (it's only used once if the user doesn't delete or corrupt the files where it is stored)
func GenerateKey() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 32)
	for i := range b {
		b[i] = availableCharacters[rand.Intn(len(availableCharacters))]
	}
	return string(b)
}

//GetToSKey is a function that returns the global key for the terms of service text file
func GetToSKey() []byte { return []byte("TkQTu6t7rWmBFS2ZmAzX6YTfpz-evVmW") }

//EncryptStringInFile encrypts a string given the key and the file
func EncryptStringInFile(key []byte, text string, file string) {
	textInBytes := []byte(text)

	//generates a new aes cipher using our 32 byte long key
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	//gcm is a mode of operation for symmetric key cryptographic block ciphers
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	//random sequence
	if _, err = io.ReadFull(crand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	//seal the new encrypted string in the file
	err = ioutil.WriteFile(file, gcm.Seal(nonce, nonce, textInBytes, nil), 0777)
	if err != nil {
		fmt.Println(err)
	}
}

//DecryptStringFromFile decrypts a string given the key and the file
func DecryptStringFromFile(key []byte, file string) string {
	ciphertext, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		fmt.Println(err)
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}

	return string(plaintext)
}
