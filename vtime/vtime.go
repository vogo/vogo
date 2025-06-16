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

package vtime

import "time"

const (
	DateTimeLayout = "2006-01-02 15:04:05"
)

var (
	// ZeroTime is the zero time.
	ZeroTime     = time.Unix(0, 0)
	TimeLocation = time.Local
)

// SetLocation sets the time location.
func SetLocation(local string) error {
	l, err := time.LoadLocation(local)
	if err != nil {
		return err
	}

	TimeLocation = l

	return nil
}

// Parse parses a string to time.
func Parse(str string) (time.Time, error) {
	return time.ParseInLocation(DateTimeLayout, str, TimeLocation)
}

// Milliseconds returns the current time in milliseconds.
func Milliseconds() int64 {
	return time.Now().UnixMilli()
}

// ToMilliseconds converts a time to milliseconds.
// Deprecated: use time.Time.UnixMilli() instead.
func ToMilliseconds(t time.Time) int64 {
	return t.UnixMilli()
}

// FromMilliseconds converts milliseconds to time.
// Deprecated: use time.UnixMilli(m) instead.
func FromMilliseconds(m int64) time.Time {
	return time.UnixMilli(m)
}

const (
	// DateHourMinuteLayout is the layout for service time
	DateHourMinuteLayout = "2006-01-02 15:04"
	DateLayout           = "2006-01-02"
	HourMinuteLayout     = "15:04"
)

func ParseDateHourMinute(t string) (time.Time, error) {
	if len(t) == len(DateHourMinuteLayout)+3 {
		t = t[:len(DateHourMinuteLayout)]
	}

	return time.Parse(DateHourMinuteLayout, t)
}

func FormatDateHourMinute(t time.Time) string {
	return t.Format(DateHourMinuteLayout)
}

func FormatHourMinute(t time.Time) string {
	return t.Format(HourMinuteLayout)
}

func FormatDate(t time.Time) string {
	return t.Format(DateLayout)
}

func FormatDateTime(t time.Time) string {
	return t.Format(DateTimeLayout)
}

func DateTimeString(t *time.Time) string {
	if t == nil || t.IsZero() {
		return ""
	}

	return t.Format(DateTimeLayout)
}

func DateString(t *time.Time) string {
	if t == nil || t.IsZero() {
		return ""
	}

	return t.Format(DateLayout)
}

func HourMinuteString(t *time.Time) string {
	if t == nil || t.IsZero() {
		return ""
	}

	return t.Format(HourMinuteLayout)
}

func DateHourMinuteString(t *time.Time) string {
	if t == nil || t.IsZero() {
		return ""
	}

	return t.Format(DateHourMinuteLayout)
}
