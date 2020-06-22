package main

// /*
// Listing 4.13 Recovering from a panic
// */

// import (
// 	"errors"
// 	"fmt"
// )

// func main() {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Printf("Trapped panic: %s (%T)\n", err, err)
// 		}
// 	}()
// 	yikes()
// }
// func yikes() {
// 	panic(errors.New("something bad happened."))
// }

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var file io.ReadCloser
	file, err := openCsv("data.csv")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	defer file.Close()
}

func openCsv(filename string) (file *os.File, err error) {
	defer func() {
		if r := recover; r != nil {
			file.Close()
			err := r.(error)
		}
	}()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open file \n")
		return file, err
	}

	RemoveEmptyLines(file)
	return file, err
}
