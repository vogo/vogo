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
	"strconv"
	"testing"
)

func TestAppendIfNotExist(t *testing.T) {
	tests := []struct {
		name     string
		slice    interface{}
		item     interface{}
		expected interface{}
	}{
		// Int slice tests
		{"int empty slice", []int{}, 1, []int{1}},
		{"int append new element", []int{1, 2, 3}, 4, []int{1, 2, 3, 4}},
		{"int append existing element", []int{1, 2, 3}, 2, []int{1, 2, 3}},
		{"int append existing element at beginning", []int{1, 2, 3}, 1, []int{1, 2, 3}},
		{"int append existing element at end", []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"int single element same", []int{42}, 42, []int{42}},
		{"int single element different", []int{42}, 43, []int{42, 43}},
		{"int slice with duplicates existing", []int{1, 2, 2, 3, 2}, 2, []int{1, 2, 2, 3, 2}},
		{"int slice with duplicates new", []int{1, 2, 2, 3, 2}, 4, []int{1, 2, 2, 3, 2, 4}},
		// String slice tests
		{"string empty slice", []string{}, "hello", []string{"hello"}},
		{"string append new element", []string{"a", "b", "c"}, "d", []string{"a", "b", "c", "d"}},
		{"string append existing element", []string{"a", "b", "c"}, "b", []string{"a", "b", "c"}},
		{"string append existing element at beginning", []string{"a", "b", "c"}, "a", []string{"a", "b", "c"}},
		{"string append existing element at end", []string{"a", "b", "c"}, "c", []string{"a", "b", "c"}},
		{"string with empty string existing", []string{"a", "", "c"}, "", []string{"a", "", "c"}},
		{"string append empty string", []string{"a", "b", "c"}, "", []string{"a", "b", "c", ""}},
		// Float64 slice tests
		{"float64 append new element", []float64{1.1, 2.2, 3.3}, 4.4, []float64{1.1, 2.2, 3.3, 4.4}},
		{"float64 append existing element", []float64{1.1, 2.2, 3.3}, 2.2, []float64{1.1, 2.2, 3.3}},
		{"float64 with zero value", []float64{1.1, 0.0, 3.3}, 0.0, []float64{1.1, 0.0, 3.3}},
		// Bool slice tests
		{"bool append true to false", []bool{false}, true, []bool{false, true}},
		{"bool append false to false", []bool{false}, false, []bool{false}},
		{"bool append true to true", []bool{true}, true, []bool{true}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result interface{}
			switch s := tt.slice.(type) {
			case []int:
				result = AppendIfNotExist(s, tt.item.(int))
			case []string:
				result = AppendIfNotExist(s, tt.item.(string))
			case []float64:
				result = AppendIfNotExist(s, tt.item.(float64))
			case []bool:
				result = AppendIfNotExist(s, tt.item.(bool))
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
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
	type Person struct {
		Name string
		Age  int
	}

	tests := []struct {
		name     string
		slice    interface{}
		item     interface{}
		checker  interface{}
		expected interface{}
	}{
		// Int slice tests
		{"int len < 5 should append", []int{1, 2, 3}, 4, func(slice []int, elem int) bool { return len(slice) < 5 }, []int{1, 2, 3, 4}},
		{"int len > 5 should not append", []int{1, 2, 3}, 4, func(slice []int, elem int) bool { return len(slice) > 5 }, []int{1, 2, 3}},
		{"int empty slice len == 0 should append", []int{}, 1, func(slice []int, elem int) bool { return len(slice) == 0 }, []int{1}},
		{"int empty slice len > 0 should not append", []int{}, 1, func(slice []int, elem int) bool { return len(slice) > 0 }, []int{}},
		{"int always true", []int{1, 2}, 3, func(slice []int, elem int) bool { return true }, []int{1, 2, 3}},
		{"int always false", []int{1, 2}, 3, func(slice []int, elem int) bool { return false }, []int{1, 2}},
		{"int single element positive", []int{5}, 10, func(slice []int, elem int) bool { return slice[0] > 0 }, []int{5, 10}},
		{"int single element negative", []int{-5}, 10, func(slice []int, elem int) bool { return slice[0] > 0 }, []int{-5}},
		{"int elem > 5 should append", []int{1, 2}, 10, func(slice []int, elem int) bool { return elem > 5 }, []int{1, 2, 10}},
		{"int elem <= 5 should not append", []int{1, 2}, 3, func(slice []int, elem int) bool { return elem > 5 }, []int{1, 2}},
		// String slice tests
		{"string no forbidden should append", []string{"a", "b"}, "c", func(slice []string, elem string) bool {
			return elem != "forbidden"
		}, []string{"a", "b", "c"}},
		{"string with forbidden should not append", []string{"a", "b", "c"}, "forbidden", func(slice []string, elem string) bool {
			return elem != "forbidden"
		}, []string{"a", "b", "c"}},
		{"string append empty string", []string{"a", "b"}, "", func(slice []string, elem string) bool {
			return true
		}, []string{"a", "b", ""}},
		// Float64 slice tests
		{"float64 sum < 10 should append", []float64{1.1, 2.2, 3.3}, 1.0, func(slice []float64, elem float64) bool {
			sum := 0.0
			for _, f := range slice {
				sum += f
			}
			return sum < 10.0
		}, []float64{1.1, 2.2, 3.3, 1.0}},
		{"float64 sum >= 10 should not append", []float64{5.0, 4.0, 2.0}, 1.0, func(slice []float64, elem float64) bool {
			sum := 0.0
			for _, f := range slice {
				sum += f
			}
			return sum < 10.0
		}, []float64{5.0, 4.0, 2.0}},
		{"float64 append zero value", []float64{1.0, 2.0}, 0.0, func(slice []float64, elem float64) bool {
			sum := 0.0
			for _, f := range slice {
				sum += f
			}
			return sum < 10.0
		}, []float64{1.0, 2.0, 0.0}},
		// Bool slice tests
		{"bool even length should append", []bool{true, false}, true, func(slice []bool, elem bool) bool { return len(slice)%2 == 0 }, []bool{true, false, true}},
		{"bool odd length should not append", []bool{true}, false, func(slice []bool, elem bool) bool { return len(slice)%2 == 0 }, []bool{true}},
		{"bool empty slice even length should append", []bool{}, true, func(slice []bool, elem bool) bool { return len(slice)%2 == 0 }, []bool{true}},
		// Custom struct tests
		{"person no one over 50 should append", []Person{{"Alice", 25}, {"Bob", 30}}, Person{"Charlie", 35}, func(slice []Person, elem Person) bool {
			return elem.Age <= 50
		}, []Person{{"Alice", 25}, {"Bob", 30}, {"Charlie", 35}}},
		{"person someone over 50 should not append", []Person{{"Alice", 25}, {"Bob", 60}}, Person{"Charlie", 35}, func(slice []Person, elem Person) bool {
			for _, p := range slice {
				if p.Age > 50 {
					return false
				}
			}
			return true
		}, []Person{{"Alice", 25}, {"Bob", 60}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result interface{}
			switch s := tt.slice.(type) {
			case []int:
				result = AppendIfCheckPass(s, tt.item.(int), tt.checker.(func([]int, int) bool))
			case []string:
				result = AppendIfCheckPass(s, tt.item.(string), tt.checker.(func([]string, string) bool))
			case []float64:
				result = AppendIfCheckPass(s, tt.item.(float64), tt.checker.(func([]float64, float64) bool))
			case []bool:
				result = AppendIfCheckPass(s, tt.item.(bool), tt.checker.(func([]bool, bool) bool))
			case []Person:
				result = AppendIfCheckPass(s, tt.item.(Person), tt.checker.(func([]Person, Person) bool))
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

// Benchmark tests for AppendIfCheckPass
func BenchmarkAppendIfCheckPass(b *testing.B) {
	slice := []int{1, 2, 3, 4, 5}
	checker := func(s []int, elem int) bool { return len(s) < 10 }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = AppendIfCheckPass(slice, i, checker)
	}
}

func TestMapTo(t *testing.T) {
	t.Run("int to string", func(t *testing.T) {
		input := []int{1, 2, 3}
		expected := []string{"1", "2", "3"}
		result := MapTo(input, func(i int) string {
			return strconv.Itoa(i)
		})
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("string to int length", func(t *testing.T) {
		input := []string{"a", "bb", "ccc"}
		expected := []int{1, 2, 3}
		result := MapTo(input, func(s string) int {
			return len(s)
		})
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []int{}
		expected := []string{}
		result := MapTo(input, func(i int) string {
			return strconv.Itoa(i)
		})
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
