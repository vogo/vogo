// Copyright 2019-2020 The vogo Authors. All rights reserved.

package vmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoundFloat64(t *testing.T) {
	assert.Equal(t, 12.23, RoundFloat64(12.2345678, 2))
	assert.Equal(t, 12.235, RoundFloat64(12.2345678, 3))
	assert.Equal(t, 12.2346, RoundFloat64(12.2345678, 4))
	assert.Equal(t, 12.23457, RoundFloat64(12.2345678, 5))
	assert.Equal(t, 12.23456, RoundFloat64(12.2345632, 5))
	assert.Equal(t, 12.234563, RoundFloat64(12.2345632, 6))
	assert.Equal(t, 12.2345632, RoundFloat64(12.2345632, 7))
	assert.Equal(t, 12.2345632, RoundFloat64(12.2345632, 8))
	assert.Equal(t, 12.2345632, RoundFloat64(12.2345632, 9))
}
