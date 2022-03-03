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
	ZeroTime     = time.Unix(0, 0)
	TimeLocation = time.Local
)

func SetLocation(local string) error {
	l, err := time.LoadLocation(local)
	if err != nil {
		return err
	}

	TimeLocation = l

	return nil
}

func Parse(str string) (time.Time, error) {
	return time.ParseInLocation(DateTimeLayout, str, TimeLocation)
}

func Milliseconds() int64 {
	return ToMilliseconds(time.Now())
}

func ToMilliseconds(t time.Time) int64 {
	return t.UnixNano() / 1000000
}

func FromMilliseconds(m int64) time.Time {
	return time.Unix(m/1000, m%1000*1000000)
}
