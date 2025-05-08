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

package vrunner

import (
	"sync"
	"sync/atomic"
	"time"
)

type Task func()

// Runner the runner status struct.
type Runner struct {
	// channel to control stop status, stop it by calling Stop().
	C chan struct{}

	done   uint32
	m      sync.Mutex
	defers []Task
}

// Defer add task called in desc order when stopper is stopped.
func (s *Runner) Defer(task Task) {
	s.doSlow(func() {
		s.defers = append(s.defers, task)
	})
}

// doStop do stop work, include closing the chan and calling all defers.
func (s *Runner) doStop() {
	defer atomic.StoreUint32(&s.done, 1)

	close(s.C)

	// call in desc order, like defer.
	for i := len(s.defers) - 1; i >= 0; i-- {
		s.defers[i]()
	}

	// help gc
	s.defers = nil
}

// Stop close the stopper.
func (s *Runner) Stop() {
	s.doSlow(s.doStop)
}

// StopWith stop the stopper and execute the task.
// the same as calling Defer(task) first, and then calling Stop().
func (s *Runner) StopWith(task Task) {
	s.doSlow(func() {
		s.defers = append(s.defers, task)
		s.doStop()
	})
}

// doSlow do func synchronously if the stopper has not been stopped.
// see sync.Once.
func (s *Runner) doSlow(f func()) {
	if atomic.LoadUint32(&s.done) == 0 {
		s.m.Lock()
		defer s.m.Unlock()

		if s.done == 0 {
			f()
		}
	}
}

// Loop run task util the stopper is stopped.
// Note, there is not an interval between the executions of two tasks.
func (s *Runner) Loop(task Task) {
	go func() {
		for {
			select {
			case <-s.C:
				return
			default:
				task()
			}
		}
	}()
}

// Interval run task at intervals util the stopper is stopped.
func (s *Runner) Interval(task Task, interval time.Duration) {
	go func() {
		// run immediately for first time.
		select {
		case <-s.C:
			return
		default:
			task()
		}

		for {
			select {
			case <-s.C:
				return
			case <-time.After(interval):
				task()
			}
		}
	}()
}

// New create a new Runner.
func New() *Runner {
	return &Runner{
		m: sync.Mutex{},
		C: make(chan struct{}),
	}
}

// NewChild create a new Runner as child of the exists chan, when which is closed the child will be stopped too.
func NewChild(stop chan struct{}) *Runner {
	child := &Runner{
		m: sync.Mutex{},
		C: make(chan struct{}),
	}

	go func() {
		select {
		case <-stop:
			child.Stop()
		case <-child.C:
		}
	}()

	return child
}

// NewChild create a new Runner as child of the exists one, when which is stopped the child will be stopped too.
func (s *Runner) NewChild() *Runner {
	return NewChild(s.C)
}

// NewParent create a new Runner as parent of the exists one, which will be stopped when the new parent stopped.
func (s *Runner) NewParent() *Runner {
	parent := &Runner{
		m: sync.Mutex{},
		C: make(chan struct{}),
	}

	go func() {
		select {
		case <-parent.C:
			s.Stop()
		case <-s.C:
		}
	}()

	return parent
}
