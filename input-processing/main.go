package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("SP// Backend Developer Test - Input Processing")
	fmt.Println()

	ch := make(chan string)

	go SearchForErrors(ch, os.Stdin)

	for message := range ch {
		fmt.Println(message)

	}

}
func SearchForErrors(c chan string, r io.Reader) {
	// Scanner was recommended by the standard library
	// However it has a max token size of 64*1024
	// Since we must support larger inputs Reader must be used instead.
	defer close(c)

	// Don't read anything bigger than 1GB
	// Can be tweaked based on performance needs.
	reader := bufio.NewReaderSize(r, 64*64*64*4096)

	for {
		token, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				c <- ""
				break
			}
			fmt.Println(err)
		}

		if strings.Contains(token, errorMsg) {
			c <- token
		}
	}
}
