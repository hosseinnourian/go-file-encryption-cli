package main

import (
	"fmt"
	"os"
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
func printHelp() {

}
func encryptHandle() {

}
func decryptHandle() {

}
func getPassword()  {
	
}