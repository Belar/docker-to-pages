package main

import (
	"fmt"
	"os"
)

func main() {
	saveFile, error := os.Create("./dist/test.txt")

	if error != nil {
		fmt.Println("File creation failed")
		return
	}

	testString := []byte("qwerty")
	saveFile.Write(testString)

	fmt.Println("Saved to file")
}
