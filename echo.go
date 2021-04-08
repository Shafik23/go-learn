package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	go echo(os.Stdin, os.Stdout)

	time.Sleep(30 * time.Second)
	fmt.Println("Timed out!")
	os.Exit(0)
}

func echo(input io.Reader, output io.Writer) {
	io.Copy(output, input)
}
