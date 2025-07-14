package set

import "testing"

func TestSet(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		s := New[int]()
		s.Add(1)
		if !s.Contains(1) {
			t.Errorf("Expected 1 to be in set, but it was not")
		}
	})

	t.Run("Remove", func(t *testing.T) {
		s := New[int]()
		s.Add(1)
		s.Remove(1)
		if s.Contains(1) {
			t.Errorf("Expected 1 to be removed from set, but it was not")
		}
	})

	t.Run("Contains", func(t *testing.T) {
		s := New[int]()
		s.Add(1)
		if !s.Contains(1) {
			t.Errorf("Expected 1 to be in set, but it was not")
		}
	})

	t.Run("Len", func(t *testing.T) {
		s := New[int]()
		s.Add(1)
		s.Add(2)
		s.Add(3)
		if s.Len() != 3 {
			t.Errorf("Expected length of 3, got %d", s.Len())
		}
	})

	t.Run("Items", func(t *testing.T) {
		s := New[int]()
		s.Add(1)
		s.Add(2)
		s.Add(3)
		items := s.Items()
		if len(items) != 3 {
			t.Errorf("Expected 3 items, got %d", len(items))
		}
		if (items[0] != 1 && items[0] != 2 && items[0] != 3) ||
			(items[1] != 1 && items[1] != 2 && items[1] != 3) ||
			(items[2] != 1 && items[2] != 2 && items[2] != 3) {
			t.Errorf("Expected items to be [1, 2, 3], got %v", items)
		}
	})

	t.Run("Union", func(t *testing.T) {
		s1 := New[int]()
		s1.Add(1)
		s1.Add(2)

		s2 := New[int]()
		s2.Add(2)
		s2.Add(3)

		s3 := s1.Union(s2)
		if s3.Len() != 3 {
			t.Errorf("Expected length of 3, got %d", s3.Len())
		}
		if !s3.Contains(1) || !s3.Contains(2) || !s3.Contains(3) {
			t.Errorf("Expected 1, 2, and 3 to be in set, but they were not")
		}
	})

	t.Run("Intersection", func(t *testing.T) {
		s1 := New[int]()
		s1.Add(1)
		s1.Add(2)

		s2 := New[int]()
		s2.Add(2)
		s2.Add(3)

		s3 := s1.Intersection(s2)
		if s3.Len() != 1 {
			t.Errorf("Expected length of 1, got %d", s3.Len())
		}
		if !s3.Contains(2) {
			t.Errorf("Expected 2 to be in set, but it was not")
		}
	})

	t.Run("Difference", func(t *testing.T) {
		s1 := New[int]()
		s1.Add(1)
		s1.Add(2)

		s2 := New[int]()
		s2.Add(2)
		s2.Add(3)

		s3 := s1.Difference(s2)
		if s3.Len() != 1 {
			t.Errorf("Expected length of 1, got %d", s3.Len())
		}
		if !s3.Contains(1) {
			t.Errorf("Expected 1 to be in set, but it was not")
		}
	})

	t.Run("IsSubset", func(t *testing.T) {
		s1 := New[int]()
		s1.Add(1)
		s1.Add(2)

		s2 := New[int]()
		s2.Add(2)
		s2.Add(3)

		if s1.IsSubset(s2) {
			t.Errorf("Expected s1 to not be a subset of s2")
		}

		s2.Add(1)
		if !s1.IsSubset(s2) {
			t.Errorf("Expected s1 to be a subset of s2")
		}
	})

	t.Run("IsSuperset", func(t *testing.T) {
		s1 := New[int]()
		s1.Add(1)
		s1.Add(2)

		s2 := New[int]()
		s2.Add(2)
		s2.Add(3)

		if s1.IsSuperset(s2) {
			t.Errorf("Expected s1 to not be a superset of s2")
		}

		s1.Add(3)
		if !s1.IsSuperset(s2) {
			t.Errorf("Expected s1 to be a superset of s2")
		}
	})

	t.Run("Equal", func(t *testing.T) {
		s1 := New[int]()
		s1.Add(1)
		s1.Add(2)

		s2 := New[int]()
		s2.Add(1)
		s2.Add(2)

		if !s1.Equal(s2) {
			t.Errorf("Expected s1 and s2 to be equal sets")
		}

		s2.Add(3)
		if s1.Equal(s2) {
			t.Errorf("Expected s1 and s2 to be different sets")
		}
	})

	t.Run("Clear", func(t *testing.T) {
		s1 := New[int]()
		s1.Add(1)
		s1.Add(2)
		s1.Clear()
		if s1.Len() != 0 {
			t.Errorf("Expected length of 0, got %d", s1.Len())
		}
	})

	t.Run("Clone", func(t *testing.T) {
		s1 := New[int]()
		s1.Add(1)
		s1.Add(2)

		s2 := s1.Clone()
		if !s1.Equal(s2) {
			t.Errorf("Expected s1 and s2 to be equal sets")
		}
		s2.Add(3)
		if s1.Equal(s2) {
			t.Errorf("Expected s1 and s2 to be different sets")
		}
	})

	t.Run("Custom Struct", func(t *testing.T) {
		type testStruct struct {
			Data int
		}

		t1 := testStruct{Data: 1}

		s1 := New[testStruct]()
		s1.Add(t1)
		s1.Add(testStruct{Data: 2})

		if !s1.Contains(t1) {
			t.Errorf("Expected testStruct{Data: 1} to be in set, but it was not")
		}
		if !s1.Contains(testStruct{Data: 2}) {
			t.Errorf("Expected testStruct{Data: 2} to be in set, but it was not")
		}
	})
}
