package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)

	f, _ := os.Create("./log.txt")
	f.Close()

	logCh := make(chan string, 50)

	go func() {
		for {
			msg, ok := <-logCh
			if ok {
				f, _ := os.OpenFile("./log.txt", os.O_APPEND, os.ModeAppend)

				logTime := time.Now().Format(time.RFC3339)
				f.WriteString(logTime + " - " + msg)
				f.Close()
			} else {
				break
			}
		}
	}()

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			go func(i, j int) {
				msg := fmt.Sprintf("%d + %d = %d\n", i, j, i+j)
				logCh <- msg
				fmt.Print(msg)
			}(i, j)
		}
	}

	fmt.Scanln()
}

/*
Step-1 Initial demo code with sync.Mutex
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(4)

	mutex := new(sync.Mutex)


	for i:=1; i < 10;i++ {
		for j:=1; j<10;j++ {
			mutex.Lock()
			go func() {
				fmt.Printf("%d + %d = %d\n", i, j, i+j)
				mutex.Unlock()
			}()
		}
	}

	fmt.Scanln()
}
*/

/*
Step-2 Update to use channel to force sequential access
package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)

	mutex := make(chan bool, 1)


	for i:=1; i < 10;i++ {
		for j:=1; j<10;j++ {
			mutex <- true
			go func() {
				fmt.Printf("%d + %d = %d\n", i, j, i+j)
				<-mutex
			}()
		}
	}

	fmt.Scanln()
}
*/

/*
Step-3 Async logger with channels
package main

import (
	"fmt"
	"runtime"
	"os"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	f, _ := os.Create("./log.txt")
	f.Close()

	logCh := make(chan string, 50)

	go func() {
		for {
			msg, ok := <- logCh
			if ok {
				f, _ := os.OpenFile("./log.txt", os.O_APPEND, os.ModeAppend)

				logTime := time.Now().Format(time.RFC3339)
				f.WriteString(logTime + " - " + msg)

				f.Close()
			} else {
				break
			}
		}
	}()

	mutex := make(chan bool, 1)


	for i:=1; i < 10;i++ {
		for j:=1; j<10;j++ {
			mutex <- true
			go func() {
				msg := fmt.Sprintf("%d + %d = %d\n", i, j, i+j)
				logCh <- msg
				fmt.Print(msg)
				<-mutex
			}()
		}
	}

	fmt.Scanln()
}
*/

/*
Step-4 Remove mutex and pass parameters into goroutine
package main

import (
	"fmt"
	"runtime"
	"os"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)

	f, _ := os.Create("./log.txt")
	f.Close()

	logCh := make(chan string, 50)

	go func() {
		for {
			msg, ok := <- logCh
			if ok {
				f, _ := os.OpenFile("./log.txt", os.O_APPEND, os.ModeAppend)

				logTime := time.Now().Format(time.RFC3339)
				f.WriteString(logTime + " - " + msg)
				f.Close()
			} else {
				break
			}
		}
	}()

	for i:=1; i < 10;i++ {
		for j:=1; j<10;j++ {
			go func(i, j int) {
				msg := fmt.Sprintf("%d + %d = %d\n", i, j, i+j)
				logCh <- msg
				fmt.Print(msg)
			}(i, j)
		}
	}

	fmt.Scanln()
}
*/