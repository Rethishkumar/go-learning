package main

// /*
// LIsting 3.1 Using a goroutine to run a task
// */

// import (
// 	"fmt"
// 	"io"
// 	"os"
// 	"time"
// )

// func main() {
// 	go echo(os.Stdin, os.Stdout)
// 	time.Sleep(30 * time.Second)
// 	fmt.Println("Timed out")
// 	os.Exit(0)
// }

// func echo(in io.Reader, out io.Writer) {
// 	io.Copy(out, in)
// }

// /*
// 	Listing 3.2 An anonymous goroutine
// */

// import (
// 	"fmt"
// 	"runtime"
// )

// func main() {
// 	fmt.Println("Outside a go routine")
// 	go func() {
// 		fmt.Println("Inside a go routine")
// 	}()

// 	fmt.Println("Outside again")

// 	runtime.Gosched()
// }

// /*
// Listing 3.4 Compressing files in parallel with a wait group
// */
// import (
// 	"compress/gzip"
// 	"fmt"
// 	"io"
// 	"os"
// 	"sync"
// )

// func main() {
// 	var wg sync.WaitGroup

// 	var i int = -1
// 	var file string
// 	for i, file = range os.Args[1:] {
// 		wg.Add(1)
// 		go func(filename string) {
// 			compress(filename)
// 			wg.Done()
// 		}(file)
// 	}

// 	wg.Wait()

// 	fmt.Printf("Compressed %d files\n", i+1)
// }

// func compress(filename string) error {
// 	in, err := os.Open(filename)
// 	if err != nil {
// 		return err
// 	}

// 	defer in.Close()

// 	out, err := os.Create(filename + ".gz")
// 	if err != nil {
// 		return err
// 	}

// 	defer out.Close()

// 	gzout := gzip.NewWriter(out)
// 	_, err = io.Copy(gzout, in)
// 	gzout.Close()
// 	return err
// }

// /*
// 	Listing 3.6 Word counter with locks
// */

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// 	"sync"
// )

// func main() {
// 	var wg sync.WaitGroup

// 	w := newWords()
// 	for _, f := range os.Args[1:] {
// 		wg.Add(1)
// 		go func(filename string) {
// 			if err := tallyWords(filename, w); err != nil {
// 				fmt.Println(err.Error())
// 			}
// 			wg.Done()
// 		}(f)
// 	}

// 	wg.Wait()

// 	fmt.Println("Words that appear more than once")
// 	for word, count := range w.found {
// 		if count > 1 {
// 			fmt.Printf("%s occured %d times \n", word, count)
// 		}
// 	}
// }

// type words struct {
// 	sync.Mutex
// 	found map[string]int
// }

// func newWords() *words {
// 	return &words{found: map[string]int{}}
// }

// func (w *words) add(word string, n int) {
// 	w.Lock()
// 	defer w.Unlock()
// 	count, ok := w.found[word]
// 	if !ok {
// 		w.found[word] = n
// 		return
// 	}
// 	w.found[word] = count + n
// }

// func tallyWords(filename string, dict *words) error {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return err
// 	}

// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	scanner.Split(bufio.ScanWords)
// 	for scanner.Scan() {
// 		word := strings.ToLower(scanner.Text())
// 		dict.add(word, 1)
// 	}
// 	return scanner.Err()
// }

// /*
// 	Listing 3.7 Using multiple channels
// */

// import (
// 	"fmt"
// 	"os"
// 	"time"
// )

// func main() {
// 	done := time.After(30 * time.Second)
// 	echo := make(chan []byte)
// 	go readStdin(echo)

// 	for {
// 		select {
// 		case buf := <-echo:
// 			os.Stdout.Write(buf)
// 		case <-done:
// 			fmt.Println("Timed out")
// 			os.Exit(0)
// 		}
// 	}
// }

// func readStdin(out chan<- []byte) {
// 	for {
// 		data := make([]byte, 1024)
// 		l, _ := os.Stdin.Read(data)
// 		if l > 0 {
// 			out <- data
// 		}
// 	}
// }

/*
Listing 3.11 Using a close channel
*/

import (
	"fmt"
	"time"
)

func main() {
	untill := time.After(5 * time.Second)
	done := make(chan bool)
	msg := make(chan string)

	go send(msg, done)

	for {
		select {
		case m := <-msg:
			fmt.Println(m)
		case <-untill:
			fmt.Println("Timeout 1")
			done <- true
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Timeout 2")
			return
		}
	}
}

func send(msg chan<- string, done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("Done")
			close(msg)
			return
		default:
			msg <- "hello"
			time.Sleep(500 * time.Millisecond)
		}
	}
}
