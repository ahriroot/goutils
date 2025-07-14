// # Package set provides a generic collection implementation based on a map. Set 包提供了基于 map 实现的泛型集合
//
// [ New ]: creates and returns a new empty Set. 创建并返回一个新的空集合
//
// [ Add ]: inserts an element into the set. 向集合中添加元素
//
// [ Remove ]: deletes an element from the set. 从集合中删除元素
//
// [ Contains ]: checks if an element exists in the set. 检查集合中是否包含某元素
//
// [ Len ]: returns the number of elements in the set. 返回集合中元素的数量
//
// [ Items ]: returns a slice containing all elements in the set. 返回包含所有元素的切片
package set

// Set is a generic collection implementation based on a map. 基于map实现的泛型集合
type Set[T comparable] struct {
	items map[T]struct{}
}

// New creates and returns a new empty Set. 创建并返回一个新的空集合
//
// Example:
//
//	s := New[int]()
//	s := New[string]()
func New[T comparable]() *Set[T] {
	return &Set[T]{
		items: make(map[T]struct{}),
	}
}

// Add inserts an element into the set. 向集合中添加元素
//
// If the element already exists, it has no effect. 如果元素已经存在则无效果
//
// Example:
//
//	s.Add(1)
//	s.Add("a")
func (s *Set[T]) Add(item T) {
	s.items[item] = struct{}{}
}

// Remove deletes an element from the set. 从集合中删除元素
//
// If the element does not exist, it has no effect. 如果元素不存在则无效果
//
// Example:
//
//	s.Remove(1)
//	s.Remove("a")
func (s *Set[T]) Remove(item T) {
	delete(s.items, item)
}

// Contains checks if an element exists in the set. 检查集合中是否包含某元素
//
// Example:
//
//	if s.Contains(1) {
//		// do something
//	}
//	if s.Contains("a") {
//		// do something
//	}
func (s *Set[T]) Contains(item T) bool {
	_, ok := s.items[item]
	return ok
}

// Len returns the number of elements in the set. 返回集合中元素的数量
//
// Example:
//
//	length := s.Len()
func (s *Set[T]) Len() int {
	return len(s.items)
}

// Items returns a slice containing all elements in the set. 返回包含所有元素的切片
//
// Example:
//
//	items := s.Items()
//	for _, item := range items {
//		// do something with item
//	}
func (s *Set[T]) Items() []T {
	items := make([]T, 0, len(s.items))
	for item := range s.items {
		items = append(items, item)
	}
	return items
}

// Union returns a new set containing all elements from both sets. 合并两个集合，返回一个新的集合
//
// Example:
//
//	s1 := New[int]()
//	s1.Add(1)
//	s1.Add(2)
//
//	s2 := New[int]()
//	s2.Add(2)
//	s2.Add(3)
//
//	s3 := s1.Union(s2)
//	fmt.Println(s3.Items()) // Output: [1 2 3]
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := New[T]()
	for item := range s.items {
		result.Add(item)
	}
	for item := range other.items {
		result.Add(item)
	}
	return result
}

// Intersection returns a new set containing elements that are present in both sets. 交集，返回一个新的集合
//
// Example:
//
//	s1 := New[int]()
//	s1.Add(1)
//	s1.Add(2)
//
//	s2 := New[int]()
//	s2.Add(2)
//	s2.Add(3)
//
//	s3 := s1.Intersection(s2)
//	fmt.Println(s3.Items()) // Output: [2]
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := New[T]()
	for item := range s.items {
		if other.Contains(item) {
			result.Add(item)
		}
	}
	return result
}

// Difference returns a new set containing elements that are present in s but not in other. 差集，返回一个新的集合
//
// Example:
//
//	s1 := New[int]()
//	s1.Add(1)
//	s1.Add(2)
//
//	s2 := New[int]()
//	s2.Add(2)
//	s2.Add(3)
//
//	s3 := s1.Difference(s2)
//	fmt.Println(s3.Items()) // Output: [1]
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := New[T]()
	for item := range s.items {
		if !other.Contains(item) {
			result.Add(item)
		}
	}
	return result
}

// IsSubset checks if s is subset of other. 判断当前集合是否是另一个集合的子集
//
// Example:
//
//	s1 := New[int]()
//	s1.Add(1)
//	s1.Add(2)
//
//	s2 := New[int]()
//	s2.Add(2)
//	s2.Add(3)
//
//	if s1.IsSubset(s2) {
//		fmt.Println("s1 is subset of s2")
//	} else {
//		fmt.Println("s1 is not subset of s2")
//	}
//
// Output:
//
//	s1 is subset of s2
func (s *Set[T]) IsSubset(other *Set[T]) bool {
	for item := range s.items {
		if !other.Contains(item) {
			return false
		}
	}
	return true
}

// IsSuperset checks if s is superset of other. 判断当前集合是否是另一个集合的超集
//
// Example:
//
//	s1 := New[int]()
//	s1.Add(1)
//	s1.Add(2)
//
//	s2 := New[int]()
//	s2.Add(2)
//	s2.Add(3)
//
//	if s1.IsSuperset(s2) {
//		fmt.Println("s1 is superset of s2")
//	} else {
//		fmt.Println("s1 is not superset of s2")
//	}
//
// Output:
//
//	s1 is not superset of s2
func (s *Set[T]) IsSuperset(other *Set[T]) bool {
	return other.IsSubset(s)
}

// Equal checks if s is equal to other. 判断两个集合是否相等
//
// Example:
//
//	s1 := New[int]()
//	s1.Add(1)
//	s1.Add(2)
//
//	s2 := New[int]()
//	s2.Add(2)
//	s2.Add(3)
//
//	if s1.Equal(s2) {
//		fmt.Println("s1 is equal to s2")
//	} else {
//		fmt.Println("s1 is not equal to s2")
//	}
//
// Output:
//
//	s1 is not equal to s2
func (s *Set[T]) Equal(other *Set[T]) bool {
	if s.Len() != other.Len() {
		return false
	}
	return s.IsSubset(other)
}

// Clear removes all elements from the set. 清空集合
func (s *Set[T]) Clear() {
	s.items = make(map[T]struct{})
}

// Clone returns a new set with the same elements as s (shallow copy). 克隆集合 (浅拷贝)
func (s *Set[T]) Clone() *Set[T] {
	result := New[T]()
	for item := range s.items {
		result.Add(item)
	}
	return result
}
