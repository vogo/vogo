// Copyright 2019-2020 The vogo Authors. All rights reserved.

package vioutil_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vogo/vogo/vio/vioutil"
)

func TestLinkLatest(t *testing.T) {
	tempDir := os.TempDir()

	sourceFile := filepath.Join(tempDir, "a.txt")
	linkFile := filepath.Join(tempDir, "b.txt")

	assert.Nil(t, ioutil.WriteFile(sourceFile, []byte("test"), 0666))

	assert.Nil(t, vioutil.LinkFile(sourceFile, linkFile))
	assert.Nil(t, vioutil.LinkFile(sourceFile, linkFile))

	_ = os.Remove(linkFile)
	_ = os.Remove(sourceFile)

	sourceDir := filepath.Join(tempDir, "d1")
	linkDir := filepath.Join(tempDir, "d2")

	_ = os.Mkdir(sourceDir, 0777)

	assert.Nil(t, vioutil.LinkFile(sourceDir, linkDir))
	assert.Nil(t, vioutil.LinkFile(sourceDir, linkDir))

	_ = os.Remove(sourceDir)
	_ = os.Remove(linkDir)
}

func TestDos2Unix(t *testing.T) {
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

	file.Close()

	err = vioutil.Dos2Unix(fileName)
	if !assert.Nil(t, err) {
		return
	}

	b, err := ioutil.ReadFile(fileName)
	if !assert.Nil(t, err) {
		return
	}

	assert.Equal(t, shellData, string(b))

	_ = os.Remove(fileName)
}
