package main

import (
	"bytes"
	"file-encryption/encryptor"
	"fmt"
	"os"

	"golang.org/x/term"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}
	function := os.Args[1]

	switch function {
	case "help":
		printHelp()
	case "encrypt":
		encryptHandle()
	case "decrypt":
		decryptHandle()
	default:
		fmt.Println("please run [*]help, [*]encrypt, [*]decrypt")
		os.Exit(1)
	}
}

// done
func printHelp() {
	fmt.Println("file encryption")
	fmt.Println("simple file encryption for daily needs.")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("\tgo run . encrypt /path/to/your/file")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("")
	fmt.Println("\t encrypt \tEncrypts a file given a password")
	fmt.Println("\t decrypt\t Tries to decrypt a file using a password")
	fmt.Println("\t help\t\tDisplay help text")
	fmt.Println("")
}

func encryptHandle() {
	if len(os.Args) < 3 {
		fmt.Println("missing the path to the file.")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("File Not Found")
	}
	password := getPassword()
	fmt.Println("\nEncrypting...")

	// use modules
	encryptor.Encrypt(file, password)
	fmt.Println("\nFile sucessfully protected")
}
func decryptHandle() {
	if len(os.Args) < 3 {
		fmt.Println("missing the path to the file.")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("File Not Found")
	}
	fmt.Println("Enter the password")
	password, _ := term.ReadPassword(0)
	fmt.Println("\nDecrypting...")
	// use modules
	encryptor.Decrypt(file, password)
	fmt.Println("\nFile sucessfully decrypted")
}
func getPassword() []byte {
	fmt.Println("Enter Password...")
	password, _ := term.ReadPassword(0)
	fmt.Println("Confirm Password...")
	password2, _ := term.ReadPassword(0)

	if !validatePassword(password, password2) {
		fmt.Println("\nPassword Is Not Match.")
		return getPassword()
	}
	return password
}
func validatePassword(pass1, pass2 []byte) bool {
	if !bytes.Equal(pass1, pass2) {
		return false
	}
	return true
}
func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
