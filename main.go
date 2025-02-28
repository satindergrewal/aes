package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"syscall/js"

	"golang.org/x/crypto/pbkdf2"
)

func deriveKey(passphrase string, salt []byte) ([]byte, []byte) {
	if salt == nil {
		salt = make([]byte, 8)
		rand.Read(salt)
	}
	return pbkdf2.Key([]byte(passphrase), salt, 1000, 32, sha256.New), salt
}

func encrypt(passphrase, plaintext string) string {
	key, salt := deriveKey(passphrase, nil)
	iv := make([]byte, 12)
	rand.Read(iv)
	b, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(b)
	data := aesgcm.Seal(nil, iv, []byte(plaintext), nil)
	return hex.EncodeToString(salt) + "-" + hex.EncodeToString(iv) + "-" + hex.EncodeToString(data)
}

func decrypt(passphrase, ciphertext string) string {
	arr := strings.Split(ciphertext, "-")
	if len(arr) != 3 {
		return ""
	}
	salt, _ := hex.DecodeString(arr[0])
	iv, _ := hex.DecodeString(arr[1])
	data, _ := hex.DecodeString(arr[2])
	key, _ := deriveKey(passphrase, salt)
	b, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(b)
	data, err := aesgcm.Open(nil, iv, data, nil)
	if err != nil {
		return ""
	}
	return string(data)
}

// encryptWrapper exposes encrypt to JavaScript
func encryptWrapper(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		return "Error: Expected 2 arguments (passphrase, plaintext)"
	}
	passphrase := args[0].String()
	plaintext := args[1].String()
	return encrypt(passphrase, plaintext)
}

// decryptWrapper exposes decrypt to JavaScript
func decryptWrapper(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		return "Error: Expected 2 arguments (passphrase, ciphertext)"
	}
	passphrase := args[0].String()
	ciphertext := args[1].String()
	result := decrypt(passphrase, ciphertext)
	if result == "" {
		return "Decryption failed: empty response. Check keys or ciphertext."
	}
	return result
}

func main() {
	// Register functions to be callable from JavaScript
	js.Global().Set("goEncrypt", js.FuncOf(encryptWrapper))
	js.Global().Set("goDecrypt", js.FuncOf(decryptWrapper))

	// Keep the program running (required for WASM)
	select {}
}