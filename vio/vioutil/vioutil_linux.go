//+build linux

// Copyright 2019 The vogo Authors. All rights reserved.

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
