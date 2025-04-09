package goshellconv

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadFile(fileName string) ([]byte, error) {
	fileData, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("error reading the file: %w", err)
	}
	return fileData, nil
}

func GenerateShellcode(data []byte) []byte {
	var shellcode bytes.Buffer
	for i := 0; i < len(data); i++ {
		fmt.Fprintf(&shellcode, "0x%02X,", data[i])
	}
	return shellcode.Bytes()
}

func RemoveExtension(fileName string) string {
	return strings.TrimSuffix(fileName, ".exe")
}
