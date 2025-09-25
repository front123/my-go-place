// main.go
package main

import (
	"fmt"
	"go-example/algo"
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

type company struct {
	Users []int64
	Name  string
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

	// models.InitDB()

	// str := `{"user": null, "name": "goodcompany"}`
	// var com company
	// err := json.Unmarshal([]byte(str), &com)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Printf("company: %+v", com)

	// t, err := time.Parse("2006-01-02 15:04:05", "2024-10-30 15:58:27")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// // log.Println(t)
	// dur := time.Since(t)
	// log.Println(dur.Seconds())

	// fmt.Printf("%s \t %s \t %d \t %s\n", "po-134", "good", 12346, "hellofdsf")

	// loc, _ := time.LoadLocation("Asia/Shanghai")
	// t, err := time.ParseInLocation("2006-01-02 15:04:05", "2025-08-04 15:38:51", loc)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println(FormatHKDateTime(t, "2006-01-02,15:04"))

	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// for i := 0; i < 100; i++ {
	// 	fmt.Println(r.Intn(10))
	// }

	// arr := []int{2, 1, 7, 3, 4, 6, 5}
	// algo.HeapSort(arr)
	// algo.HeapSortMin(arr)

	// arr = algo.QuickSort2(arr)
	// fmt.Println(arr)

	// algo.RunQuickSort()

	// var studs StudentArray
	// studs = append(studs, &Student{Name: "front"})
	// changestuds(&studs)
	// bs, _ := json.Marshal(studs)
	// fmt.Println(string(bs))

	// var studs1, studs2 []*Student
	// var studss []*Student
	// studss = append(studss, studs1..., studs2...)
	// bs, _ := json.Marshal(stud)
	// fmt.Println(string(bs))

	// fmt.Println(^uint64(0))

	graph := algo.NewGraph(8)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 3)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 5)
	graph.AddEdge(4, 6)
	graph.AddEdge(5, 7)
	graph.AddEdge(6, 7)

	graph.BFS(0, 7)
}

type StudentArray []*Student

type Student struct {
	Name string
}

func changestuds(studs *StudentArray) {
	*studs = append(*studs, &Student{Name: "johb"})
	(*studs)[0].Name = "Front1"
}

func FormatHKDateTime(dateTime time.Time, format string) (dateString string) {
	if dateTime.IsZero() {
		return ""
	}
	location, _ := time.LoadLocation("Asia/Hong_Kong")
	dateString = dateTime.In(location).Format(format)
	if dateString == "01/01/0001" {
		dateString = ""
	}
	return
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
