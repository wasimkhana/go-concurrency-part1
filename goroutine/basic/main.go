package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(2)
	godur, _ := time.ParseDuration("10ms")

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Hello")
			time.Sleep(godur)

		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Go")
			time.Sleep(godur)
		}
	}()

	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)

}

/*
	step-1 Two funcs called in sequence
	func(){
		println("Hello")
	}()

	func() {
		println("Go")
	}()
*/

/*
	step-2 Main thread exits before go routines have chance to run
	go func() {
		println("Hello")
	}()

	go func() {
		println("Go")
	}()
*/

/*
	step-3 Add sleep to main thread to allow go routines to execute
	go func() {
		println("Hello")
	}()

	go func() {
		println("Go")
	}()

	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
*/

/*
	step-4 Goroutines will swap execution contexts
	godur, _ := time.ParseDuration("10ms")

	go func() {
		for i:=0; i< 100; i++ {
			println("Hello")
			time.Sleep(godur)

		}
	}()

	go func() {
		for i:=0; i< 100; i++ {
			println("Go")
			time.Sleep(godur)
		}
	}()


	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
*/

/*
	step-5 With multiple processor threads, application becomes non-determinant
	runtime.GOMAXPROCS(2)
	godur, _ := time.ParseDuration("10ms")

	go func() {
		for i:=0; i< 100; i++ {
			println("Hello")
			time.Sleep(godur)

		}
	}()

	go func() {
		for i:=0; i< 100; i++ {
			println("Go")
			time.Sleep(godur)
		}
	}()


	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
*/
