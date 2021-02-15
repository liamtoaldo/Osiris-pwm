package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"

	"os"
)

//available characters for keys
var availableCharacters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!-$£%&/()=?^'*@°#§")

//GenerateKey is the main function for generating a new key (it's only used once if the user doesn't delete or corrupt the files where it is stored)
func GenerateKey() string {
	b := make([]rune, 32)
	for i := range b {
		//a random position generated with the crypto library, so it's way more secure and
		randomPos, _ := rand.Int(rand.Reader, big.NewInt(int64(len(availableCharacters)-2))) //minus 2 because otherwise the key length will be 33, so invalid
		b[i] = availableCharacters[randomPos.Int64()]
	}
	return string(b)
}

//GetToSKey is a function that returns the global key for the settings files
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
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(file, gcm.Seal(nonce, nonce, textInBytes, nil), 770)
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

//GetGlobalKey returns the global key that encrypts and decrypts the global
func GetGlobalKey() []byte { return []byte("q1ozLRcb3YrlpvxZqUE7LviP5MFwnT5w") }

//EncryptDataStringInFile encrypts a string given the key and the file to a data file (one containing services, usernames and passwords.)
func EncryptDataStringInFile(key []byte, text string, file string) {
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
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	//seal the new encrypted string in the file
	for i := 0; i < 999; i++ {
		file += fmt.Sprint(i)
		println(file)
		if _, err := os.Stat(file); os.IsNotExist(err) {
			break
		}
	}

	err = ioutil.WriteFile(file, gcm.Seal(nonce, nonce, textInBytes, nil), 770)
	if err != nil {
		fmt.Println(err)
	}
}

//DecryptDataStringFromFile decrypts a string given the key and the file with DATA
func DecryptDataStringFromFile(key []byte, file string) string {

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

	for i := 0; i < 999; i++ {
		file += fmt.Sprint(i)
		if _, err := os.Stat(file); os.IsNotExist(err) {
			break
		}
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}

	return string(plaintext)
}
