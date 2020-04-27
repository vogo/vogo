// Copyright 2019-2020 The vogo Authors. All rights reserved.

package vjson

import "encoding/json"

// EnsureUnmarshal unmarshal data and panic if has error.
func EnsureUnmarshal(data []byte, v interface{}) {
	if err := json.Unmarshal(data, v); err != nil {
		panic(err)
	}
}

// EnsureMarshal marshal interface and panic if has error.
func EnsureMarshal(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return b
}
