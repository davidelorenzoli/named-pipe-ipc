package main

import (
	"flag"
	"lab/named-pipe-ipc/pkg/ipc"
	"log"
	"os"
	"os/exec"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	var executablePath string

	flag.StringVar(&executablePath, "executablePath", "./spawned_process", "path to spawned_process")
	flag.Parse()

	namedPipe := ipc.CreateNamedPipe()

	waitGroup.Add(1)

	go execute(executablePath, namedPipe)

	go ipc.Read(namedPipe)

	waitGroup.Wait()

	if err := os.Remove(namedPipe); err != nil {
		log.Printf("Failed to delete named pipe %s. Error: %s", namedPipe, err.Error())
	} else {
		log.Printf("Successfully deleted named pipe %s", namedPipe)
	}
}

func execute(executablePath string, namedPipe string) {
	cmd := exec.Command(executablePath, namedPipe)
	// Just to forward the stdout
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		log.Printf("Failed to execute %s. Error: %s", executablePath, err.Error())
	}

	defer waitGroup.Done()
}
