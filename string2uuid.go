package main

import (
	"flag"

	"github.com/atotto/clipboard"
	"github.com/google/uuid"
)

func main() {
	copyToClipboard := flag.Bool("c", false, "copy to clipboard")
	flag.Parse()
	args := flag.Args()
	uuidString := args[0]
	id, err := uuid.Parse(uuidString)
	if err != nil {
		panic(err)
	}
	println(id.String())
	if *copyToClipboard {
		err = clipboard.WriteAll(id.String())
		if err != nil {
			panic(err)
		}
		println("copied to clipboard")
	}

}
