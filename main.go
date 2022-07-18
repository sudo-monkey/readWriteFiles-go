package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	writeFile()
	defer readFile("./fromString.txt")

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func writeFile() {
	content := "Hello from Go!"
	file, err := os.Create("./fromString.txt")
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Wrote a file with %v characters\n", length)
	defer file.Close()
}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Println("Text read from file:", string(data))
}
