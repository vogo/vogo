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

// Copyright 2019-2020 The vogo Authors. All rights reserved.

package vrand_test

import (
	"testing"
	"time"

	"github.com/vogo/vogo/vrand"
)

const (
	randomSrc = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

func BenchmarkRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		vrand.RandomString(randomSrc, 16)
	}
}

func BenchmarkRandomSeedString(b *testing.B) {
	seed := time.Now().UnixNano()

	for i := 0; i < b.N; i++ {
		vrand.RandomSeedString(seed, randomSrc, 16)
	}
}

func TestIntn64Range(t *testing.T) {
	tests := []struct {
		name string
		min  int64
		max  int64
	}{
		{"positive range", 1, 10},
		{"negative range", -10, -1},
		{"mixed range", -5, 5},
		{"zero range", 0, 10},
		{"single value", 5, 5},
		{"reversed params", 10, 1}, // should swap internally
		{"large range", 1000000, 2000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test multiple iterations to ensure consistency
			for i := 0; i < 100; i++ {
				result := vrand.Intn64Range(tt.min, tt.max)

				// Determine expected min and max (function swaps if min >= max)
				expectedMin, expectedMax := tt.min, tt.max
				if tt.min >= tt.max {
					expectedMin, expectedMax = tt.max, tt.min
				}

				// Check if result is in expected range [min, max]
				if result < expectedMin || result > expectedMax {
					t.Errorf("Intn64Range(%d, %d) = %d, want value in range [%d, %d]",
						tt.min, tt.max, result, expectedMin, expectedMax)
				}
			}
		})
	}
}

func TestIntnRange(t *testing.T) {
	tests := []struct {
		name string
		min  int
		max  int
	}{
		{"positive range", 1, 10},
		{"negative range", -10, -1},
		{"mixed range", -5, 5},
		{"zero range", 0, 10},
		{"single value", 5, 6},     // [5, 6) means only 5 is possible
		{"reversed params", 10, 1}, // should swap internally
		{"large range", 1000, 2000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test multiple iterations to ensure consistency
			for i := 0; i < 100; i++ {
				result := vrand.IntnRange(tt.min, tt.max)

				// Determine expected min and max (function swaps if min >= max)
				expectedMin, expectedMax := tt.min, tt.max
				if tt.min >= tt.max {
					expectedMin, expectedMax = tt.max, tt.min
				}

				// Check if result is in expected range [min, max)
				if result < expectedMin || result >= expectedMax {
					t.Errorf("IntnRange(%d, %d) = %d, want value in range [%d, %d)",
						tt.min, tt.max, result, expectedMin, expectedMax)
				}
			}
		})
	}
}

func TestFloat64Range(t *testing.T) {
	tests := []struct {
		name string
		min  float64
		max  float64
	}{
		{"positive range", 1.0, 10.0},
		{"negative range", -10.0, -1.0},
		{"mixed range", -5.0, 5.0},
		{"zero range", 0.0, 10.0},
		{"small range", 1.0, 1.1},
		{"reversed params", 10.0, 1.0}, // should swap internally
		{"large range", 1000.0, 2000.0},
		{"fractional range", 0.1, 0.9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test multiple iterations to ensure consistency
			for i := 0; i < 100; i++ {
				result := vrand.Float64Range(tt.min, tt.max)

				// Determine expected min and max (function swaps if min >= max)
				expectedMin, expectedMax := tt.min, tt.max
				if tt.min >= tt.max {
					expectedMin, expectedMax = tt.max, tt.min
				}

				// Check if result is in expected range [min, max)
				if result < expectedMin || result >= expectedMax {
					t.Errorf("Float64Range(%f, %f) = %f, want value in range [%f, %f)",
						tt.min, tt.max, result, expectedMin, expectedMax)
				}
			}
		})
	}
}

func TestFloat64Range_EdgeCases(t *testing.T) {
	// Test equal values
	result := vrand.Float64Range(5.0, 5.0)
	if result != 5.0 {
		t.Errorf("Float64Range(5.0, 5.0) = %f, want 5.0", result)
	}

	// Test very small differences
	min, max := 1.0, 1.0000001
	for i := 0; i < 10; i++ {
		result := vrand.Float64Range(min, max)
		if result < min || result >= max {
			t.Errorf("Float64Range(%f, %f) = %f, want value in range [%f, %f)",
				min, max, result, min, max)
		}
	}
}
