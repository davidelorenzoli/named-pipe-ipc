package main

import (
	"flag"
	"fmt"
	"lab/named-pipe-ipc/pkg/ipc"
)

func main() {
	flag.Parse()
	namedPipe := flag.Args()[0]

	fmt.Println("Opening named pipe for writing")
	ipc.Write(namedPipe, []byte("hello"))
}
