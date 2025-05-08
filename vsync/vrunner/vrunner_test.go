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

package vrunner_test

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vogo/vogo/vsync/vrunner"
)

const goroutineScheduleInterval = time.Millisecond * 10

func TestRunnerStop(t *testing.T) {
	t.Parallel()

	s1 := vrunner.New()

	s1.Defer(func() {
		t.Log("s1 stopped 2")
	})
	s1.Defer(func() {
		t.Log("s1 stopped 1")
	})

	// run loop task until s1 closed.
	s1.Loop(func() {
		t.Log("s1 run loop task")
		time.Sleep(time.Millisecond * 3)
	})

	// run interval task until s1 closed.
	s1.Interval(func() {
		t.Log("s1 run interval task")
	}, time.Millisecond)

	go func() {
		ticker := time.NewTicker(time.Millisecond * 2)

		for {
			select {
			case <-s1.C:
				return
			case <-ticker.C:
				t.Log("run ticker task until s1 stopped")
			}
		}
	}()

	s2 := s1.NewChild()
	s2.Defer(func() {
		t.Log("s2 stopped")
	})

	s3 := s2.NewChild()
	s3.Defer(func() {
		t.Log("s3 stopped")
	})

	time.Sleep(goroutineScheduleInterval)

	s1.Stop()

	s1.Interval(func() {
		t.Fatal("should not run interval task after s1 stopped")
	}, time.Millisecond)

	time.Sleep(goroutineScheduleInterval)
}

func TestRunner(t *testing.T) {
	t.Parallel()

	s := vrunner.New()

	var (
		status1 int64
		status2 int64
	)

	s.Defer(func() {
		atomic.StoreInt64(&status1, 1)
	})

	s.Defer(func() {
		atomic.StoreInt64(&status2, 2)
	})

	s.StopWith(func() {
		t.Log("s stopped")
	})

	assert.Equal(t, int64(1), atomic.LoadInt64(&status1))
	assert.Equal(t, int64(2), atomic.LoadInt64(&status2))

	// stop again wont panic
	s.Stop()
}

func TestNewChild(t *testing.T) {
	t.Parallel()

	s := vrunner.New()
	doTestParentChildRunner(t, s, s.NewChild())
}

func TestNewParent(t *testing.T) {
	t.Parallel()

	s := vrunner.New()
	doTestParentChildRunner(t, s.NewParent(), s)
}

func doTestParentChildRunner(t *testing.T, parent, child *vrunner.Runner) {
	t.Helper()

	var (
		status1 int64
		status2 int64
	)

	child.Defer(func() {
		atomic.StoreInt64(&status1, 1)
	})

	parent.Defer(func() {
		atomic.StoreInt64(&status2, 2)
	})

	parent.Stop()

	time.Sleep(goroutineScheduleInterval)

	assert.Equal(t, int64(1), atomic.LoadInt64(&status1))
	assert.Equal(t, int64(2), atomic.LoadInt64(&status2))
}

func TestNewChildFromChan(t *testing.T) {
	t.Parallel()

	c := make(chan struct{})
	s := vrunner.NewChild(c)

	var status1 int64

	s.Defer(func() {
		atomic.AddInt64(&status1, 1)
	})

	time.Sleep(goroutineScheduleInterval)

	close(c)

	time.Sleep(goroutineScheduleInterval)

	assert.Equal(t, int64(1), atomic.LoadInt64(&status1))

	s.Stop()

	assert.Equal(t, int64(1), atomic.LoadInt64(&status1))
}
