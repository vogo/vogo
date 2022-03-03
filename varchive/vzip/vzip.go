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

package vzip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const DefaultZipLimitSize = 8 * 1024 * 1024

// Unzip will decompress a zip archive, moving all files and folders
// within the zip file (parameter 1) to an output directory (parameter 2).
func Unzip(src, destDir string) error {
	return LimitUnzip(src, destDir, DefaultZipLimitSize)
}

// LimitUnzip decompress a zip archive, while limit the size of a single containing file.
func LimitUnzip(src, destDir string, limitSize int64) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	// Make File
	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		return err
	}

	for _, zipFile := range r.File {
		if err := unzipFile(destDir, zipFile, limitSize); err != nil {
			return err
		}
	}

	return nil
}

func unzipFile(destDir string, f *zip.File, limitSize int64) error {
	// Check for ZipSlip
	fileName := strings.ReplaceAll(f.Name, "..", "")

	if fileName != f.Name {
		return fmt.Errorf("%w: invalid zip file %s", os.ErrInvalid, f.Name)
	}

	targetPath := filepath.Join(destDir, fileName)

	if f.FileInfo().IsDir() {
		return os.MkdirAll(targetPath, f.Mode())
	}

	if err := os.MkdirAll(filepath.Dir(targetPath), os.ModePerm); err != nil {
		return err
	}

	outFile, err := os.OpenFile(targetPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}

	rc, err := f.Open()
	if err != nil {
		return err
	}

	_, err = io.CopyN(outFile, rc, limitSize)

	outFile.Close()
	_ = rc.Close()

	// ignore EOF for copyN
	if err == io.EOF {
		return nil
	}

	return err
}

// ZipDir compresses a directory into a zip archive file.
func ZipDir(zipPath, dir string) error {
	newZipFile, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	baseDirLen := len(dir)
	if dir[len(dir)-1] != '/' {
		baseDirLen = len(filepath.Dir(dir))
	}

	return AddDirToZip(zipWriter, baseDirLen, dir)
}

// AddDirToZip add all files under the target directory into a zip file.
func AddDirToZip(writer *zip.Writer, baseDirLen int, dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if dir == path || info.IsDir() {
			return nil
		}

		return AddFileToZip(writer, path, path[baseDirLen:])
	})
}

// AddFileToZip add a single file into a zip file.
func AddFileToZip(zipWriter *zip.Writer, filePath, pathInZip string) error {
	fileToZip, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	// Using FileInfoHeader() above only uses the basename of the file. If we want
	// to preserve the folder structure we can overwrite this with the full path.
	header.Name = pathInZip

	header.Method = chooseCompressMethod(filepath.Ext(filePath))

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, fileToZip)

	return err
}

func chooseCompressMethod(ext string) uint16 {
	switch ext {
	case ".jar", ".z", ".gz", ".tar", ".zip":
		return zip.Store
	default:
		return zip.Deflate
	}
}
