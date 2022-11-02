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

package vos

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"

	"github.com/vogo/logger"
	"github.com/vogo/vogo/vstrings"
)

var (
	ErrPortNotFound = errors.New("port not found")
	ErrNoProcess    = errors.New("no process")
)

// PidExist whether pid exists
// see: https://stackoverflow.com/questions/15204162/check-if-a-process-exists-in-go-way
func PidExist(pid int) bool {
	p, err := os.FindProcess(pid)

	return err == nil && p.Signal(syscall.Signal(0)) == nil
}

// Kill process.
func Kill(pid int) error {
	logger.Infof("kill process %d", pid)

	proc, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	if proc == nil {
		return nil
	}

	return proc.Kill()
}

// GetPidByPort get process pid by port.
// command example: netstat -anp|grep '8888 ' |grep 'LISTEN'|awk '{printf $7}'|cut -d/ -f1 .
func GetPidByPort(port int) (int, error) {
	fullCommand := fmt.Sprintf("lsof -iTCP:%d -sTCP:LISTEN -n -P |grep LISTEN | awk '{print $2}'", port)

	result, err := ExecShell(fullCommand)
	if err != nil {
		return 0, fmt.Errorf("command error: %w", err)
	}

	if result == nil {
		return 0, fmt.Errorf("%w: starting at %d", ErrNoProcess, port)
	}

	lines := bytes.Split(result, []byte{'\n'})
	pid := lines[len(lines)-1]

	if len(pid) == 0 && len(lines) > 1 {
		pid = lines[len(lines)-2]
	}

	logger.Debugf("the pid of port %d is %s", port, pid)

	p, err := strconv.Atoi(string(pid))
	if err != nil {
		logger.Warnf("can't find pid for port %d, result: %s", port, pid)

		return -1, ErrPortNotFound
	}

	return p, nil
}

// GetProcessUser the user of process by pid
// example: ps -o ruser -p 16787 | tail -1
func GetProcessUser(pid int) (string, error) {
	fullCommand := fmt.Sprintf("ps -o ruser -p %d | tail -1", pid)

	return SingleCommandResult(fullCommand)
}

func ReadProcEnv(pid []byte) map[string]string {
	env := make(map[string]string)

	environData, err := os.ReadFile(fmt.Sprintf("/proc/%s/environ", pid))
	if err != nil {
		return env
	}

	environ := bytes.Split(environData, []byte{0x0})

	for _, e := range environ {
		items := strings.SplitN(string(e), "=", 2)
		if len(items) > 1 && vstrings.ContainsAny(items[0], "JAVA", "JRE", "PATH", "CATALINA", "USER", "HOME") {
			env[items[0]] = items[1]
		}
	}

	return env
}
