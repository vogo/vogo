// Copyright 2019 The vogo Authors. All rights reserved.

package vsnowflake

import "testing"

func TestSonySnowflake_NextID(t *testing.T) {
	snowflake := New()

	t.Logf("machine id: %d", machineIDFromIP)

	for i := 0; i < 10; i++ {
		t.Log(snowflake.NextID())
	}
}
