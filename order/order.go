package order

import (
	"cmp"
)

// BubbleSort 冒泡排序
func BubbleSort[T cmp.Ordered](arr []T) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

// SelectionSort 选择排序
func SelectionSort[T cmp.Ordered](arr []T) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// InsertionSort 插入排序
func InsertionSort[T cmp.Ordered](arr []T) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// QuickSort 快速排序
func QuickSort[T cmp.Ordered](arr []T) {
	if len(arr) <= 1 {
		return
	}
	quickSort(arr, 0, len(arr)-1)
}

func quickSort[T cmp.Ordered](arr []T, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partition[T cmp.Ordered](arr []T, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// MergeSort 归并排序
func MergeSort[T cmp.Ordered](arr []T) {
	if len(arr) <= 1 {
		return
	}
	mid := len(arr) / 2
	left := make([]T, mid)
	right := make([]T, len(arr)-mid)
	copy(left, arr[:mid])
	copy(right, arr[mid:])

	MergeSort(left)
	MergeSort(right)

	merge(arr, left, right)
}

func merge[T cmp.Ordered](arr, left, right []T) {
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}
}

// HeapSort 堆排序
func HeapSort[T cmp.Ordered](arr []T) {
	n := len(arr)

	// 构建最大堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// 逐个提取元素
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func heapify[T cmp.Ordered](arr []T, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

// ShellSort 希尔排序
func ShellSort[T cmp.Ordered](arr []T) {
	n := len(arr)
	gap := n / 2

	for gap > 0 {
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i
			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = temp
		}
		gap /= 2
	}
}

// CountingSort 计数排序 (仅适用于整数)
func CountingSort(arr []int) {
	if len(arr) == 0 {
		return
	}

	max := arr[0]
	min := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	rangeSize := max - min + 1
	count := make([]int, rangeSize)

	for _, num := range arr {
		count[num-min]++
	}

	idx := 0
	for i := 0; i < rangeSize; i++ {
		for count[i] > 0 {
			arr[idx] = i + min
			idx++
			count[i]--
		}
	}
}

// RadixSort 基数排序 (仅适用于非负整数)
func RadixSort(arr []int) {
	if len(arr) == 0 {
		return
	}

	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	for exp := 1; max/exp > 0; exp *= 10 {
		countSort(arr, exp)
	}
}

func countSort(arr []int, exp int) {
	n := len(arr)
	output := make([]int, n)
	count := make([]int, 10)

	for i := 0; i < n; i++ {
		index := (arr[i] / exp) % 10
		count[index]++
	}

	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		index := (arr[i] / exp) % 10
		output[count[index]-1] = arr[i]
		count[index]--
	}

	for i := 0; i < n; i++ {
		arr[i] = output[i]
	}
}

// BucketSort 桶排序
func BucketSort(arr []float64) {
	if len(arr) == 0 {
		return
	}

	// 确定桶的数量
	n := len(arr)
	buckets := make([][]float64, n)

	// 将元素分配到桶中
	for _, num := range arr {
		index := int(float64(n) * num)
		if index == n {
			index = n - 1
		}
		buckets[index] = append(buckets[index], num)
	}

	// 对每个桶进行排序
	for i := 0; i < n; i++ {
		InsertionSort(buckets[i])
	}

	// 合并桶
	index := 0
	for i := 0; i < n; i++ {
		for _, num := range buckets[i] {
			arr[index] = num
			index++
		}
	}
}
