package main

import (
	"fmt"
	"log"

	packr "github.com/gobuffalo/packr/v2"
)

func main() {
	box := packr.New("myBox", "../templates")

	s, err := box.FindString("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(s)
}
