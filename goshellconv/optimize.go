package goshellconv

func OptimizeShellcode(shellcode []byte) []byte {
	optimizedShellcode := []byte{}
	for _, b := range shellcode {
		if b != ',' {
			optimizedShellcode = append(optimizedShellcode, b)
		}
	}
	return optimizedShellcode
}
