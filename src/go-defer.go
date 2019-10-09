package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func readFile(path string) ([]byte, error) {
	parentPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	fullPath := filepath.Join(parentPath, path)
	file, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return read(file)
}

func read(file *os.File) ([]byte, error) {
	br := bufio.NewReader(file)
	var buf bytes.Buffer
	var rtn_err error
	for {
		ba, isPrefix, err := br.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error: %s\n", err)
			rtn_err = errors.New("read error!!! ")
			break
		}
		buf.Write(ba)
		if !isPrefix {
			buf.WriteByte('\n')
		}
	}

	return buf.Bytes(), rtn_err
}

func deferIt2() {
	for i := 1; i < 5; i++ {
		defer fmt.Print(i)
	}
}

func main9() {
	deferIt2()

	path := "./src/go-defer.go"
	ba, err := readFile(path)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("The content of '%s':\n%s\n", path, ba)
}
