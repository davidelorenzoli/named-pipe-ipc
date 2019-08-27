package ipc

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"syscall"
)

func CreateNamedPipe() string {
	tmpDir, _ := ioutil.TempDir("", "named-pipes")
	// Create named pipe
	namedPipe := filepath.Join(tmpDir, "stdout")

	if err := syscall.Mkfifo(namedPipe, 0600); err != nil {
		log.Printf("Failed to create named pipe %s. Error: %s", tmpDir, err.Error())
	} else {
		log.Printf("Create named pipe %s", tmpDir)
	}

	return namedPipe
}

func Read(namedPipe string) {
	// Open named pipe for reading
	log.Println("Opening named pipe for reading")
	stdout, _ := os.OpenFile(namedPipe, os.O_RDONLY, 0600)
	log.Println("Reading")

	var buff bytes.Buffer
	log.Println("Waiting for someone to write something")

	if _, err := io.Copy(&buff, stdout); err != nil {
		log.Printf("Failed to read pipe. Error: %s", err)
	}

	if err := stdout.Close(); err != nil {
		log.Printf("Failed to close stream. Error: %s", err)
	}

	log.Printf("Data: %s\n", buff.String())
}

func Write(namedPipe string, content []byte) {
	stdout, _ := os.OpenFile(namedPipe, os.O_RDWR, 0600)
	fmt.Println("Writing")
	if _, err := stdout.Write(content); err != nil {
		log.Printf("Error writing bytes: %s", err.Error())
	}
	if err := stdout.Close(); err != nil {
		log.Printf("Error closing writer: %s", err.Error())
	}
}
