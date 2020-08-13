package main

import (
	"fmt"
	"os"

	"github.com/richardlehane/mscfb"
)

// Parse parses a given hwp file
func Parse(f *os.File) error {
	doc, err := mscfb.New(f)
	if err != nil {
		return err
	}

	// weird c style for loop
	for entry, err := doc.Next(); err == nil; entry, err = doc.Next() {

		buf := make([]byte, 512)
		i, _ := doc.Read(buf)
		fmt.Println(entry.Name)
		if i > 0 {
			fmt.Println(buf[:i])
		}
	}

	// err from the for loop is caught here
	// EOF is also returned as an error
	return err
}
