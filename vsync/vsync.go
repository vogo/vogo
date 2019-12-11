// Copyright 2019 The vogo Authors. All rights reserved.

package vsync

func IsChanClosed(c chan struct{}) bool {
	if c == nil {
		return true
	}
	select {
	case <-c:
		return true
	default:
		return false
	}
}

func SafeClosed(c chan struct{}) {
	if !IsChanClosed(c) {
		defer func() {
			_ = recover()
		}()

		close(c)
	}
}
