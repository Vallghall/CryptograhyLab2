package main

import (
	"fmt"
	"log"
	"secondlab/cmd"
	"secondlab/pkg/modulo"
	"secondlab/pkg/xor"
)

func main() {
	conf := cmd.NewConfigs()
	switch conf.Task() {
	case 1:
		output, err := xor.Cipher(conf.Text(), conf.Key())
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("The ciphered text: %v\n", output)
	case 2:
		output, err := modulo.Cipher(conf.Text(), conf.Key())
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("The ciphered text: %v\n", output)
	case 3:
		fallthrough
	default:
		log.Fatalln("Wrong task number")
	}
}
