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
	"encoding/json"
	"time"
)

type Date struct {
	time.Time
}

func NewDate(t time.Time) *Date {
	return &Date{Time: t}
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Format(DateLayout))
}

func (d *Date) UnmarshalJSON(data []byte) error {
	var dateStr string
	err := json.Unmarshal(data, &dateStr)
	if err != nil {
		return err
	}
	parsed, err := time.Parse(DateLayout, dateStr)
	if err != nil {
		return err
	}
	d.Time = parsed
	return nil
}

func (d Date) String() string {
	return d.Format(DateLayout)
}

func (d *Date) UnmarshalText(data []byte) error {
	parsed, err := time.Parse(DateLayout, string(data))
	if err != nil {
		return err
	}
	d.Time = parsed
	return nil
}
