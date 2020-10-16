package main


import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"
)

/**
Golang implementation of the Inline Images Protocol
https://www.iterm2.com/documentation-images.html
Usage: imgcat <path/to/image> <path/to/other/image> ...
 */

func main() {
	if len(os.Args) < 2 {
		//TODO: print usage
		fmt.Println("Please provide at least one argument")
		os.Exit(1)
	}

	pipeReader, pipeWriter := io.Pipe()

	go func() {
		defer pipeWriter.Close()
		err := writeImagesFromPaths(os.Args[1:], pipeWriter)
		if err != nil {
			panic(err)
		}
	}()

	_, err := io.Copy(os.Stdout, pipeReader)
	if err != nil {
		panic(err)
	}
}

func writeImagesFromPaths(args []string, writer io.Writer) error {
	for _, filepath := range args {
		err := writeInlineImage(filepath, writer)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeInlineImage(filepath string, writer io.Writer) error {
	file, err := getFile(filepath)
	if err != nil {
		return err
	}
	err = imgcat(file, writer)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

func getFile(filepath string) (file *os.File, err error) {
	file, _err := os.Open(filepath)
	if _err != nil {
		return nil, _err
	}
	return file,nil
}

func imgcat(file *os.File, writer io.Writer) error {
	_, err := io.Copy(writer, strings.NewReader("\033]1337;File=inline=1:"))
	if err != nil {
		return err
	}
	err = writeAsBase64(file, writer)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, strings.NewReader("\a\n"))
	if err != nil {
		return err
	}

	return nil
}

func writeAsBase64(reader io.Reader, writer io.Writer) error {
	writeCloser := getBase64WriteCloser(writer)
	_, err := io.Copy(writeCloser, reader)
	if err != nil {
		return err
	}
	err = writeCloser.Close()
	if err != nil {
		return err
	}

	return nil
}

func getBase64WriteCloser(writer io.Writer) (writeCloser io.WriteCloser) {
	return base64.NewEncoder(base64.StdEncoding, writer)
}

