package array

// Return a new slice containing all selected elements
// 返回一个新切片, 包含所有通过选择的元素
func Filter[T any](arr []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range arr {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Return the first element that satisfies the predicate, and whether it was found
// 返回第一个满足条件的元素和是否找到的布尔值
func Find[T any](arr []T, predicate func(T) bool) (T, bool) {
	for _, item := range arr {
		if predicate(item) {
			return item, true
		}
	}
	var zero T
	return zero, false
}

// Check if at least one element in the array satisfies the predicate
// 检查数组中是否有至少一个元素满足条件
func Some[T any](arr []T, predicate func(T) bool) bool {
	for _, item := range arr {
		if predicate(item) {
			return true
		}
	}
	return false
}

// Check if all elements in the array satisfy the predicate
// 检查数组中的所有元素是否都满足条件
func Every[T any](arr []T, predicate func(T) bool) bool {
	for _, item := range arr {
		if !predicate(item) {
			return false
		}
	}
	return true
}

// Return a new array with the results of calling a provided function on every element in the calling array
// 对数组中的每个元素执行函数并返回新数组
func Map[T any, U any](arr []T, mapper func(T) U) []U {
	result := make([]U, len(arr))
	for i, item := range arr {
		result[i] = mapper(item)
	}
	return result
}

// Reduce applies a reducer function against an accumulator and each element in the array (from left to right)
// 对数组中的每个元素执行累加器函数
func Reduce[T any, U any](arr []T, reducer func(U, T) U, initial U) U {
	result := initial
	for _, item := range arr {
		result = reducer(result, item)
	}
	return result
}

// Check if the array contains a certain element
// 检查数组是否包含某个元素
func Includes[T comparable](arr []T, item T) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}

// Return the index of the first occurrence of the specified element in the array, or -1 if it is not present
// 返回元素在数组中的第一个索引，不存在则返回 -1
func IndexOf[T comparable](arr []T, item T) int {
	for i, v := range arr {
		if v == item {
			return i
		}
	}
	return -1
}

// Return the index of the last occurrence of the specified element in the array, or -1 if it is not present
// 返回元素在数组中的最后一个索引，不存在则返回-1
func LastIndexOf[T comparable](arr []T, item T) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == item {
			return i
		}
	}
	return -1
}

// Return the index of the first element that satisfies the predicate, or -1 if no element is found
// 返回第一个满足条件的元素的索引，找不到返回-1
func FindIndex[T any](arr []T, predicate func(T) bool) int {
	for i, item := range arr {
		if predicate(item) {
			return i
		}
	}
	return -1
}

// Return the index of the last element that satisfies the predicate, or -1 if no element is found
// 返回最后一个满足条件的元素的索引，找不到返回-1
func FindLastIndex[T any](arr []T, predicate func(T) bool) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if predicate(arr[i]) {
			return i
		}
	}
	return -1
}

// Return the sum of all elements in the array
// 数组元素求和
func Sum[T int | int32 | int64 | float32 | float64 | uint | uint32 | uint64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// Reverse the order of the elements in the array
// 数组元素反转
func Reverse[T any](slice []T) []T {
	result := make([]T, len(slice))
	for i, v := range slice {
		result[len(slice)-1-i] = v
	}
	return result
}
