package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/echenim/solawind-repl/services"
)

const READ string = "READ"
const WRITE string = "WRITE"
const DELETE string = "DELETE"
const START string = "START"
const COMMIT string = "COMMIT"
const ABORT string = "ABORT"
const QUIT string = "QUIT"

func main() {
	// s := "write a hello"
	// cm := strings.Fields(s)
	// fmt.Println(s[len(cm[0]):])
	header()
	sw := services.SolaWind{}
	sw.IsTransaction = false
	sw.Store = make(map[string]string)
	sw.Transactions = make(map[string]string)
	repl(&sw)
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

func repl(sw *services.SolaWind) {
	fmt.Print("solawind-repl> ")
	read := bufio.NewScanner(os.Stdin)
	read.Scan()
	data := read.Text()
	evaluate(data, sw)
	repl(sw)
}

func evaluate(s string, sw *services.SolaWind) {
	cmd := strings.Fields(s)
	switch strings.ToUpper(string(cmd[0])) {
	case WRITE:
		data := strings.Fields(s)
		takeOutCmd := s[len(data[0]):]
		sw.Write(takeOutCmd)
	case READ:
		data := strings.Fields(s)
		takeOutCmd := s[len(data[0]):]
		sw.Read(takeOutCmd)
	case DELETE:
		data := strings.Fields(s)
		takeOutCmd := s[len(data[0]):]
		sw.Delete(takeOutCmd)
	case COMMIT:
		sw.Commit()
	case ABORT:
		sw.Abort()
	case QUIT:
		sw.Quit()
	case START:
		sw.Start()
	default:
		fmt.Println("invalid command")
	}
}
