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
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/vogo/vogo/vlog"
	"github.com/vogo/vogo/vos/vuser"
)

func EnsureEnvString(key string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		panic("env not found: " + key)
	}

	return v
}

func EnvString(key string) string {
	v, _ := os.LookupEnv(key)

	return v
}

func GetEnvStr(key string, defaultValue string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return v
}

func EnsureEnvInt(key string) int {
	v, ok := os.LookupEnv(key)
	if !ok {
		panic("env not found: " + key)
	}

	intValue, err := strconv.Atoi(v)
	if err != nil {
		panic(fmt.Sprintf("invalid config %s: %s", key, v))
	}

	return intValue
}

func EnvInt(key string) int {
	v, _ := os.LookupEnv(key)
	intValue, err := strconv.Atoi(v)
	if err != nil {
		return 0
	}
	return intValue
}

func GetEnvInt(key string, defaultValue int) int {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	intValue, err := strconv.Atoi(v)
	if err != nil {
		return defaultValue
	}

	return intValue
}

func EnsureEnvInt64(key string) int64 {
	return int64(EnsureEnvInt(key))
}

func EnvInt64(key string) int64 {
	return int64(EnvInt(key))
}

func GetEnvInt64(key string, defaultValue int64) int64 {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	intValue, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return defaultValue
	}

	return intValue
}

func EnvBool(key string) bool {
	return GetEnvStr(key, "false") == "true"
}

func GetEnvBool(key string, defaultValue bool) bool {
	return GetEnvStr(key, strconv.FormatBool(defaultValue)) == "true"
}

func EnsureEnvBool(key string) bool {
	return EnsureEnvString(key) == "true"
}

var ignoreLoadEnvs = []string{"JAVA_OPTS", "CLASSPATH"}

func isLoadIgnoreEnv(e string) bool {
	return slices.Contains(ignoreLoadEnvs, e)
}

func LoadUserEnv() {
	profiles := vuser.GetUserEnvProfiles()
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
		if err := os.Setenv("PATH", path+vuser.EnvValueSplit+bin); err != nil {
			vlog.Printf("set env error | err: %v", err)
		}
	}
}

func EnvPathContains(path, bin string) bool {
	return strings.HasPrefix(path, bin+vuser.EnvValueSplit) ||
		strings.Contains(path, vuser.EnvValueSplit+bin+vuser.EnvValueSplit) ||
		strings.HasSuffix(path, vuser.EnvValueSplit+bin)
}

func loadEnvFromProfile(profile string) {
	commandStr := fmt.Sprintf("source %s && env", profile)
	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, "/bin/sh", "-c", commandStr)

	go func() {
		time.Sleep(time.Second * 2)
		cancel()
	}()

	result, err := cmd.CombinedOutput()
	if err != nil {
		vlog.Printf("load env error | err: %v", err)

		return
	}

	reg := regexp.MustCompile(`[A-Za-z0-9]+=.*`)
	lines := bytes.SplitSeq(result, []byte{'\n'})

	for line := range lines {
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

		err = os.Setenv(key, string(line[index+1:]))
		if err != nil {
			vlog.Printf("failed to set env | err: %v", err)
		}
	}
}
