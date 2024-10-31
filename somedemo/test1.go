package somedemo

import (
	"fmt"
	"time"
)

type Stu struct {
	Name string
	Age  int
	c    *Clas
}

type Clas struct {
	Name string
}

func testReadWriteMap() {
	m := make(map[string]int)

	go func ()  {
		for i:=0; i<100; i++ {
			k := fmt.Sprintf("k_%d", i)
			m[k] = i
		}
	}()

	go func ()  {
		for i:=0; i<100; i++ {
			k := fmt.Sprintf("k_%d", i)
			fmt.Println(m[k])
		}
	}()

	time.Sleep(2*time.Second)
}

func slicetest2() {
	// sliceGrow
	// s1 := []int{1,2,4}
	// for  i:=0; i<16; i++ {
	// 	s1 = append(s1, i)
	// 	fmt.Printf("addr: %p, len: %d, cap: %d\n", s1, len(s1), cap(s1))
	// }

	// set the len = cap = 10
	// s1 := make([]int, 10)
	// fmt.Printf("addr: %p, elements: %v, len: %d, cap: %d\n", s1, s1, len(s1), cap(s1))
	// s1 = append(s1, 10)
	// // after add first element, the slice grow, cap double
	// fmt.Printf("addr: %p, elements: %v, len: %d, cap: %d\n", s1, s1, len(s1), cap(s1))

	// cap=10 len=0   so no element in slice
	// s1 := make([]int, 0, 10)
	// fmt.Printf("addr: %p, elements: %v, len: %d, cap: %d\n", s1, s1, len(s1), cap(s1))
	// s1 = append(s1, 10)
	// // after add first element, not grow
	// fmt.Printf("addr: %p, elements: %v, len: %d, cap: %d\n", s1, s1, len(s1), cap(s1))

	s1 := make([]int, 10, 11)
	fmt.Printf("addr: %p, elements: %v, len: %d, cap: %d\n", s1, s1, len(s1), cap(s1))
	// s1 = append(s1, 10)
	// after add first element, not grow
	// fmt.Printf("addr: %p, elements: %v, len: %d, cap: %d\n", s1, s1, len(s1), cap(s1))

	s2 := s1[8:] // here s2 use same memory with s1, now s2 len=2, cap=3
	s2 = append(s2, []int{2,3,4,5,2}...) 
	// here cap 3 not enough, so s2 grow, this would allocate a new memory to s2
	fmt.Printf("addr: %p, elements: %v, len: %d, cap: %d\n", s2, s2, len(s2), cap(s2))

}

func Slicetest() {

	// s1 := []int{1,2,3}
	// s2 := s1[:]
  // // this is shallow copy, they address are same
	// fmt.Printf("s1 address: %p\n", s1)
	// fmt.Printf("s2 address: %p\n", s2)
	
	// s3 := make([]int, 0, 3)
	// copy(s3, s2)
	// // this is deep copy, they address are different
	// fmt.Printf("s3 address: %p\n", s3)
	// var s []int
	s := make([]int, 3) // init len=cap
	s = append(s, 1) // trigger slice grow
	fmt.Printf("before modifySlice s addr:%p, len:%d, cap:%d\n",  &s, len(s), cap(s))
	modifySlice(s)
	// s len=3 cap=3
	// the before slice grow not change current s slice header
	fmt.Printf("after modifySlice s addr:%p, len:%d, cap:%d\n",  &s, len(s), cap(s))
	// fmt.Println(s[3]) // panic, becasuse s len=3 s[3]out of range

	fmt.Println(s)
	s = append(s, 2)
	fmt.Println(s)
	fmt.Printf("add 2 to s addr:%p, len:%d, cap:%d\n",  s, len(s), cap(s))

	// new slice header, but use common base array
	// len=cap
	// s1 := s[1:]
	// fmt.Printf("s1 addr:%p, len:%d, cap:%d\n",  s1, len(s1), cap(s1))

	s4 := []int{1,2,3}
	fmt.Printf("s4 addr:%p, len:%d, cap:%d\n",  s4, len(s4), cap(s4))
	s5 := s4[1:] // s5 has a offset address from s4
	// s5 := s4[3:] // s5 address same as s4 but len=0 cap=0
	fmt.Println(s5)
	fmt.Printf("s5 addr:%p, len:%d, cap:%d\n",  s5, len(s5), cap(s5))
	s5[0] = 22 // s5 and s4 share a part of same memory, so this would change s4[1]
	fmt.Printf("s4 addr:%p, eles:%v len:%d, cap:%d\n",  s4, s4, len(s4), cap(s4))

	s6 := append([]int{1,2,3}, s4[2:]...) // s6 address != s4 address
	// s6 := append(s4[:1], s4[2:]...) // s6 address = s4 address
	fmt.Printf("s6 addr:%p, len:%d, cap:%d\n",  s6, len(s6), cap(s6))
	s7 := s4[1:]
	fmt.Println(s5)
	fmt.Printf("s5 addr:%p, len:%d, cap:%d\n",  s7, len(s7), cap(s7))
}
func modifySlice(s []int) {
	s = append(s, 2)
	// slice grow, addr/len/cap change
	fmt.Printf("in modifySlice s addr:%p, len:%d, cap:%d\n",  &s, len(s), cap(s))
}

func fortest() {
	ch := make(chan int, 1)
	go func ()  {
		priority:
			for {
				select {
				case v := <-ch:
					fmt.Println(v)
				default:
					break priority
				}
			}
			
			fmt.Println("after priority work")
	}()

	ch <- 1
	time.Sleep(2*time.Second)
}


type Slice[T string|int] []T
func generic() {

	printArray([]string{"a", "b"})
	printArray([]int{1, 2})
	// printArray([]bool{}) // error

	Add(2,3)
	Add[string]("2", "3")
	
	// Add(2, "3") // error
}

func printArray[T string|int|Stu](arr []T) {
	for _, v := range arr {
		println(v)
	}
}

func Add[T int|float64|string](a T, b T) T {return a+b}