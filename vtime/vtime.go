// Copyright 2019 The vogo Authors. All rights reserved.

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
