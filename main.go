package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var msg = `
Usage: goodhangul FILENAME [OPTION]
A converter for hwp files

OPTIONS:
  xml	parse and output the file to xml
  pdf	parse and output the file to pdf
`

// bit of a hack. Stdandard flag lib doesn't allow flag.Parse(os.Args[2]). You need a subcommand to do so.
var optionCmd = flag.NewFlagSet("", flag.ExitOnError)

func main() {
	//check if enough arguments were given
	if len(os.Args) < 2 {
		fmt.Println(msg)
		os.Exit(1)
	}

	err := optionCmd.Parse(os.Args[0:])
	if err != nil {
		fmt.Println(msg)
		os.Exit(1)
	}

	fileName := os.Args[1]

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(msg)
		os.Exit(1)
	}
	err = Parse(f)
	if err != nil && err != io.EOF {
		fmt.Println(msg)
		os.Exit(1)
	}
}
