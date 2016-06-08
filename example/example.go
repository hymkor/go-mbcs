package main

import (
	"fmt"
	"os"

	mbcs ".."
)

func main() {
	var ansi []byte
	var ansi_err error

	ansi, ansi_err = mbcs.UtoA("UTF8文字列")
	if ansi_err != nil {
		fmt.Fprintln(os.Stderr, ansi_err)
		return
	}

	var utf8 string
	var utf8_err error

	utf8, utf8_err = mbcs.AtoU(ansi)
	if utf8_err != nil {
		fmt.Fprintln(os.Stderr, utf8_err)
		return
	}
	fmt.Printf("Ok: %s\n", utf8)
}

// vim:set fenc=utf8:
