package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/gabriel-vasile/mimetype"
)

func main() {
	fileBytes, err := os.ReadFile(os.Args[len(os.Args)-1])
	mimeType, mimeErr := mimetype.DetectFile(os.Args[len(os.Args)-1])

	if err != nil {
		panic(err)
	}

	if mimeErr != nil {
		panic(mimeErr)
	}

	base64 := base64.StdEncoding.EncodeToString(fileBytes)

	fmt.Print("data:" + mimeType.String() + ";base64," + base64)
}
