//go:build linux
// +build linux

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

package vioutil

import (
	"os"
	"os/user"
	"strconv"
	"syscall"

	"github.com/vogo/vogo/vos"

	"github.com/vogo/logger"
)

func LockFile(file *os.File) error {
	return syscall.Flock(int(file.Fd()), syscall.LOCK_EX)
}

func UnLockFile(file *os.File) error {
	return syscall.Flock(int(file.Fd()), syscall.LOCK_UN)
}

// Touch create file if not exists
func Touch(fileName, userName string) error {
	if !ExistFile(fileName) {
		f, err := os.Create(fileName)
		if err != nil {
			logger.Infof("failed to create file %s, error: %v", fileName, err)
			return err
		}

		defer f.Close()

		if userName != "" && userName != vos.CurrUserHome() {
			u, err := user.Lookup(userName)
			if err != nil {
				logger.Infof("failed to change file owner %s, error: %v", fileName, err)
				return err
			}
			uid, _ := strconv.Atoi(u.Uid)
			gid, _ := strconv.Atoi(u.Gid)
			return os.Chown(fileName, uid, gid)
		}
	}
	return nil
}
