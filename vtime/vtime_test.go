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

package vtime

import (
	"testing"
	"time"
)

func TestSetLocation(t *testing.T) {
	tests := []struct {
		name        string
		location    string
		expectError bool
	}{
		{"valid UTC", "UTC", false},
		{"valid Local", "Local", false},
		{"valid timezone", "America/New_York", false},
		{"valid timezone Asia", "Asia/Shanghai", false},
		{"invalid timezone", "Invalid/Timezone", true},
		{"empty string", "", false}, // LoadLocation("") returns Local, which is valid
	}

	// Save original location
	originalLocation := TimeLocation
	defer func() {
		TimeLocation = originalLocation
	}()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := SetLocation(tt.location)
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error for location %q, got nil", tt.location)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error for location %q, got %v", tt.location, err)
				}
				// For empty string, Go's LoadLocation("") returns UTC
				expectedLocation := tt.location
				if tt.location == "" {
					expectedLocation = "UTC"
				}
				if TimeLocation.String() != expectedLocation {
					t.Errorf("Expected location %q, got %q", expectedLocation, TimeLocation.String())
				}
			}
		})
	}
}

func TestParse(t *testing.T) {
	// Set location to UTC for consistent testing
	originalLocation := TimeLocation
	TimeLocation = time.UTC
	defer func() {
		TimeLocation = originalLocation
	}()

	tests := []struct {
		name        string
		input       string
		expected    time.Time
		expectError bool
	}{
		{
			"valid datetime",
			"2023-12-25 15:30:45",
			time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC),
			false,
		},
		{
			"valid datetime with zeros",
			"2023-01-01 00:00:00",
			time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			false,
		},
		{"invalid format", "2023/12/25 15:30:45", time.Time{}, true},
		{"invalid date", "2023-13-25 15:30:45", time.Time{}, true},
		{"invalid time", "2023-12-25 25:30:45", time.Time{}, true},
		{"empty string", "", time.Time{}, true},
		{"partial string", "2023-12-25", time.Time{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Parse(tt.input)
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error for input %q, got nil", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error for input %q, got %v", tt.input, err)
				}
				if !result.Equal(tt.expected) {
					t.Errorf("Expected %v, got %v", tt.expected, result)
				}
			}
		})
	}
}

func TestMilliseconds(t *testing.T) {
	// Test that Milliseconds returns a reasonable value
	before := time.Now().UnixMilli()
	result := Milliseconds()
	after := time.Now().UnixMilli()

	if result < before || result > after {
		t.Errorf("Milliseconds() returned %d, expected between %d and %d", result, before, after)
	}
}

func TestToMilliseconds(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected int64
	}{
		{
			"epoch",
			time.Unix(0, 0),
			0,
		},
		{
			"specific time",
			time.Date(2023, 12, 25, 15, 30, 45, 123000000, time.UTC),
			time.Date(2023, 12, 25, 15, 30, 45, 123000000, time.UTC).UnixMilli(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToMilliseconds(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestFromMilliseconds(t *testing.T) {
	tests := []struct {
		name     string
		input    int64
		expected time.Time
	}{
		{
			"epoch",
			0,
			time.Unix(0, 0),
		},
		{
			"specific time",
			time.Date(2023, 12, 25, 15, 30, 45, 123000000, time.UTC).UnixMilli(),
			time.Date(2023, 12, 25, 15, 30, 45, 123000000, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FromMilliseconds(tt.input)
			if !result.Equal(tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestParseDateHourMinute(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    time.Time
		expectError bool
	}{
		{
			"valid format",
			"2023-12-25 15:30",
			time.Date(2023, 12, 25, 15, 30, 0, 0, time.UTC),
			false,
		},
		{
			"with seconds (should be truncated)",
			"2023-12-25 15:30:45",
			time.Date(2023, 12, 25, 15, 30, 0, 0, time.UTC),
			false,
		},
		{"invalid format", "2023/12/25 15:30", time.Time{}, true},
		{"invalid date", "2023-13-25 15:30", time.Time{}, true},
		{"invalid time", "2023-12-25 25:30", time.Time{}, true},
		{"empty string", "", time.Time{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseDateHourMinute(tt.input)
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error for input %q, got nil", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error for input %q, got %v", tt.input, err)
				}
				if !result.Equal(tt.expected) {
					t.Errorf("Expected %v, got %v", tt.expected, result)
				}
			}
		})
	}
}

func TestFormatDateHourMinute(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected string
	}{
		{
			"normal time",
			time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC),
			"2023-12-25 15:30",
		},
		{
			"zero time",
			time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			"2023-01-01 00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatDateHourMinute(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestFormatHourMinute(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected string
	}{
		{
			"afternoon",
			time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC),
			"15:30",
		},
		{
			"midnight",
			time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			"00:00",
		},
		{
			"noon",
			time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
			"12:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatHourMinute(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestFormatDate(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected string
	}{
		{
			"normal date",
			time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC),
			"2023-12-25",
		},
		{
			"new year",
			time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			"2023-01-01",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatDate(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestFormatDateTime(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected string
	}{
		{
			"normal datetime",
			time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC),
			"2023-12-25 15:30:45",
		},
		{
			"zero datetime",
			time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			"2023-01-01 00:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatDateTime(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestDateTimeString(t *testing.T) {
	tests := []struct {
		name     string
		input    *time.Time
		expected string
	}{
		{
			"normal time",
			&[]time.Time{time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC)}[0],
			"2023-12-25 15:30:45",
		},
		{"nil pointer", nil, ""},
		{"zero time", &[]time.Time{{}}[0], ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DateTimeString(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestDateString(t *testing.T) {
	tests := []struct {
		name     string
		input    *time.Time
		expected string
	}{
		{
			"normal time",
			&[]time.Time{time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC)}[0],
			"2023-12-25",
		},
		{"nil pointer", nil, ""},
		{"zero time", &[]time.Time{{}}[0], ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DateString(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestHourMinuteString(t *testing.T) {
	tests := []struct {
		name     string
		input    *time.Time
		expected string
	}{
		{
			"normal time",
			&[]time.Time{time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC)}[0],
			"15:30",
		},
		{"nil pointer", nil, ""},
		{"zero time", &[]time.Time{{}}[0], ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HourMinuteString(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestDateHourMinuteString(t *testing.T) {
	tests := []struct {
		name     string
		input    *time.Time
		expected string
	}{
		{
			"normal time",
			&[]time.Time{time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC)}[0],
			"2023-12-25 15:30",
		},
		{"nil pointer", nil, ""},
		{"zero time", &[]time.Time{{}}[0], ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DateHourMinuteString(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

// Test constants and variables
func TestConstants(t *testing.T) {
	if DateTimeLayout != "2006-01-02 15:04:05" {
		t.Errorf("Expected DateTimeLayout to be %q, got %q", "2006-01-02 15:04:05", DateTimeLayout)
	}

	if DateHourMinuteLayout != "2006-01-02 15:04" {
		t.Errorf("Expected DateHourMinuteLayout to be %q, got %q", "2006-01-02 15:04", DateHourMinuteLayout)
	}

	if DateLayout != "2006-01-02" {
		t.Errorf("Expected DateLayout to be %q, got %q", "2006-01-02", DateLayout)
	}

	if HourMinuteLayout != "15:04" {
		t.Errorf("Expected HourMinuteLayout to be %q, got %q", "15:04", HourMinuteLayout)
	}

	if !ZeroTime.Equal(time.Unix(0, 0)) {
		t.Errorf("Expected ZeroTime to be %v, got %v", time.Unix(0, 0), ZeroTime)
	}
}

// Benchmark tests
func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Parse("2023-12-25 15:30:45")
	}
}

func BenchmarkFormatDateTime(b *testing.B) {
	t := time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = FormatDateTime(t)
	}
}

func BenchmarkMilliseconds(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Milliseconds()
	}
}

// Example tests
func ExampleParse() {
	t, err := Parse("2023-12-25 15:30:45")
	if err != nil {
		// handle error
	}
	_ = t
}

func ExampleFormatDateTime() {
	t := time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC)
	result := FormatDateTime(t)
	_ = result // "2023-12-25 15:30:45"
}

func ExampleSetLocation() {
	err := SetLocation("Asia/Shanghai")
	if err != nil {
		// handle error
	}
	// TimeLocation is now set to Asia/Shanghai
}
