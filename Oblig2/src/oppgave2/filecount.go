package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"io/ioutil"
	"strings"

)
var antall []int
var rune []string

func main() {

	filnavn := os.Args[1]
	filnavn, err := getFilnavn()

	b, err := ioutil.ReadFile(filnavn)
	if err != nil {
		fmt.Print(err)
	}

	if err != nil {
		logError(err)
		os.Exit(1)
	}

	lines, err := countLinesInFile(filnavn)
	if err != nil {
		logError(err)
		os.Exit(2)
	}

	fmt.Println("\nInformation about <" + filnavn + ">:\n")
	fmt.Printf("Number of lines in file: "+"%d\n", lines)
	fmt.Println("")
	str := string(b)
	for i := 32; i <= 126; i++ {
		iterate(i, str);
	}
	var symbol string
	var tall= 0
	var n int = 0
	var teller int = 0

	for j := 0; j < 5; j++ {
		tall = 0
		for l := 0; l < len(antall); l++ {
			n = antall[l]
			if n > tall {
				tall = n
				teller = l

			}
		}
		symbol = rune[teller]
		antall[teller] = 0
		fmt.Println(j+1, ".", "Rune:", symbol, ",", "Counts:", tall)
	}
}

func iterate(tall int, tekst string) {
	st := string(tall)
	symbol := fmt.Sprintf("%s", st)
	antallet := strings.Count(tekst, symbol)
	antall = append(antall, antallet)
	rune = append(rune, symbol)

}



func logError(e error) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", e.Error())

}

func getFilnavn() (string, error) {
	if len(os.Args) == 1 {
		return "", errors.New(
			fmt.Sprintf("No file arg provided\nUsage: %s "+
				"[file]\n", os.Args[0]))
	}

	return os.Args[1], nil
}

func countLinesInFile(filnavn string) (int, error) {
	file, err := os.Open(filnavn)

	if err != nil {
		return 0, err
	}

	buf := make([]byte, 1024)
	lines := 0

	for {
		readBytes, err := file.Read(buf)

		if err != nil {
			if readBytes == 0 && err == io.EOF {
				err = nil
			}
			return lines, err
		}

		lines += bytes.Count(buf[:readBytes], []byte{'\n'})
	}

	return lines, nil
}
