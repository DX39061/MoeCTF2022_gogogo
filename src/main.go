package main

import (
	"log"
)

func main() {
	if err := RouterInit(); err != nil {
		log.Panic(err)
	}
	err := Run()
	if err != nil {
		log.Panic(err)
	}
}
