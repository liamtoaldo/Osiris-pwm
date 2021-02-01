package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
)

//GetToSKey is a function that returns the global key for the terms of service text file
func GetToSKey() []byte { return []byte("TkQTu6t7rWmBFS2ZmAzX6YTfpz-evVmW") }

//EncryptStringInFile encrypts a string given the key and the file
func EncryptStringInFile(key []byte, text string, file string) {
	textInBytes := []byte(text)

	// generate a new aes cipher using our 32 byte long key
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())
	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	// here we encrypt our text using the Seal function
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
