package main

import (
	"fmt"
	"log"
	"secondlab/cmd"
	"secondlab/internal/xor"
	"secondlab/pkg/modulo"
)

func main() {
	conf := cmd.NewConfigs()
	switch conf.Task() {
	case 1:
		output, err := modulo.CipherXOR(conf.Text(), conf.Key())
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("The ciphered text:\n %v\t%08b\n", output, []rune(output))
	case 2:
		output, err := modulo.CipherMod32(conf.Text(), conf.Key())
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("The ciphered text: %v\n", output)
	case 3:
		output, err := xor.CipherString(conf.Text(), conf.Key())
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("The ciphered text: %v\n", output)
	default:
		log.Fatalln("Wrong task number")
	}
}
