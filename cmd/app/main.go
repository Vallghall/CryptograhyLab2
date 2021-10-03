package main

import (
	"fmt"
	"log"
	"os"
	"secondlab/internal/xor"
)

func main() {
	output, err := xor.CipherString(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("The ciphered text: %v\n", output)
}
