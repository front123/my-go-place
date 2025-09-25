package algo

import "fmt"

/*
 * Heap sort - https://en.wikipedia.org/wiki/Heapsort
 */

func sift(arr []int, i int, arrLen int) []int {
	done := false

	tmp := 0
	maxChild := 0

	for (i*2+1 < arrLen) && (!done) {
		if i*2+1 == arrLen-1 {
			maxChild = i*2 + 1
		} else if arr[i*2+1] > arr[i*2+2] {
			maxChild = i*2 + 1
		} else {
			maxChild = i*2 + 2
		}

		if arr[i] < arr[maxChild] {
			tmp = arr[i]
			arr[i] = arr[maxChild]
			arr[maxChild] = tmp
			i = maxChild
		} else {
			done = true
		}
	}

	return arr
}

/*
      2
    1   7
  3 4  6 5
*/

// []int{2, 1, 7, 3, 4, 6, 5}

func HeapSort(arr []int) {
	i := 0
	tmp := 0

	for i = len(arr)/2 - 1; i >= 0; i-- {
		arr = sift(arr, i, len(arr))
	}

	fmt.Println(arr)
	for i = len(arr) - 1; i >= 1; i-- {
		tmp = arr[0]    //7
		arr[0] = arr[i] //2
		arr[i] = tmp    //[2 4 6 3 1 7]
		fmt.Println("-- ", arr)
		arr = sift(arr, 0, i)
		fmt.Println(arr)
	}
}

func sift_min(arr []int, i int, arrlen int) {
	for 2*i+1 < arrlen {
		changed := false
		min_idx := i
		p := arr[i]
		l := arr[2*i+1]
		if l < p {
			min_idx = 2*i + 1
			changed = true
		}
		if 2*i+2 < arrlen && arr[2*i+2] < arr[min_idx] {
			min_idx = 2*i + 2
			changed = true
		}
		if !changed {
			return
		}

		arr[i], arr[min_idx] = arr[min_idx], arr[i]
		i++
	}
}

// []int{2, 1, 7, 3, 4, 6, 5}
func HeapSortMin(arr []int) {
	i := 0
	tmp := 0

	for i = len(arr)/2 - 1; i >= 0; i-- {
		sift_min(arr, i, len(arr))
	}

	fmt.Println(arr)
	for i = len(arr) - 1; i >= 1; i-- {
		tmp = arr[0]    //7
		arr[0] = arr[i] //2
		arr[i] = tmp    //[2 4 6 3 1 7]
		fmt.Println("-- ", arr)
		sift_min(arr, 0, i)
		fmt.Println(arr)
	}

	// reverse desc -> asc
	for i, j := 0, len(arr)-1; i < j; {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	fmt.Println(arr)
}
