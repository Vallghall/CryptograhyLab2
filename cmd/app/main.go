package main

import (
	"log"
	"secondlab/cmd"
)

func main() {
	conf := cmd.NewConfigs()
	switch conf.Task() {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fallthrough
	default:
		log.Fatalln("Wrong task number")
	}
}
