package algo

import "fmt"

func quickSort(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		quickSort(arr, low, pivotIndex-1)

		fmt.Println("------LF------")
		quickSort(arr, pivotIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	// i .... j .. pivot
	// 1 8 5 6 2 3 7 9 4
	// next one of i is always bigger than pivot
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	fmt.Println(arr)
	return i + 1
}

func RunQuickSort() {
	// Example usage of quickSort
	arr := []int{1, 8, 5, 6, 2, 3, 7, 9, 4}
	fmt.Println("Unsorted array:", arr)

	quickSort(arr, 0, len(arr)-1)

	fmt.Println("Sorted array:", arr)
}
