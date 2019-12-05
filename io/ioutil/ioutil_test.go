//author: wongoo
//date: 20190628

package vioutil

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkLatest(t *testing.T) {
	tempDir := os.TempDir()
	targetDir := filepath.Join(tempDir, "test_link_latest")
	err := os.Mkdir(targetDir, os.ModePerm)
	assert.Nil(t, err)

	prefix := "temp-test-file"
	suffix := ".txt"

	versionFile1 := filepath.Join(tempDir, prefix+"-1.0"+suffix)
	versionFile2 := filepath.Join(tempDir, prefix+"-2.0"+suffix)

	err = ioutil.WriteFile(versionFile1, []byte("test"), 0666)
	assert.Nil(t, err)

	err = ioutil.WriteFile(versionFile2, []byte("test"), 0666)
	assert.Nil(t, err)

	// -----> create first
	err = LinkLatest(tempDir, targetDir, prefix, suffix)
	assert.Nil(t, err)

	latestFile := filepath.Join(targetDir, prefix+"-latest"+suffix)
	_, err = os.Lstat(latestFile)
	assert.Nil(t, err)

	// -----> create again
	err = LinkLatest(tempDir, targetDir, prefix, suffix)
	assert.Nil(t, err)

	latestFile = filepath.Join(targetDir, prefix+"-latest"+suffix)
	_, err = os.Lstat(latestFile)
	assert.Nil(t, err)

	_ = os.Remove(latestFile)
	_ = os.Remove(versionFile1)
	_ = os.Remove(versionFile2)
	_ = os.Remove(targetDir)
}

func TestParsePackageNameVersion(t *testing.T) {
	name, version, ok := ParsePackageNameVersion("ucar-debug-module-1.2.0.1.jar")
	assert.True(t, ok)
	assert.Equal(t, "ucar-debug-module", name)
	assert.Equal(t, "1.2.0.1", version)
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

	err = Dos2Unix(fileName)
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
