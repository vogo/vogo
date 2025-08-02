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

package vstrings

import "testing"

func TestContainsIn(t *testing.T) {
	tests := []struct {
		name     string
		items    []string
		item     string
		expected bool
	}{
		{"empty slice", []string{}, "test", false},
		{"item exists", []string{"a", "b", "c"}, "b", true},
		{"item not exists", []string{"a", "b", "c"}, "d", false},
		{"empty string in slice", []string{"a", "", "c"}, "", true},
		{"empty string not in slice", []string{"a", "b", "c"}, "", false},
		{"single item match", []string{"test"}, "test", true},
		{"single item no match", []string{"test"}, "other", false},
		{"case sensitive", []string{"Test", "test"}, "TEST", false},
		{"duplicate items", []string{"a", "b", "a", "c"}, "a", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ContainsIn(tt.items, tt.item)
			if result != tt.expected {
				t.Errorf("Expected %t, got %t", tt.expected, result)
			}
		})
	}
}

func TestContainsAny(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		test     []string
		expected bool
	}{
		{"no test strings", "hello world", []string{}, false},
		{"single match", "hello world", []string{"world"}, true},
		{"multiple matches", "hello world", []string{"hello", "world"}, true},
		{"no matches", "hello world", []string{"foo", "bar"}, false},
		{"partial match", "hello world", []string{"ell"}, true},
		{"empty string in test", "hello world", []string{""}, true},
		{"empty source string", "", []string{"test"}, false},
		{"empty source with empty test", "", []string{""}, true},
		{"case sensitive", "Hello World", []string{"hello"}, false},
		{"first match wins", "hello world", []string{"hello", "xyz"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ContainsAny(tt.s, tt.test...)
			if result != tt.expected {
				t.Errorf("Expected %t, got %t", tt.expected, result)
			}
		})
	}
}

func TestAfterFirst(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		sep      string
		expected string
	}{
		{"normal case", "hello-world-test", "-", "world-test"},
		{"separator not found", "hello world", "-", "hello world"},
		{"separator at beginning", "-hello world", "-", "hello world"},
		{"separator at end", "hello world-", "-", ""},
		{"empty string", "", "-", ""},
		{"empty separator", "hello world", "", "hello world"},
		{"multiple separators", "a-b-c-d", "-", "b-c-d"},
		{"separator longer than string", "hi", "hello", "hi"},
		{"multi-char separator", "hello::world::test", "::", "world::test"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AfterFirst(tt.s, tt.sep)
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestAfterLast(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		sep      string
		expected string
	}{
		{"normal case", "hello-world-test", "-", "test"},
		{"separator not found", "hello world", "-", "hello world"},
		{"separator at beginning", "-hello world", "-", "hello world"},
		{"separator at end", "hello world-", "-", ""},
		{"empty string", "", "-", ""},
		{"empty separator", "hello world", "", ""},
		{"multiple separators", "a-b-c-d", "-", "d"},
		{"single character", "a", "a", ""},
		{"multi-char separator", "hello::world::test", "::", "test"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AfterLast(tt.s, tt.sep)
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestBeforeFirst(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		sep      string
		expected string
	}{
		{"normal case", "hello-world-test", "-", "hello"},
		{"separator not found", "hello world", "-", "hello world"},
		{"separator at beginning", "-hello world", "-", ""},
		{"separator at end", "hello world-", "-", "hello world"},
		{"empty string", "", "-", ""},
		{"empty separator", "hello world", "", ""},
		{"multiple separators", "a-b-c-d", "-", "a"},
		{"single character", "a", "a", ""},
		{"multi-char separator", "hello::world::test", "::", "hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BeforeFirst(tt.s, tt.sep)
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestBeforeFirstInclude(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		sep      string
		expected string
	}{
		{"normal case", "hello-world-test", "-", "hello-"},
		{"separator not found", "hello world", "-", "hello world"},
		{"separator at beginning", "-hello world", "-", "-"},
		{"separator at end", "hello world-", "-", "hello world-"},
		{"empty string", "", "-", ""},
		{"empty separator", "hello world", "", ""},
		{"multiple separators", "a-b-c-d", "-", "a-"},
		{"single character", "a", "a", "a"},
		{"multi-char separator", "hello::world::test", "::", "hello::"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BeforeFirstInclude(tt.s, tt.sep)
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestBeforeLast(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		sep      string
		expected string
	}{
		{"normal case", "hello-world-test", "-", "hello-world"},
		{"separator not found", "hello world", "-", "hello world"},
		{"separator at beginning", "-hello world", "-", ""},
		{"separator at end", "hello world-", "-", "hello world"},
		{"empty string", "", "-", ""},
		{"empty separator", "hello world", "", "hello world"},
		{"multiple separators", "a-b-c-d", "-", "a-b-c"},
		{"single character", "a", "a", ""},
		{"multi-char separator", "hello::world::test", "::", "hello::world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BeforeLast(tt.s, tt.sep)
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestBeforeLastInclude(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		sep      string
		expected string
	}{
		{"normal case", "hello-world-test", "-", "hello-world-"},
		{"separator not found", "hello world", "-", "hello world"},
		{"separator at beginning", "-hello world", "-", "-"},
		{"separator at end", "hello world-", "-", "hello world-"},
		{"empty string", "", "-", ""},
		{"empty separator", "hello world", "", "hello world"},
		{"multiple separators", "a-b-c-d", "-", "a-b-c-"},
		{"single character", "a", "a", "a"},
		{"multi-char separator", "hello::world::test", "::", "hello::world::"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BeforeLastInclude(tt.s, tt.sep)
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

// Benchmark tests
func BenchmarkContainsIn(b *testing.B) {
	items := []string{"apple", "banana", "cherry", "date", "elderberry"}
	b.ResetTimer()

	b.Run("found", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ContainsIn(items, "cherry")
		}
	})

	b.Run("not found", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ContainsIn(items, "grape")
		}
	})

	b.Run("empty slice", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ContainsIn([]string{}, "test")
		}
	})
}

func BenchmarkContainsAny(b *testing.B) {
	s := "hello world this is a test string"
	tests := []string{"test", "example", "sample"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = ContainsAny(s, tests...)
	}
}

func BenchmarkAfterFirst(b *testing.B) {
	s := "path/to/some/file.txt"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = AfterFirst(s, "/")
	}
}

func BenchmarkBeforeLast(b *testing.B) {
	s := "path/to/some/file.txt"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = BeforeLast(s, ".")
	}
}

// Example tests
func ExampleContainsIn() {
	items := []string{"apple", "banana", "cherry"}
	result := ContainsIn(items, "banana")
	_ = result // true

	result = ContainsIn(items, "grape")
	_ = result // false
}

func ExampleContainsAny() {
	s := "hello world"
	result := ContainsAny(s, "world", "test")
	_ = result // true

	result = ContainsAny(s, "foo", "bar")
	_ = result // false
}

func ExampleAfterFirst() {
	result := AfterFirst("path/to/file.txt", "/")
	_ = result // "to/file.txt"

	result = AfterFirst("no-separator", "/")
	_ = result // "no-separator"
}

func ExampleAfterLast() {
	result := AfterLast("path/to/file.txt", "/")
	_ = result // "file.txt"

	result = AfterLast("file.txt", ".")
	_ = result // "txt"
}

func ExampleBeforeFirst() {
	result := BeforeFirst("path/to/file.txt", "/")
	_ = result // "path"

	result = BeforeFirst("no-separator", "/")
	_ = result // "no-separator"
}

func ExampleBeforeLast() {
	result := BeforeLast("path/to/file.txt", "/")
	_ = result // "path/to"

	result = BeforeLast("file.txt", ".")
	_ = result // "file"
}