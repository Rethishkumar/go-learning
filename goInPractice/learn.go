package main

// import (
// 	"fmt"
// )

// func Names() (string, string) {
// 	first = "Foo"
// 	second = "Bar"
// 	return
// }

// func main() {
// 	n1, n2 := Names()
// 	fmt.Println(n1, n2)

// 	n3, _ := Names()
// 	fmt.Println(n3)
// }

// import (
// 	"fmt"
// )

// func Names() (first string, second string) {
// 	first = "Foo"
// 	second = "Bar"
// 	return
// }

// func main() {
// 	n1, n2 := Names()
// 	fmt.Println(n1, n2)

// 	n3, _ := Names()
// 	fmt.Println(n3)
// }

/* Listing 1.3 Read TCP status: read_status.go */

// import (
// 	"bufio"
// 	"fmt"
// 	"net"
// )

// func main() {
// 	conn, _ := net.Dial("tcp", "golang.org:80")
// 	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
// 	status, _ := bufio.NewReader(conn).ReadString('\n')
// 	fmt.Println(status)
// }

/* Listing 1.4 HTTP GET: http_get.go  */

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// func main() {
// 	resp, _ := http.Get("http://example.com")
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	fmt.Println(string(body))
// 	resp.Body.Close()
// }

/* Listing 1.6 Printing concurrently */

// import (
// 	"fmt"
// 	"time"
// )

// func count() {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(i)
// 		time.Sleep(time.Millisecond * 1)
// 	}
// }

// func main() {
// 	go count()
// 	time.Sleep(time.Millisecond * 2)
// 	fmt.Println("Hello World")
// 	time.Sleep(time.Millisecond * 5)
// }

/* Listing 1.7 Using channels: channel.go */

// import (
// 	"fmt"
// 	"time"
// )

// func printCount(c chan int) {
// 	num := 0
// 	for num >= 0 {
// 		num = <-c
// 		fmt.Println(num, " ")
// 	}
// }

// func main() {
// 	c := make(chan int)
// 	// below gives error because the goroutine stops at -1
// 	// fatal error: all goroutines are asleep - deadlock!
// 	// 	goroutine 1 [chan send]:
// 	// 	main.main()
// 	/Users/rnair/work/go/goInPractice/learn.go:111 +0x11d
// 	a := []int{8, 6, 7, 5, 3, 0, 9, -1, -2}
// 	// Actual example as per book a := []int{8, 6, 7, 5, 3, 0, 9, -1}
// 	go printCount(c)

// 	for _, v := range a {
// 		c <- v
// 	}
// 	time.Sleep(time.Millisecond * 1)
// 	fmt.Println("End of Main")
// }

/* Listing 1.10 Hello World: hello.go */
// import (
// 	"fmt"
// )

// func getName() string {
// 	return "World"
// }

// func main() {
// 	name := getName()
// 	fmt.Println("Hello ", name)
// }

/* Listing 1.16 Hello World web server: inigo.go */

import (
	"fmt"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello my name is Rethish. Help Me God")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:4000", nil)
}
