//author: wongoo
//date: 20190612

package vioutil

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	vbytes "github.com/vogo/vogo/bytes"

	"github.com/vogo/logger"
)

// ReadFile read file to string
func ReadFile(filePath string) string {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		logger.Info(err.Error())
	}

	return string(bytes)
}

func IsDirEmpty(dirPath string) bool {
	files, _ := ioutil.ReadDir(dirPath)

	for _, fi := range files {
		if fi.IsDir() {
			jars, _ := ioutil.ReadDir(dirPath + string(os.PathSeparator) + fi.Name())
			if len(jars) > 0 {
				return false
			}
		}
	}

	return true
}
func MkdirByRecursive(dirPath string) {
	logger.Infof("mkdir %s", dirPath)
	err := os.MkdirAll(dirPath, os.ModePerm)

	if err != nil {
		logger.Infof("mkdir business dir fail, path: %s, err: %s", dirPath, err.Error())
	}
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}

	return s.IsDir()
}

func IsFile(path string) bool {
	return !IsDir(path)
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		logger.Infof("open src file fail, err: " + err.Error())
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		logger.Infof("open dst file fail, err: " + err.Error())
		return
	}
	defer dst.Close()
	logger.Infof("copy file success, dst: %s, src: %s", dst.Name(), src.Name())

	return io.Copy(dst, src)
}

// ExistFile check file exists
func ExistFile(file string) bool {
	if s, err := os.Stat(file); err == nil {
		return !s.IsDir()
	} else if os.IsNotExist(err) {
		return false
	}

	return false
}

// ExistDir check dir exists
func ExistDir(file string) bool {
	s, err := os.Stat(file)
	if err != nil {
		return false
	}

	return s.IsDir()
}

// AppendFile append data to file
func AppendFile(filePath string, data []byte, perm os.FileMode) error {
	if !ExistFile(filePath) {
		return ioutil.WriteFile(filePath, data, perm)
	}

	// the following append file data
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, perm)
	if err != nil {
		return err
	}

	_, err = f.Write(data)

	if err1 := f.Close(); err == nil {
		err = err1
	}

	return err
}

// LinkLatest link the latest
func LinkLatest(dirPath, targetDirPath, prefix, suffix string) error {
	files, err := ListFileNames(dirPath, prefix, suffix)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return fmt.Errorf("can't find file with prefix [%s], suffix [%s] in directory %s", prefix, suffix, dirPath)
	}

	sort.Strings(files)
	sourceFilePath := filepath.Join(dirPath, files[len(files)-1])
	sourceFile, err := os.Stat(sourceFilePath)

	if err != nil {
		return err
	}

	targetFileName := "latest"
	if !sourceFile.IsDir() {
		targetFileName = prefix + "-latest" + suffix
	}

	targetFilePath := filepath.Join(targetDirPath, targetFileName)

	logger.Infof("create symbolic link %s to %s", sourceFilePath, targetFilePath)
	// remove the exists link file before create
	if _, err := os.Lstat(targetFilePath); err == nil {
		_ = os.Remove(targetFilePath)
	}

	return os.Symlink(sourceFilePath, targetFilePath)
}

// ListFileNames list file names which match the given prefix and suffix
func ListFileNames(dirPath, prefix, suffix string) ([]string, error) {
	f, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}

	fileInfos, err := f.Readdir(-1)
	_ = f.Close()

	if err != nil {
		return nil, err
	}

	var files []string

	for _, info := range fileInfos {
		name := info.Name()

		if prefix != "" && !strings.HasPrefix(name, prefix) {
			continue
		}

		if suffix != "" && !strings.HasSuffix(name, suffix) {
			continue
		}

		files = append(files, name)
	}

	return files, nil
}

// parse package name version
func ParsePackageNameVersion(fileName string) (name, version string, ok bool) {
	index := strings.LastIndex(fileName, "-")
	if index <= 0 {
		logger.Infof("invalid package name %s", fileName)
		return "", "", false
	}

	name = fileName[:index]
	version = fileName[index+1 : len(fileName)-4]
	ok = true

	return
}

// Move file
func Move(from, to string) error {
	os.Remove(to)
	return os.Rename(from, to)
}

// Dos2Unix change file format to unix
func Dos2Unix(fileName string) error {
	var err error

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	tmpFileName := fileName + ".tmp"
	wFile, err := os.Create(tmpFileName)

	if err != nil {
		return err
	}
	defer wFile.Close()
	w := bufio.NewWriter(wFile)

	if err := vbytes.CopyFilterBytes(file, wFile, []byte{'\r'}); err != nil {
		return err
	}

	if err := w.Flush(); err != nil {
		return err
	}

	return Move(tmpFileName, fileName)
}

// WriteDataToFile read data and write to file in a give limit time
// it will write to a temp file first, and then rename to the target file
func WriteDataToFile(filePath string, data io.Reader, timeout time.Duration) error {
	tempPath := filePath + ".tmp"

	// Create temp file
	out, err := os.Create(tempPath)
	if err != nil {
		logger.Infof("can't create file: %v", err)
		return err
	}

	// Write the body to file in a limit time
	err = vbytes.TimeoutCopy(out, data, timeout)

	// close file
	_ = out.Close()

	// delete temp file if download error occurs
	if err != nil {
		_ = os.Remove(tempPath)
		return err
	}

	// remove exists file first
	_ = os.Remove(filePath)

	// rename
	return os.Rename(tempPath, filePath)
}
