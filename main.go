package main

import (
	"encoding/base64"
	"io"
	"log"
	"os"
)

func main() {
	// read input from stdin
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Printf("Error reading from stdin: %v", err)
		return
	}
	if isDecode(os.Args[1:]) {
		output, err := base64.RawStdEncoding.DecodeString(string(data))
		if err != nil {
			log.Printf("Error decoding base64: %v", err)
			return
		}
		os.Stdout.Write(output)
		return
	}

	output := base64.RawStdEncoding.EncodeToString(data)
	os.Stdout.WriteString(output)
}

func isDecode(args []string) bool {
	for _, x := range args {
		if x == "-d" || x == "-D" || x == "--decode" {
			return true
		}
	}
	return false
}
