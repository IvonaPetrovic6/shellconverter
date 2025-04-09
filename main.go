package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"goexetoshell/goshellconv"
)

func printUsage() {
	fmt.Println("Usage: goshellconv -i <input_file> -o <output_file> -crypt <passphrase> -opt")
	fmt.Println()
	fmt.Println("Arguments:")
	fmt.Println("-i <input_file>    Input file (e.g., .exe or .dll)")
	fmt.Println("-o <output_file>   Output file (e.g., .bin)")
	fmt.Println("-crypt <passphrase> Optional AES encryption passphrase")
	fmt.Println("-opt                Optional flag to optimize shellcode (removes commas)")
	fmt.Println()
	fmt.Println("Example:")
	fmt.Println("goshellconv -i input.exe -o output.bin -crypt=\"mysecurepassword\" -opt")
}

func main() {
	inputFile := flag.String("i", "", "Input file (.exe or .dll)")
	outputFile := flag.String("o", "", "Output file (.bin)")
	passphrase := flag.String("crypt", "", "Encryption passphrase (AES)")
	optimize := flag.Bool("opt", false, "Optimize the shellcode")
	flag.Parse()

	if *inputFile == "" || *outputFile == "" {
		printUsage()
		os.Exit(1)
	}

	fmt.Println("[+] received file conversion in progress..")

	shellcode, err := goshellconv.ConvertToShellcode(*inputFile, *passphrase, *optimize)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[#] converted 1/2")

	err = os.WriteFile(*outputFile, shellcode, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[#] converted 1/2")
	fmt.Println("[+] success convertation to shellcode x64.")
}
