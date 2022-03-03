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
	"context"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/vogo/logger"
)

var ignoreLoadEnvs = []string{"JAVA_OPTS", "CLASSPATH"}

func isLoadIgnoreEnv(e string) bool {
	for _, env := range ignoreLoadEnvs {
		if env == e {
			return true
		}
	}

	return false
}

func LoadUserEnv() {
	profiles := getUserEnvProfiles()
	for _, profile := range profiles {
		if _, err := os.Stat(profile); err != nil {
			continue
		}

		loadEnvFromProfile(profile)
	}

	adjustPathEnv()
}

func adjustPathEnv() {
	addEnvPathBin("/bin")
	addEnvPathBin("/sbin")
	addEnvPathBin("/usr/bin")
	addEnvPathBin("/usr/sbin")
	addEnvPathBin("/usr/local/bin")
	addEnvPathBin("/usr/local/sbin")
}

func addEnvPathBin(bin string) {
	path := os.Getenv("PATH")
	if !EnvPathContains(path, bin) {
		if err := os.Setenv("PATH", path+EnvValueSplit+bin); err != nil {
			logger.Warnf("set env error: %v", err)
		}
	}
}

func EnvPathContains(path, bin string) bool {
	return strings.HasPrefix(path, bin+EnvValueSplit) ||
		strings.Contains(path, EnvValueSplit+bin+EnvValueSplit) ||
		strings.HasSuffix(path, EnvValueSplit+bin)
}

func loadEnvFromProfile(profile string) {
	logger.Infof("load env from %s", profile)

	commandStr := fmt.Sprintf("source %s && env", profile)
	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, "/bin/sh", "-c", commandStr)

	go func() {
		time.Sleep(time.Second * 2)
		cancel()
	}()

	result, err := cmd.CombinedOutput()
	if err != nil {
		logger.Debugf("load env error: %v", err)

		return
	}

	reg := regexp.MustCompile(`[A-Za-z0-9]+=.*`)
	lines := bytes.Split(result, []byte{'\n'})

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		if !reg.Match(line) {
			continue
		}

		index := bytes.Index(line, []byte{'='})

		key := string(line[:index])
		if isLoadIgnoreEnv(key) {
			continue
		}

		logger.Debugf("set env: %s", line)

		err := os.Setenv(key, string(line[index+1:]))
		if err != nil {
			logger.Errorf("failed to set env: %v", err)
		}
	}
}
