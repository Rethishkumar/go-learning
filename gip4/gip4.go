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

// /*
// 	Listing 4.16 Cleanup
// */

// import (
// 	"errors"
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main() {
// 	var file io.ReadCloser
// 	file, err := openCsv("data.csv")
// 	if err != nil {
// 		fmt.Printf("Error: %s", err)
// 		return
// 	}

// 	defer file.Close()
// }

// func openCsv(filename string) (file *os.File, err error) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			file.Close()
// 			err = r.(error)
// 		}
// 	}()

// 	file, err := os.Open(filename)
// 	if err != nil {
// 		fmt.Printf("Failed to open file \n")
// 		return file, err
// 	}

// 	RemoveEmptyLines(file)
// 	return file, err
// }

// func RemoveEmptyLines(f *os.File) {
// 	panic(errors.New("Failed parse"))
// }

// /*
// 	Listing 4.19 Handle panics on a goroutine
// */

// import (
// 	"bufio"
// 	"errors"
// 	"fmt"
// 	"net"
// )

// func main() {
// 	listen()
// }

// func listen() {
// 	listener, err := net.Listen("tcp", ":1026")
// 	if err != nil {
// 		fmt.Println("Failed to open port 1026")
// 		return
// 	}

// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			fmt.Println("Error accepting connection")
// 			continue
// 		}

// 		go handle(conn)
// 	}
// }

// func handle(conn net.Conn) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println("closing connection and recovering from error")
// 			conn.Close()
// 		}
// 	}()

// 	reader := bufio.NewReader(conn)
// 	data, err := reader.ReadBytes('\n')
// 	if err != nil {
// 		fmt.Println("Failed to read from socket.")
// 		conn.Close()
// 	}

// 	response(data, conn)
// }

// func response(data []byte, conn net.Conn) {
// 	defer func() {
// 		fmt.Println("closing client connection")
// 		conn.Close()
// 	}()

// 	conn.Write(data)
// 	panic(errors.New("panic in response"))
// }

/*
	Listing 4.23 Using safely.Go to trap panics
*/

import (
	"errors"
	"fmt"
	"time"

	"github.com/Masterminds/cookoo/safely"
)

func message() {
	fmt.Println("Inside Message")
	panic(errors.New("Oops!"))
}

func main() {
	safely.Go(message)
	fmt.Println("In main")
	time.Sleep(1000)
}
