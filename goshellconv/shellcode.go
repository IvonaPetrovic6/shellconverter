package goshellconv

import (
	"fmt"
	"io/ioutil"
)

func ConvertToShellcode(inputFile string, passphrase string, optimize bool) ([]byte, error) {
	fileData, err := ReadFile(inputFile)
	if err != nil {
		return nil, err
	}

	shellcode := GenerateShellcode(fileData)

	if passphrase != "" {
		encryptedShellcode, err := EncryptShellcode(shellcode, passphrase)
		if err != nil {
			return nil, err
		}
		shellcode = encryptedShellcode
	}

	if optimize {
		shellcode = OptimizeShellcode(shellcode)
	}

	if passphrase != "" {
		err := ioutil.WriteFile("passcode.txt", []byte("Do not forget your password: "+passphrase), 0644)
		if err != nil {
			return nil, fmt.Errorf("error writing passcode file: %w", err)
		}
	}

	return shellcode, nil
}
