package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"io/ioutil"
)

func main() {
	// Create tcp address
	tcpAddr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:9999")
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}

	// Create listener
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}

	// Accept connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error: ", err.Error())
			continue
		}
		go handleClient(conn)
	}

}

//Function that handles client connection, receives the contents of a file writes it in another file
func handleClient(conn net.Conn) {
	defer conn.Close()
	
	// Create buffer
	var buf [512]byte

	// Read data
	n, err := conn.Read(buf[0:])
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	// Convert data to string
	message := string(buf[0:n])

	//Add empty line at the end of the message
	message += "\n"

	// Split message by empty lines
	lines := strings.Split(message, "\n")

	// Create file
	file, err := os.Create("resultado.txt")
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	defer file.Close()

	// Writes content in file
	for _, line := range lines {
		file.WriteString(line + "\n")
	}

	// Get words
	words := getWords("resultado.txt")
	// Get palindromes
	palindromes := getPalindromes(words)
	// empty file
	file.Truncate(0)

	// Send palindromes to client
	for _, palindrome := range palindromes {
		if palindrome == "" {
			continue
		}
		file.WriteString(palindrome + "\n")
	}

	// Read file
	content, err := ioutil.ReadFile("resultado.txt")
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	// Send data
	conn.Write(content)
}



//Function that returns all the words of a file in a slice
func getWords(file string) []string {
	// Read file
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return nil
	}

	// Convert content to string
	text := string(content)

	// Split content by lines
	lines := strings.Split(text, "\n")

	// Split lines by words
	var words []string
	for _, line := range lines {
		words = append(words, strings.Split(line, " ")...)
	}

	return words
}

//Function that receives a slice of words and returns a slice of palindromes
func getPalindromes(words []string) []string {
	var palindromes []string
	for _, word := range words {
		if esPalindroma(word) {
			palindromes = append(palindromes, word)
		}
	}
	return palindromes
}

//Function that receives a word and returns true if it is a palindrome
func esPalindroma(word string) bool {
	// Convert word to lowercase
	word = strings.ToLower(word)

	// Convert word to slice of strings
	letters := strings.Split(word, "")

	fmt.Println(letters)
	// Reverse letters
	var reverse []string
	for i := len(letters) - 1; i >= 0; i-- {
		reverse = append(reverse, letters[i])
	}

	// Convert reverse to string
	reverseWord := strings.Join(reverse, "")

	// Compare words
	return word == reverseWord
}
