package array

import "testing"

func TestSet(t *testing.T) {
	t.Run("Filter", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		expected := []int{4, 5, 6, 7, 8, 9}
		result := Filter(array, func(num int) bool {
			return num > 3
		})
		if len(result) != 6 {
			t.Errorf("Expected length of 6, got %d", len(result))
		}
		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("Expected %d, got %d", expected[i], result[i])
			}
		}
	})

	t.Run("Find", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		result, found := Find(array, func(num int) bool {
			return num == 5
		})
		if !found {
			t.Errorf("Expected to find 5 in array, but it was not found")
		}
		if result != 5 {
			t.Errorf("Expected to find 5, but found %d", result)
		}
	})

	t.Run("Find", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		result, found := Find(array, func(num int) bool {
			return num == 15
		})
		if found {
			t.Errorf("Expected to not find 15 in array, but it was found")
		}
		if result != 0 {
			t.Errorf("Expected to find 0, but found %d", result)
		}
	})

	t.Run("Some", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		result := Some(array, func(num int) bool {
			return num > 5
		})
		if !result {
			t.Errorf("Expected to find at least one element greater than 5, but it was not found")
		}
	})

	t.Run("Some", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		result := Some(array, func(num int) bool {
			return num > 10
		})
		if result {
			t.Errorf("Expected to not find any element greater than 10, but it was found")
		}
	})

	t.Run("Every", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		result := Every(array, func(num int) bool {
			return num > 0 && num < 10
		})
		if !result {
			t.Errorf("Expected to find all elements between 1 and 9, but it was not found")
		}
	})

	t.Run("Every", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		result := Every(array, func(num int) bool {
			return num > 0 && num < 5
		})
		if result {
			t.Errorf("Expected to not find any element between 1 and 4, but it was found")
		}
	})

	t.Run("Map", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		result := Map(array, func(num int) int {
			return num * 2
		})
		expected := []int{2, 4, 6, 8, 10, 12, 14, 16, 18}
		if len(result) != len(expected) {
			t.Errorf("Expected length of %d, got %d", len(expected), len(result))
		}
		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("Expected %d, got %d", expected[i], result[i])
			}
		}
	})

	t.Run("Reduce", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		result := Reduce(array, func(acc int, num int) int {
			return acc + num
		}, 0)
		if result != 45 {
			t.Errorf("Expected sum of array to be 45, but got %d", result)
		}
	})

	t.Run("Includes", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		result := Includes(array, 5)
		if !result {
			t.Errorf("Expected to find 5 in array, but it was not found")
		}
	})

	t.Run("Includes", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		result := Includes(array, 15)
		if result {
			t.Errorf("Expected to not find 15 in array, but it was found")
		}
	})

	t.Run("IndexOf", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		result := IndexOf(array, 5)
		if result != 4 {
			t.Errorf("Expected to find 5 at index 4, but found at index %d", result)
		}
	})

	t.Run("IndexOf", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		result := IndexOf(array, 15)
		if result != -1 {
			t.Errorf("Expected to not find 15 in array, but it was found at index %d", result)
		}
	})

	t.Run("LastIndexOf", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		result := LastIndexOf(array, 5)
		if result != 4 {
			t.Errorf("Expected to find 5 at last index 4, but found at index %d", result)
		}
	})

	t.Run("LastIndexOf", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		result := LastIndexOf(array, 15)
		if result != -1 {
			t.Errorf("Expected to not find 15 in array, but it was found at index %d", result)
		}
	})

	t.Run("FindIndex", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		result := FindIndex(array, func(num int) bool {
			return num == 5
		})
		if result != 4 {
			t.Errorf("Expected to find 5 at index 4, but found at index %d", result)
		}
	})

	t.Run("FindIndex", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		result := FindIndex(array, func(num int) bool {
			return num == 15
		})
		if result != -1 {
			t.Errorf("Expected to not find 15 in array, but it was found at index %d", result)
		}
	})

	t.Run("FindLastIndex", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		result := FindLastIndex(array, func(num int) bool {
			return num == 5
		})
		if result != 4 {
			t.Errorf("Expected to find 5 at last index 4, but found at index %d", result)
		}
	})

	t.Run("FindLastIndex", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		result := FindLastIndex(array, func(num int) bool {
			return num == 15
		})
		if result != -1 {
			t.Errorf("Expected to not find 15 in array, but it was found at index %d", result)
		}
	})

	t.Run("Sum", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		result := Sum(array)
		if result != 45 {
			t.Errorf("Expected sum of array to be 45, but got %d", result)
		}
	})

	t.Run("Reverse", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		result := Reverse(array)
		expected := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
		if len(result) != len(expected) {
			t.Errorf("Expected length of %d, got %d", len(expected), len(result))
		}
		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("Expected %d, got %d", expected[i], result[i])
			}
		}
	})
}
