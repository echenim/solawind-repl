package main

import "fmt"

const READ string = "READ"
const WRITE string = "WRITE"
const DELETE string = "DELETE"
const START string = "START"
const COMMIT string = "COMMIT"
const ABORT string = "ABORT"
const QUIT string = "QUIT"

func main() {
	header()
}

func header() {

	fmt.Println("solawind-repl> List of commands: ")
	fmt.Println("solawind-repl> WRITE     - Stores val in key. ")
	fmt.Println("solawind-repl> START     - Start a transaction ")
	fmt.Println("solawind-repl> COMMIT    - Commit a transaction ")
	fmt.Println("solawind-repl> DELETE    - Removes a key from store ")
	fmt.Println("solawind-repl> CLS       - Clear the Terminal Screen ")
	fmt.Println("solawind-repl> READ      - Reads and prints, to stdout, the val associated with key ")
	fmt.Println("solawind-repl> ABORT     - Abort a transaction. All actions in the current transaction are discarded ")
	fmt.Println("solawind-repl> QUIT      - Exit the REPL cleanly")
}
