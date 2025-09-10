/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package vslices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAppendIfNotExist(t *testing.T) {
	// Test with int slice
	t.Run("int slice", func(t *testing.T) {
		// Test appending to empty slice
		result := AppendIfNotExist([]int{}, 1)
		expected := []int{1}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}

		// Test appending new element
		result = AppendIfNotExist([]int{1, 2, 3}, 4)
		expected = []int{1, 2, 3, 4}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}

		// Test appending existing element
		result = AppendIfNotExist([]int{1, 2, 3}, 2)
		expected = []int{1, 2, 3}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}

		// Test appending existing element at beginning
		result = AppendIfNotExist([]int{1, 2, 3}, 1)
		expected = []int{1, 2, 3}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}

		// Test appending existing element at end
		result = AppendIfNotExist([]int{1, 2, 3}, 3)
		expected = []int{1, 2, 3}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test with string slice
	t.Run("string slice", func(t *testing.T) {
		// Test appending to empty slice
		result := AppendIfNotExist([]string{}, "hello")
		expected := []string{"hello"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}

		// Test appending new element
		result = AppendIfNotExist([]string{"a", "b", "c"}, "d")
		expected = []string{"a", "b", "c", "d"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}

		// Test appending existing element
		result = AppendIfNotExist([]string{"a", "b", "c"}, "b")
		expected = []string{"a", "b", "c"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}

		// Test with empty string
		result = AppendIfNotExist([]string{"a", "", "c"}, "")
		expected = []string{"a", "", "c"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}

		// Test appending empty string to slice without empty string
		result = AppendIfNotExist([]string{"a", "b", "c"}, "")
		expected = []string{"a", "b", "c", ""}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test with float64 slice
	t.Run("float64 slice", func(t *testing.T) {
		// Test appending new element
		result := AppendIfNotExist([]float64{1.1, 2.2, 3.3}, 4.4)
		expected := []float64{1.1, 2.2, 3.3, 4.4}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}

		// Test appending existing element
		result = AppendIfNotExist([]float64{1.1, 2.2, 3.3}, 2.2)
		expected = []float64{1.1, 2.2, 3.3}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}

		// Test with zero value
		result = AppendIfNotExist([]float64{1.1, 0.0, 3.3}, 0.0)
		expected = []float64{1.1, 0.0, 3.3}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test with bool slice
	t.Run("bool slice", func(t *testing.T) {
		// Test appending true to slice with false
		result := AppendIfNotExist([]bool{false}, true)
		expected := []bool{false, true}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}

		// Test appending false to slice with false
		result = AppendIfNotExist([]bool{false}, false)
		expected = []bool{false}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}

		// Test appending true to slice with true
		result = AppendIfNotExist([]bool{true}, true)
		expected = []bool{true}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test with single element slice
	t.Run("single element slice", func(t *testing.T) {
		// Test appending same element
		result := AppendIfNotExist([]int{42}, 42)
		expected := []int{42}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}

		// Test appending different element
		result = AppendIfNotExist([]int{42}, 43)
		expected = []int{42, 43}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test with duplicate elements in slice
	t.Run("slice with duplicates", func(t *testing.T) {
		// Test appending element that appears multiple times
		result := AppendIfNotExist([]int{1, 2, 2, 3, 2}, 2)
		expected := []int{1, 2, 2, 3, 2}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}

		// Test appending new element to slice with duplicates
		result = AppendIfNotExist([]int{1, 2, 2, 3, 2}, 4)
		expected = []int{1, 2, 2, 3, 2, 4}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

// Benchmark tests
func BenchmarkAppendIfNotExist(b *testing.B) {
	slice := []int{1, 2, 3, 4, 5}
	b.ResetTimer()

	b.Run("existing element", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = AppendIfNotExist(slice, 3)
		}
	})

	b.Run("new element", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = AppendIfNotExist(slice, 6)
		}
	})

	b.Run("empty slice", func(b *testing.B) {
		emptySlice := []int{}
		for i := 0; i < b.N; i++ {
			_ = AppendIfNotExist(emptySlice, 1)
		}
	})
}

func TestAppendIfCheckPass(t *testing.T) {
	// Test with int slice
	t.Run("int slice", func(t *testing.T) {
		// Test checker returns true - should append
		checker := func(slice []int) bool { return len(slice) < 5 }
		result := AppendIfCheckPass([]int{1, 2, 3}, 4, checker)
		expected := []int{1, 2, 3, 4}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}

		// Test checker returns false - should not append
		checker = func(slice []int) bool { return len(slice) > 5 }
		result = AppendIfCheckPass([]int{1, 2, 3}, 4, checker)
		expected = []int{1, 2, 3}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}

		// Test with empty slice - checker returns true
		checker = func(slice []int) bool { return len(slice) == 0 }
		result = AppendIfCheckPass([]int{}, 1, checker)
		expected = []int{1}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}

		// Test with empty slice - checker returns false
		checker = func(slice []int) bool { return len(slice) > 0 }
		result = AppendIfCheckPass([]int{}, 1, checker)
		expected = []int{}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test with string slice
	t.Run("string slice", func(t *testing.T) {
		// Test checker based on slice content
		checker := func(slice []string) bool {
			for _, s := range slice {
				if s == "forbidden" {
					return false
				}
			}
			return true
		}

		// Should append - no forbidden string
		result := AppendIfCheckPass([]string{"a", "b"}, "c", checker)
		expected := []string{"a", "b", "c"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}

		// Should not append - contains forbidden string
		result = AppendIfCheckPass([]string{"a", "forbidden", "b"}, "c", checker)
		expected = []string{"a", "forbidden", "b"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}

		// Test with empty string
		result = AppendIfCheckPass([]string{"a", "b"}, "", checker)
		expected = []string{"a", "b", ""}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test with float64 slice
	t.Run("float64 slice", func(t *testing.T) {
		// Test checker based on sum
		checker := func(slice []float64) bool {
			sum := 0.0
			for _, f := range slice {
				sum += f
			}
			return sum < 10.0
		}

		// Should append - sum is less than 10
		result := AppendIfCheckPass([]float64{1.1, 2.2, 3.3}, 1.0, checker)
		expected := []float64{1.1, 2.2, 3.3, 1.0}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}

		// Should not append - sum would exceed 10
		result = AppendIfCheckPass([]float64{5.0, 4.0, 2.0}, 1.0, checker)
		expected = []float64{5.0, 4.0, 2.0}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}

		// Test with zero value
		result = AppendIfCheckPass([]float64{1.0, 2.0}, 0.0, checker)
		expected = []float64{1.0, 2.0, 0.0}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test with bool slice
	t.Run("bool slice", func(t *testing.T) {
		// Test checker that allows append only if slice has even length
		checker := func(slice []bool) bool { return len(slice)%2 == 0 }

		// Should append - even length
		result := AppendIfCheckPass([]bool{true, false}, true, checker)
		expected := []bool{true, false, true}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}

		// Should not append - odd length
		result = AppendIfCheckPass([]bool{true}, false, checker)
		expected = []bool{true}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}

		// Test with empty slice - even length (0)
		result = AppendIfCheckPass([]bool{}, true, checker)
		expected = []bool{true}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test with custom struct
	t.Run("custom struct", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		// Test checker based on struct field
		checker := func(slice []Person) bool {
			for _, p := range slice {
				if p.Age > 50 {
					return false
				}
			}
			return true
		}

		// Should append - no one over 50
		people := []Person{{"Alice", 25}, {"Bob", 30}}
		newPerson := Person{"Charlie", 35}
		result := AppendIfCheckPass(people, newPerson, checker)
		expected := []Person{{"Alice", 25}, {"Bob", 30}, {"Charlie", 35}}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}

		// Should not append - someone over 50
		people = []Person{{"Alice", 25}, {"Bob", 60}}
		result = AppendIfCheckPass(people, newPerson, checker)
		expected = []Person{{"Alice", 25}, {"Bob", 60}}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test edge cases
	t.Run("edge cases", func(t *testing.T) {
		// Test with nil checker (this would panic, but we test with always true)
		alwaysTrue := func(slice []int) bool { return true }
		result := AppendIfCheckPass([]int{1, 2}, 3, alwaysTrue)
		expected := []int{1, 2, 3}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}

		// Test with always false checker
		alwaysFalse := func(slice []int) bool { return false }
		result = AppendIfCheckPass([]int{1, 2}, 3, alwaysFalse)
		expected = []int{1, 2}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}

		// Test with single element slice
		checker := func(slice []int) bool { return slice[0] > 0 }
		result = AppendIfCheckPass([]int{5}, 10, checker)
		expected = []int{5, 10}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}

		result = AppendIfCheckPass([]int{-5}, 10, checker)
		expected = []int{-5}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

// Benchmark tests for AppendIfCheckPass
func BenchmarkAppendIfCheckPass(b *testing.B) {
	slice := []int{1, 2, 3, 4, 5}
	checker := func(s []int) bool { return len(s) < 10 }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = AppendIfCheckPass(slice, i, checker)
	}
}

func ExampleAppendIfCheckPass() {
	// Example with length-based checker
	slice := []int{1, 2, 3}
	lengthChecker := func(s []int) bool { return len(s) < 5 }

	result := AppendIfCheckPass(slice, 4, lengthChecker)
	fmt.Println(result) // [1 2 3 4]

	result = AppendIfCheckPass([]int{1, 2, 3, 4, 5}, 6, lengthChecker)
	fmt.Println(result) // [1 2 3 4 5]

	// Example with content-based checker
	strings := []string{"apple", "banana"}
	contentChecker := func(s []string) bool {
		for _, str := range s {
			if str == "forbidden" {
				return false
			}
		}
		return true
	}

	result2 := AppendIfCheckPass(strings, "cherry", contentChecker)
	fmt.Println(result2) // [apple banana cherry]

	// Output:
	// [1 2 3 4]
	// [1 2 3 4 5]
	// [apple banana cherry]
}

// Example test
func ExampleAppendIfNotExist() {
	// Adding new element
	slice := []int{1, 2, 3}
	result := AppendIfNotExist(slice, 4)
	// result: [1, 2, 3, 4]
	_ = result

	// Adding existing element
	result = AppendIfNotExist(slice, 2)
	// result: [1, 2, 3] (unchanged)
	_ = result
}
