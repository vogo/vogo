//+build !linux

package vioutil

import (
	"errors"
	"os"
)

var (
	errUnsupported = errors.New("unsupported")
)

func LockFile(file *os.File) error {
	return nil
}

func UnLockFile(file *os.File) error {
	return nil
}

// Touch create file if not exists
func Touch(fileName, userName string) error {
	return errUnsupported
}
