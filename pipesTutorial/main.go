package main

import (
	"fmt"
	"io"
	//"os"
	"os/exec"
)

func main() {
	// create two pipes
	pipe1Reader, pipe1Writer := io.Pipe()
	pipe2Reader, pipe2Writer := io.Pipe()

	// go routine to read from pipe1 and execute cat command
	go func() {
		// read the file path from pipe1
		filePathBytes, err := io.ReadAll(pipe1Reader)
		if err != nil {
			fmt.Println("Error reading from pipe1:", err)
			return
		}
		pipe1Reader.Close()
		filePath := string(filePathBytes)

		// execute the cat command on the file path and write the output to pipe2
		cmd := exec.Command("cat", filePath)
		cmd.Stdout = pipe2Writer
		if err := cmd.Run(); err != nil {
			fmt.Println("Error executing cat command:", err)
			return
		}
		pipe2Writer.Close()

	}()

	// get the file path from the user
	fmt.Print("Enter document path: ")
	filePath := ""
	fmt.Scanln(&filePath)

	// write the file path to pipe1
	if _, err := io.WriteString(pipe1Writer, filePath); err != nil {
		fmt.Println("Error writing to pipe1:", err)
		return
	}
	pipe1Writer.Close()

	// read the output of the cat command from pipe2
	outputBytes, err := io.ReadAll(pipe2Reader)
	if err != nil {
		fmt.Println("Error reading from pipe2:", err)
		return
	}
	fmt.Println(string(outputBytes))

	pipe2Reader.Close()
}