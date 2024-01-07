package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/G1A021037-ANJASFEDO/go-file-encryption/filecrypt"
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
		fmt.Println("Run encrypt or decrypt")
	}
}

func printHelp() {
	fmt.Println("file encryption")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("\t go run . encrypt (file path)")
	fmt.Println("")
	fmt.Println("Command:")
	fmt.Println("\t encrypt\tEncrypts the file by give it a password")
	fmt.Println("\t decrypt\tDecrypts the file by using the password")
	fmt.Println("\t help\t\tDisplay help text")
	fmt.Println("")
}

func encryptHandle() {
	if len(os.Args) < 3 {
		println("Missing the path to file, For more info type go run . help")
	}

	file := os.Args[2]

	if !validateFile(file) {
		panic("File not Found")
	}

	password := getPassword()

	fmt.Println("\n Encrypting...")
	filecrypt.Encrypt(file, password)
	fmt.Println("\n file successfuly protected")
}

func decryptHandle() {
	if len(os.Args) < 3 {
		println("Missing the path to file, For more info type go run . help")
	}

	file := os.Args[2]

	if !validateFile(file) {
		panic("File not Found")
	}

	fmt.Print("Enter Password: ")
	password, _ := term.ReadPassword(0)
	fmt.Println("\n Decrypting...")

	filecrypt.Decrypt(file, password)
	fmt.Println("\n file successfuly decrypted")
}

func getPassword() []byte {
	fmt.Print("Enter Password")
	password1, _ := term.ReadPassword(0)
	fmt.Print("\nConfirm Password: ")
	password2, _ := term.ReadPassword(0)

	if !validatePassword(password1, password2) {
		fmt.Printf("\n Password do not Match, please try again")
		return getPassword()
	}
	return password1
}

func validatePassword(password1 []byte, password2 []byte) bool {
	if !bytes.Equal(password1, password2) {
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
