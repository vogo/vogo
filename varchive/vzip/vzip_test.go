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

package vzip_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vogo/vogo/varchive/vzip"
	"github.com/vogo/vogo/vio/vioutil"
)

func TestZipDir(t *testing.T) {
	t.Parallel()

	workDir := filepath.Join(os.TempDir(), "test_zip_dir")
	defer os.RemoveAll(workDir)

	assert.NoError(t, os.MkdirAll(filepath.Join(workDir, "a", "b"), os.ModePerm))

	assert.NoError(t, os.WriteFile(filepath.Join(workDir, "a", "a1.txt"), []byte("aaa1"), 0o600))
	assert.NoError(t, os.WriteFile(filepath.Join(workDir, "a", "a2.txt"), []byte("aaa2"), 0o600))
	assert.NoError(t, os.WriteFile(filepath.Join(workDir, "a", "b", "b1.txt"), []byte("bbb1"), 0o600))
	assert.NoError(t, os.WriteFile(filepath.Join(workDir, "a", "b", "b2.txt"), []byte("bbb2"), 0o600))

	zipPath := filepath.Join(workDir, "test.zip")
	zipDir := filepath.Join(workDir, "a")
	outputDir := filepath.Join(workDir, "output")

	t.Logf("zip path: %s", zipPath)
	t.Logf("zip dir: %s", zipDir)

	assert.NoError(t, vzip.ZipDir(zipPath, zipDir))
	assert.NoError(t, vzip.Unzip(zipPath, outputDir))
	assert.True(t, vioutil.ExistFile(filepath.Join(outputDir, "a", "a1.txt")))
	assert.True(t, vioutil.ExistFile(filepath.Join(outputDir, "a", "a2.txt")))
	assert.True(t, vioutil.ExistFile(filepath.Join(outputDir, "a", "b", "b1.txt")))
	assert.True(t, vioutil.ExistFile(filepath.Join(outputDir, "a", "b", "b2.txt")))
	assert.Equal(t, "aaa1", vioutil.ReadFile(filepath.Join(outputDir, "a", "a1.txt")))

	assert.NoError(t, os.RemoveAll(zipPath))
	assert.NoError(t, os.RemoveAll(outputDir))

	zipDir = filepath.Join(workDir, "a") + "/"

	t.Logf("zip path: %s", zipPath)
	t.Logf("zip dir: %s", zipDir)

	assert.NoError(t, vzip.ZipDir(zipPath, zipDir))
	assert.NoError(t, vzip.Unzip(zipPath, outputDir))
	assert.True(t, vioutil.ExistFile(filepath.Join(outputDir, "a1.txt")))
	assert.True(t, vioutil.ExistFile(filepath.Join(outputDir, "a2.txt")))
	assert.True(t, vioutil.ExistFile(filepath.Join(outputDir, "b", "b1.txt")))
	assert.True(t, vioutil.ExistFile(filepath.Join(outputDir, "b", "b2.txt")))
	assert.Equal(t, "bbb1", vioutil.ReadFile(filepath.Join(outputDir, "b", "b1.txt")))
}
