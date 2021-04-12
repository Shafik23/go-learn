package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	doneChan := time.After(30 * time.Second)
	echoChan := make(chan []byte)

	go readStdin(echoChan)

	for {
		select {
		case buf := <-echoChan:
			os.Stdout.Write(buf)
		case <-doneChan:
			fmt.Println("Timeout Expired - Goodbye!")
			os.Exit(0)
		}
	}

}

func readStdin(out chan<- []byte) error {
	for {
		data := make([]byte, 1024)
		l, _ := os.Stdin.Read(data)

		if l > 0 {
			out <- data
		}
	}
}
