// main.go
package main

import (
	"fmt"
	"go-example/models"
	"net/http"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	_ "net/http/pprof"
)

func add(a, b int) *int {
	res := 0

	res = a + b

	return &res
}

func main() {

	// var ch = make(chan int)
	// close(ch) // panic on nil channel
	// ch <- 1 // deadlock
	// <-ch // deadlock

	// testRespClose()

	// usePprof()

	// useSyncPool()

	// useAtomicCount()

	models.InitDB()

}

func startPprof() {
	// http://localhost:6060/debug/pprof
	go func() {
		http.ListenAndServe(":6060", nil)
	}()
}

func useSyncPool() {
	type t struct {
		Name string
	}

	t0 := &t{Name: "front"}
	fmt.Printf("t0 addr: %p\n", t0)

	p := sync.Pool{
		New: func() any {
			return t0
		},
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		v := p.Get().(*t)
		fmt.Printf("g1 get t from pool addr: %p\n", v)
		p.Put(v)
	}()

	go func() {
		defer wg.Done()
		v := p.Get().(*t)
		fmt.Printf("g2 get t from pool addr: %p\n", v)
		p.Put(v)
	}()

	wg.Wait()
}

func useAtomicCount() {
	var count = int32(0)
	var wg sync.WaitGroup
	var gNum = 1000
	wg.Add(gNum)

	for i := 0; i < gNum; i++ {
		go func() {
			defer wg.Done()
			// atomic.AddInt32(&count, 1) // use this is OK

			// use below has problem
			c := count + 1 // issue this
			atomic.StoreInt32(&count, c)
		}()
	}

	wg.Wait()
	fmt.Println(count)
}

func usePprof() {
	for i := 0; i < 100; i++ {
		go func() {
			select {}
		}()
	}

	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	select {}
}

func testRespClose() {
	var wg = sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			requestWithNoClose()
			// requestWithClose()
		}()
	}

	wg.Wait()
	time.Sleep(5 * time.Second)
	fmt.Println(runtime.NumGoroutine())
}

func requestWithNoClose() {
	_, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
}

func requestWithClose() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
}

func copyReforCopyVal() {
	arr := []int{1, 2, 3}
	fmt.Println(len(arr))
	var p = arr
	fmt.Printf("in main, addr p: %p\n", p)
	fmt.Printf("in main, addr arr: %p\n", arr)
	fmt.Printf("in main, addr &p: %p\n", &p)
	fmt.Printf("in main, addr &arr: %p\n", &arr)

	arr2 := [3]int{1, 2, 3}
	var p2 = arr2
	fmt.Printf("in main, addr p2: %p\n", &p2)
	fmt.Printf("in main, addr arr2: %p\n", &arr2)
}

func m1(i *int64) {
	fmt.Printf("in m1, addr i: %p\n", i)
	fmt.Printf("in m1, addr &i: %p\n", &i)
}
