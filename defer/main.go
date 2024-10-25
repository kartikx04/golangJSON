package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if err := write("readme.txt", "Hi there, welcome to golang project!"); err != nil {
		log.Fatal("failed to create file")
	}
	if err := copyFile("readme.txt", "anotherFile.txt"); err != nil {
		log.Fatal("failed to copy file")
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

func copyFile(source, destination string) error {
	src, err := os.Open(source)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer dst.Close()

	n, err := io.Copy(dst, src)
	if err != nil {
		return err
	}

	fmt.Printf("copied %d bytes from %s to %s.\n", n, source, destination)

	if err := src.Close(); err != nil {
		return err
	}

	return dst.Close()
}
