package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"
)

func createFile(name, b64 string) error {

	fmt.Println(b64)

	reader := strings.NewReader(b64)
	decoder := base64.NewDecoder(base64.StdEncoding, reader)

	output, err := os.Create(name)

	if err != nil {
		return err
	}
	defer output.Close()

	_, err = io.Copy(output, decoder)

	return err
}

func deleteFile(name string) error {
	err := os.Remove(name)
	return err
}
