package main

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed data.zip
var zipFile []byte

func main() {
	f, err := os.Create("sync.zip")
	if err != nil {
		fmt.Printf("failed: %s", err)
	}
	defer f.Close()

	f.Write(zipFile)
}
