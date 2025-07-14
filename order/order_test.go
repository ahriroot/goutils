package order

import "testing"

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestSet(t *testing.T) {
	t.Run("BubbleSort", func(t *testing.T) {
		array := []int{4, 3, 1, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		BubbleSort(array)
		if !equal(array, expected) {
			t.Errorf("Expected %v, got %v", expected, array)
		}
	})

	t.Run("SelectionSort", func(t *testing.T) {
		array := []int{4, 3, 1, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		SelectionSort(array)
		if !equal(array, expected) {
			t.Errorf("Expected %v, got %v", expected, array)
		}
	})

	t.Run("InsertionSort", func(t *testing.T) {
		array := []int{4, 3, 1, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		InsertionSort(array)
		if !equal(array, expected) {
			t.Errorf("Expected %v, got %v", expected, array)
		}
	})

	t.Run("QuickSort", func(t *testing.T) {
		QuickSort([]int{})
		array := []int{4, 3, 1, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		QuickSort(array)
		if !equal(array, expected) {
			t.Errorf("Expected %v, got %v", expected, array)
		}
	})

	t.Run("MergeSort", func(t *testing.T) {
		array := []int{4, 3, 1, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		MergeSort(array)
		if !equal(array, expected) {
			t.Errorf("Expected %v, got %v", expected, array)
		}
	})

	t.Run("HeapSort", func(t *testing.T) {
		array := []int{4, 3, 1, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		HeapSort(array)
		if !equal(array, expected) {
			t.Errorf("Expected %v, got %v", expected, array)
		}
	})

	t.Run("ShellSort", func(t *testing.T) {
		array := []int{4, 3, 1, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		ShellSort(array)
		if !equal(array, expected) {
			t.Errorf("Expected %v, got %v", expected, array)
		}
	})

	t.Run("CountingSort", func(t *testing.T) {
		CountingSort([]int{})
		array := []int{4, 3, 1, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		CountingSort(array)
		if !equal(array, expected) {
			t.Errorf("Expected %v, got %v", expected, array)
		}
	})

	t.Run("RadixSort", func(t *testing.T) {
		RadixSort([]int{})
		array := []int{4, 3, 1, 1, 5, 9, 2, 6, 5, 3, 5}
		expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}
		RadixSort(array)
		if !equal(array, expected) {
			t.Errorf("Expected %v, got %v", expected, array)
		}
	})

	t.Run("BucketSort", func(t *testing.T) {
		BucketSort([]float64{})
		array := []float64{0.78, 0.17, 0.39, 0.26, 0.72, 0.94, 0.21, 1.08, 0.12, 0.23}
		expected := []float64{0.12, 0.17, 0.21, 0.23, 0.26, 0.39, 0.72, 0.78, 0.94, 1.08}
		BucketSort(array)
		for i := range array {
			if array[i] != expected[i] {
				t.Errorf("Expected %v, got %v", expected, array)
				break
			}
		}
	})
}
