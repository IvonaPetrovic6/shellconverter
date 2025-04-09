package goshellconv

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

func EncryptShellcode(shellcode []byte, passphrase string) ([]byte, error) {
	key := []byte(passphrase)
	for len(key) < 16 {
		key = append(key, 0)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("error creating AES cipher: %w", err)
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, fmt.Errorf("error generating IV: %w", err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	encryptedShellcode := make([]byte, len(shellcode))
	stream.XORKeyStream(encryptedShellcode, shellcode)

	encryptedShellcode = append(iv, encryptedShellcode...)
	return encryptedShellcode, nil
}
