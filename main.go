package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit()
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

}

func encryptHandle() {

}

func decryptHandle() {

}

func getPassword() {

}

func validatePassword() {

}

func validateFile() {
	
}
