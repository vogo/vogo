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

package vstrconv

import (
	"math"
	"testing"
)

// Test Ensure functions (should panic on invalid input)
func TestEnsureInt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
		shouldPanic bool
	}{
		{"valid positive", "123", 123, false},
		{"valid negative", "-456", -456, false},
		{"valid zero", "0", 0, false},
		{"invalid string", "abc", 0, true},
		{"empty string", "", 0, true},
		{"float string", "123.45", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Expected panic for input %q", tt.input)
					}
				}()
			}
			result := EnsureInt(tt.input)
			if !tt.shouldPanic && result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestEnsureInt64(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int64
		shouldPanic bool
	}{
		{"valid positive", "9223372036854775807", 9223372036854775807, false},
		{"valid negative", "-9223372036854775808", -9223372036854775808, false},
		{"valid zero", "0", 0, false},
		{"invalid string", "abc", 0, true},
		{"empty string", "", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Expected panic for input %q", tt.input)
					}
				}()
			}
			result := EnsureInt64(tt.input)
			if !tt.shouldPanic && result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestEnsureInt32(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int32
		shouldPanic bool
	}{
		{"valid positive", "2147483647", 2147483647, false},
		{"valid negative", "-2147483648", -2147483648, false},
		{"valid zero", "0", 0, false},
		{"invalid string", "abc", 0, true},
		{"overflow", "2147483648", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Expected panic for input %q", tt.input)
					}
				}()
			}
			result := EnsureInt32(tt.input)
			if !tt.shouldPanic && result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestEnsureUint(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected uint
		shouldPanic bool
	}{
		{"valid positive", "123", 123, false},
		{"valid zero", "0", 0, false},
		{"invalid string", "abc", 0, true},
		{"negative number", "-1", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Expected panic for input %q", tt.input)
					}
				}()
			}
			result := EnsureUint(tt.input)
			if !tt.shouldPanic && result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestEnsureUint64(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected uint64
		shouldPanic bool
	}{
		{"valid positive", "18446744073709551615", 18446744073709551615, false},
		{"valid zero", "0", 0, false},
		{"invalid string", "abc", 0, true},
		{"negative number", "-1", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Expected panic for input %q", tt.input)
					}
				}()
			}
			result := EnsureUint64(tt.input)
			if !tt.shouldPanic && result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestEnsureBool(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
		shouldPanic bool
	}{
		{"true", "true", true, false},
		{"false", "false", false, false},
		{"1", "1", true, false},
		{"0", "0", false, false},
		{"invalid string", "abc", false, true},
		{"empty string", "", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Expected panic for input %q", tt.input)
					}
				}()
			}
			result := EnsureBool(tt.input)
			if !tt.shouldPanic && result != tt.expected {
				t.Errorf("Expected %t, got %t", tt.expected, result)
			}
		})
	}
}

func TestEnsureFloat64(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
		shouldPanic bool
	}{
		{"valid positive", "123.45", 123.45, false},
		{"valid negative", "-123.45", -123.45, false},
		{"valid zero", "0", 0, false},
		{"valid integer", "123", 123, false},
		{"invalid string", "abc", 0, true},
		{"empty string", "", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Expected panic for input %q", tt.input)
					}
				}()
			}
			result := EnsureFloat64(tt.input)
			if !tt.shouldPanic && result != tt.expected {
				t.Errorf("Expected %f, got %f", tt.expected, result)
			}
		})
	}
}

func TestEnsureFloat32(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float32
		shouldPanic bool
	}{
		{"valid positive", "123.45", 123.45, false},
		{"valid negative", "-123.45", -123.45, false},
		{"valid zero", "0", 0, false},
		{"invalid string", "abc", 0, true},
		{"empty string", "", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Expected panic for input %q", tt.input)
					}
				}()
			}
			result := EnsureFloat32(tt.input)
			if !tt.shouldPanic && math.Abs(float64(result-tt.expected)) > 1e-6 {
				t.Errorf("Expected %f, got %f", tt.expected, result)
			}
		})
	}
}

// Test non-Ensure functions (should return default values for empty strings)
func TestInt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"valid positive", "123", 123},
		{"valid negative", "-456", -456},
		{"valid zero", "0", 0},
		{"empty string", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Int(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int64
	}{
		{"valid positive", "123", 123},
		{"valid negative", "-456", -456},
		{"valid zero", "0", 0},
		{"empty string", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Int64(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestInt32(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int32
	}{
		{"valid positive", "123", 123},
		{"valid negative", "-456", -456},
		{"valid zero", "0", 0},
		{"empty string", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Int32(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestUint(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected uint
	}{
		{"valid positive", "123", 123},
		{"valid zero", "0", 0},
		{"empty string", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Uint(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestUint64(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected uint64
	}{
		{"valid positive", "123", 123},
		{"valid zero", "0", 0},
		{"empty string", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Uint64(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestBool(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"true", "true", true},
		{"false", "false", false},
		{"1", "1", true},
		{"0", "0", false},
		{"empty string", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Bool(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %t, got %t", tt.expected, result)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
	}{
		{"valid positive", "123.45", 123.45},
		{"valid negative", "-123.45", -123.45},
		{"valid zero", "0", 0},
		{"empty string", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Float64(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %f, got %f", tt.expected, result)
			}
		})
	}
}

func TestFloat32(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float32
	}{
		{"valid positive", "123.45", 123.45},
		{"valid negative", "-123.45", -123.45},
		{"valid zero", "0", 0},
		{"empty string", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Float32(tt.input)
			if math.Abs(float64(result-tt.expected)) > 1e-6 {
				t.Errorf("Expected %f, got %f", tt.expected, result)
			}
		})
	}
}

// Test formatting functions
func TestI64toa(t *testing.T) {
	tests := []struct {
		name     string
		input    int64
		expected string
	}{
		{"positive", 123, "123"},
		{"negative", -456, "-456"},
		{"zero", 0, "0"},
		{"max int64", 9223372036854775807, "9223372036854775807"},
		{"min int64", -9223372036854775808, "-9223372036854775808"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := I64toa(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestI32toa(t *testing.T) {
	tests := []struct {
		name     string
		input    int32
		expected string
	}{
		{"positive", 123, "123"},
		{"negative", -456, "-456"},
		{"zero", 0, "0"},
		{"max int32", 2147483647, "2147483647"},
		{"min int32", -2147483648, "-2147483648"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := I32toa(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestFtoa(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected string
	}{
		{"positive", 123.45, "123.45"},
		{"negative", -123.45, "-123.45"},
		{"zero", 0.0, "0"},
		{"integer", 123.0, "123"},
		{"small decimal", 0.123, "0.123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Ftoa(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestF32toa(t *testing.T) {
	tests := []struct {
		name     string
		input    float32
		expected string
	}{
		{"positive", 123.45, "123.45"},
		{"negative", -123.45, "-123.45"},
		{"zero", 0.0, "0"},
		{"integer", 123.0, "123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := F32toa(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

// Benchmark tests
func BenchmarkEnsureInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = EnsureInt("123")
	}
}

func BenchmarkInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Int("123")
	}
}

func BenchmarkIntEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Int("")
	}
}

func BenchmarkI64toa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = I64toa(123456789)
	}
}

// Example tests
func ExampleEnsureInt() {
	// Valid input
	result := EnsureInt("123")
	_ = result // 123

	// Invalid input will panic
	// EnsureInt("abc") // This will panic
}

func ExampleInt() {
	// Valid input
	result := Int("123")
	_ = result // 123

	// Empty string returns default value
	result = Int("")
	_ = result // 0
}

func ExampleI64toa() {
	result := I64toa(123456789)
	_ = result // "123456789"
}