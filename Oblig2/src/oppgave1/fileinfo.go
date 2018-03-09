package main

import (
	"fmt"
	"log"
	"os"
	"io/ioutil"
)


func main() {
	filename := os.Args[1]

	fi, err := os.Lstat(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("information about file <" + filename + ">:")

	data, err := ioutil.ReadFile("fileinfo.go")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nSize: %d bytes", len(data))


		i := len(data)
		f := float64(i)
		fmt.Printf(" %f Kilobytes", f/1e3)
		fmt.Printf(" %f Megabytes", f/1e6)
		fmt.Printf(" %e Gigabytes", f/1e9)

	switch mode := fi.Mode(); {
	case mode.IsDir():
		fmt.Println("\nis a directory")
	case !mode.IsDir():
		fmt.Println("\nis not a directory")
	}

	switch mode1 := fi.Mode(); {
	case mode1.IsRegular():
		fmt.Println("is a regular file")
	case !mode1.IsRegular():
		fmt.Println("is not a regular file")
	}
	
	switch mode2 := fi.Mode(); {
	case mode2&os.ModePerm == 0777:
		fmt.Println("Has Unix permission bits")
	case mode2&os.ModePerm != 0777:
		fmt.Println("Has not Unix permission bits")
	}

	switch mode3 := fi.Mode(); {
	case mode3&os.ModeAppend != 0:
		fmt.Println("is append only")
	case mode3&os.ModeAppend == 0:
		fmt.Println("is not append only")
	}

	switch mode4 := fi.Mode(); {
	case mode4&os.ModeDevice != 0:
		fmt.Println("is device file")
	case mode4&os.ModeDevice == 0:
		fmt.Println("is not device file")
	}

	switch mode5 := fi.Mode(); {
	case mode5&os.ModeCharDevice != 0:
		fmt.Println("is Unix character device")
	case mode5&os.ModeCharDevice == 0:
		fmt.Println("is not Unix character device")
	}
	
	switch mode6 := fi.Mode(); {
	case mode6&os.ModeSymlink != 0:
		fmt.Println("symbolic link")
	case mode6&os.ModeSymlink == 0:
		fmt.Println("is not symbolic link")
	}
}
