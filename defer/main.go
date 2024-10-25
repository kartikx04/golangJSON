package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if err := write("readme.txt", "Hi there, welcome to golang project!"); err != nil {
		log.Fatal(err)
	}
}

func write(filename, text string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()
	_, err = io.WriteString(file, text)
	if err != nil {
		return err
	}

	return file.Close()
}
