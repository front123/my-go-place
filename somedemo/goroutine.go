package somedemo

import (
	"fmt"
	"time"
)

// panic can only be recovered in itself goroutine
func go1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("main recover:", r)
		}
	}()

	defer t1()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("gorountine recover:", r)
			}
		}()

		panic("goroutine panic...")

	}()

	time.Sleep(2 * time.Second)

	fmt.Println("main done")
}

func t1() {
	fmt.Println("t1 defer run")
}

func chann1() {
	ch := make(chan int, 1)

	ch<-1
	fmt.Println("ok send 1")

	// buf full and no goroutine reading data from chan, so deadlock
	// this line will block goroutine
	ch<-2 
	fmt.Println("ok send 2")
	go loop(ch)

	time.Sleep(2*time.Second)
}

func loop(ch chan int) {
	for  {
		select {
		case v := <-ch:
			fmt.Println("get data:", v)
		}
	}
}