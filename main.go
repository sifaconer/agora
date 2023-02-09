package main

import (
	"agora/cmd"
	"log"
)

func main() {
	if err := cmd.NewServer().Run(); err != nil {
		log.Fatal(err)
	}
}
