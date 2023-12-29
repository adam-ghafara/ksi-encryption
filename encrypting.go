package ksi

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

// encrypt encrypts data using a 256-bit key
func encrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	b := base64.StdEncoding.EncodeToString(data)
	ciphertext := make([]byte, aes.BlockSize+len(b))

	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))

	return ciphertext, nil
}

// decrypt decrypts data using a 256-bit key
func decrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(data) < aes.BlockSize {
		return nil, errors.New("Ciphertext too short.")
	}

	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(data, data)

	dbuf := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	n, err := base64.StdEncoding.Decode(dbuf, data)
	if err != nil {
		return nil, err
	}

	return dbuf[:n], nil
}

func encryptProcess() {
	// Sample data
	data := []byte("Sample data for encryption")

	// AES-256 Key
	key := []byte("0123456789012345678901234567890")

	// Encrypting the data
	ciphertext, err := encrypt(data, key)
	if err != nil {
		fmt.Println("Encryption Error:", err)
		return
	}

	// Printing encrypted data
	fmt.Println("Encrypted Data:", string(ciphertext))
}

func decryptProcess() {
	// Sample data
	data := []byte("Sample data for decryption")

	// AES-256 Key
	key := []byte("0123456789012345678901234567890")

	// Decrypting the data
	decryptedtext, err := decrypt(data, key)
	if err != nil {
		fmt.Println("Decryption Error:", err)
		return
	}

	// Printing decrypted data
	fmt.Println("Decrypted Data:", string(decryptedtext))
}
