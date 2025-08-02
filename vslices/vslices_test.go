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