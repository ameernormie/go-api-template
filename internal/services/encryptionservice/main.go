package encryptionservice

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
)

// Encrypt string to base64 crypto using AES
func Encrypt(text string) (string, error) {
	encKey := os.Getenv("ENC_KEY")
	encIv := os.Getenv("ENC_IV")
	key := []byte(encKey)
	iv := []byte(encIv)
	plainText := []byte(text)
	plainText, err := pad(plainText, aes.BlockSize)

	if err != nil {
		return "", fmt.Errorf(`plainText: "%s" has error`, plainText)
	}
	if len(plainText)%aes.BlockSize != 0 {
		err := fmt.Errorf(`plainText: "%s" has the wrong block size`, plainText)
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, len(plainText))

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText, plainText)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// Decrypt decrypts cipher text string into plain text string
func Decrypt(encrypted string) (string, error) {
	encKey := os.Getenv("ENC_KEY")
	encIv := os.Getenv("ENC_IV")
	key := []byte(encKey)
	iv := []byte(encIv)
	cipherText, _ := base64.StdEncoding.DecodeString(encrypted)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	if len(cipherText) < aes.BlockSize {
		err := fmt.Errorf(`cipherText too short: "%s"`, cipherText)
		return "", err
	}
	if len(cipherText)%aes.BlockSize != 0 {
		err := fmt.Errorf(`cipherText is not multiple of the block size: "%s"`, cipherText)
		return "", err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	cipherText, _ = unpad(cipherText, aes.BlockSize)
	return fmt.Sprintf("%s", cipherText), nil
}

func pad(buf []byte, size int) ([]byte, error) {
	bufLen := len(buf)
	padLen := size - bufLen%size
	padded := make([]byte, bufLen+padLen)
	copy(padded, buf)
	for i := 0; i < padLen; i++ {
		padded[bufLen+i] = byte(padLen)
	}
	return padded, nil
}

func unpad(padded []byte, size int) ([]byte, error) {
	if len(padded)%size != 0 {
		return nil, errors.New("padded value wasn't in correct size")
	}

	bufLen := len(padded) - int(padded[len(padded)-1])
	buf := make([]byte, bufLen)
	copy(buf, padded[:bufLen])
	return buf, nil
}
