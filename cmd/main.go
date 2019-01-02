package main

import (
	"fmt"
	"log"

	"github.com/gobuffalo/packd"
	packr "github.com/gobuffalo/packr/v2"
)

func main() {
	box := packr.New("myBox", "../templates")

	s, err := box.FindString("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(s)

	walk := func(fileName string, file packd.File) error {
		fmt.Println(fileName)
		return nil
	}

	if err := box.WalkPrefix("", walk); err != nil {
		log.Fatal(err)
	}
}
