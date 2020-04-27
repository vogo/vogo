// Copyright 2019-2020 The vogo Authors. All rights reserved.

package vos_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vogo/vogo/vos"
)

func TestGetProcessUser(t *testing.T) {
	u, err := vos.GetProcessUser(os.Getpid())
	assert.Nil(t, err)
	t.Logf("user: %s", u)
}
