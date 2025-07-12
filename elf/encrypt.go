package elf

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

func encrypt(key, plaintext []byte) (ciphertxt, nonce []byte, e error) {
	block, e := aes.NewCipher(key)
	if e != nil {
		return nil, nil, e
	}

	aesGCM, e := cipher.NewGCM(block)
	if e != nil {
		return nil, nil, e
	}

	nonce = make([]byte, aesGCM.NonceSize())
	rand.Read(nonce)

	ciphertxt = aesGCM.Seal(nil, nonce, plaintext, nil)
	return ciphertxt, nonce, nil
}

func decrypt(key, ciphertext, nonce []byte) ([]byte, error) {
	block, e := aes.NewCipher(key)
	if e != nil {
		return nil, e
	}

	aesGCM, e := cipher.NewGCM(block)
	if e != nil {
		return nil, e
	}

	plaintext, e := aesGCM.Open(nil, nonce, ciphertext, nil)
	if e != nil {
		return nil, e
	}

	return plaintext, nil
}
