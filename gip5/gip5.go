package main

// /*
// 	Listing 5.1 Simple log usage
// */

// import (
// 	"log"
// )

// func main() {
// 	log.Println("This ia a regular Message")
// 	log.Fatalln("This ia a fatal Message")
// 	log.Println("This is at end of function. After fatal Message")
// }

// /*
// 	Listing 5.2 Logging to a file
// */

// import (
// 	"log"
// 	"os"
// )

// func main() {
// 	logfile, _ := os.Create("./log.txt")
// 	defer logfile.Close()

// 	logger := log.New(logfile, "example ", log.LstdFlags|log.Lshortfile)
// 	logger.Println("This is a regular log")
// 	logger.Fatal("This is a fatal log")
// 	logger.Println("This a log after fatal. Should not come")
// }

// /*
// 	Listing 5.3 Network log client
// */

// import (
// 	"log"
// 	"net"
// )

// func main() {
// 	conn, err := net.Dial("tcp", "localhost:1902")
// 	if err != nil {
// 		panic("unbale to connect to localhost:1902")
// 	}

// 	defer conn.Close()

// 	f := log.LstdFlags | log.Lshortfile
// 	logger := log.New(conn, "example ", f)
// 	logger.Println("This is a regular message")
// 	logger.Panicln("This is a panic message")

// }

// /*
// Listing 5.6 Logging to the system log

// For checking logs:
// Mac: tail -F /var/log/system.log
// Unix: /var/log/messages
// */

// import (
// 	"log/syslog"
// )

// func main() {
// 	logger, err := syslog.New(syslog.LOG_LOCAL3, "narwhal")
// 	if err != nil {
// 		panic("Cannot attach to syslog")
// 	}

// 	logger.Debug("Debug Message")
// 	logger.Notice("Notice Message")
// 	logger.Warning("Warning Message")
// 	logger.Alert("Alert Message")
// }

/*
	Listing 5.8 Using the Stack function
*/

import (
	"fmt"
	"runtime"
)

func main() {
	foo()
}

func foo() {
	bar()
}

func bar() {
	buf := make([]byte, 1024)
	runtime.Stack(buf, false)
	fmt.Printf("Stack Trace is:\n%s\n", buf)
}
