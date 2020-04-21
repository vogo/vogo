// Copyright 2019-2020 The vogo Authors. All rights reserved.

package vurl

import (
	"net"
	"net/url"
)

// IsURLNetError whether net error for url request
func IsURLNetError(err error) bool {
	urlErr, ok := err.(*url.Error)
	if !ok {
		return false
	}

	_, ok = urlErr.Err.(*net.OpError)

	return ok
}
