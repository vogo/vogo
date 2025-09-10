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

// Copyright 2019-2020 The vogo Authors. All rights reserved.

package vioutil_test

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vogo/vogo/vio/vioutil"
)

func TestLinkLatest(t *testing.T) {
	t.Parallel()

	tempDir := os.TempDir()

	sourceFile := filepath.Join(tempDir, "a.txt")
	linkFile := filepath.Join(tempDir, "b.txt")

	assert.Nil(t, os.WriteFile(sourceFile, []byte("test"), 0o600))

	assert.Nil(t, vioutil.LinkFile(sourceFile, linkFile))
	assert.Nil(t, vioutil.LinkFile(sourceFile, linkFile))

	_ = os.Remove(linkFile)
	_ = os.Remove(sourceFile)

	sourceDir := filepath.Join(tempDir, "d1")
	linkDir := filepath.Join(tempDir, "d2")

	_ = os.Mkdir(sourceDir, 0o777)

	assert.Nil(t, vioutil.LinkFile(sourceDir, linkDir))
	assert.Nil(t, vioutil.LinkFile(sourceDir, linkDir))

	_ = os.Remove(sourceDir)
	_ = os.Remove(linkDir)
}

//nolint:dupword // ignore this.
func TestDos2Unix(t *testing.T) {
	t.Parallel()

	fileName := filepath.Join(os.TempDir(), "unit_test_dos2unix.sh")
	_ = os.Remove(fileName)

	file, err := os.Create(fileName)
	if !assert.Nil(t, err) {
		return
	}

	shellData := `#!/bin/bash
PID_FILE="${DIRNAME}/../agent.pid"

if [ -f  "$PID_FILE" ]; then
  if [ ! -z "$PID"  ]; then
    kill -0 $PID >/dev/null 2>&1
    if [ $? -eq 0 ]; then
        echo "process was existed,plz stop it before start!!!"
        exit 1;
    fi
  fi
fi

java \
  -Dsz.framework.projectName=demo-agent \
  -Xms64m -Xmx64m -XX:PermSize=100m -XX:MaxPermSize=100m \
  -Dlog.dir=${DIRNAME}/../../../logs \
  -Dspring.profiles.active=dev \
  $@ \
  -jar  ${DIRNAME}/../agent.jar  &


PID=$!
result=$?
echo $PID > "$PID_FILE"
echo $result

`
	fileData := bytes.ReplaceAll([]byte(shellData), []byte{'\n'}, []byte{'\r', '\n'})

	_, err = file.Write(fileData)
	if !assert.Nil(t, err) {
		return
	}

	if err = file.Close(); err != nil {
		t.Fatalf("Failed to close file: %v", err)
	}

	err = vioutil.Dos2Unix(fileName)
	if !assert.Nil(t, err) {
		return
	}

	b, err := os.ReadFile(fileName)
	if !assert.Nil(t, err) {
		return
	}

	assert.Equal(t, shellData, string(b))

	_ = os.Remove(fileName)
}
