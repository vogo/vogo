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

// Recommend log format style: `"log msg | key1: %s | key2: %d | key3: %s | err: %v"` (pipe-separated key-value pairs, the name of key is in snake case)

package vlog

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

const (
	TagTrace = "TRAC"
	TagDebug = "DEBG"
	TagInfo  = "INFO"
	TagWarn  = "WARN"
	TagError = "ERRO"
	TagFatal = "FATL"
	TagPanic = "PNIC"
	TagPrint = "PRNT"

	LevelTrace = 5
	LevelDebug = 4
	LevelInfo  = 3
	LevelWarn  = 2
	LevelError = 1
	LevelFatal = 0
)

const (
	Lnone     = 0             // none file
	Lfile     = 1             // d.go:23
	Lfunc     = 1 << 1        // foo
	LfileFunc = Lfunc | Lfile // d.go:foo:23
)

var (
	Level              = LevelInfo
	output   io.Writer = os.Stdout
	flag     int
	instance []byte
)

// SetLevel set logger Level
// the Level variable is exported and can be set directly.
func SetLevel(l int) {
	Level = l
}

// SetOutput set logger output writer
func SetOutput(w io.Writer) {
	output = w
}

// SetFlags set logger flags
func SetFlags(f int) {
	flag = f
}

// SetInstance set logger instance
func SetInstance(s []byte) {
	instance = s
}

// Writer return the logger writer
func Writer() io.Writer {
	return output
}

func Trace(a ...any) {
	if Level < LevelTrace {
		return
	}
	WriteLog(TagTrace, fmt.Sprint(a...))
}

func Debug(a ...any) {
	if Level < LevelDebug {
		return
	}
	WriteLog(TagDebug, fmt.Sprint(a...))
}

func Info(a ...any) {
	if Level < LevelInfo {
		return
	}
	WriteLog(TagInfo, fmt.Sprint(a...))
}

func Warn(a ...any) {
	if Level < LevelWarn {
		return
	}
	WriteLog(TagWarn, fmt.Sprint(a...))
}

func Error(a ...any) {
	if Level < LevelError {
		return
	}
	WriteLog(TagError, fmt.Sprint(a...))
}

func Tracef(format string, a ...any) {
	if Level < LevelTrace {
		return
	}
	WriteLog(TagTrace, fmt.Sprintf(format, a...))
}

func Debugf(format string, a ...any) {
	if Level < LevelDebug {
		return
	}
	WriteLog(TagDebug, fmt.Sprintf(format, a...))
}

func Infof(format string, a ...any) {
	if Level < LevelInfo {
		return
	}
	WriteLog(TagInfo, fmt.Sprintf(format, a...))
}

func Warnf(format string, a ...any) {
	if Level < LevelWarn {
		return
	}
	WriteLog(TagWarn, fmt.Sprintf(format, a...))
}

func Errorf(format string, a ...any) {
	if Level < LevelError {
		return
	}
	WriteLog(TagError, fmt.Sprintf(format, a...))
}

func Fatal(a ...any) {
	WriteLog(TagFatal, fmt.Sprint(a...))
	os.Exit(1)
}

func Fatalf(format string, a ...any) {
	WriteLog(TagFatal, fmt.Sprintf(format, a...))
	os.Exit(1)
}

func Fatalln(a ...any) {
	WriteLog(TagFatal, fmt.Sprint(a...))
	os.Exit(1)
}

func Print(a ...any) {
	WriteLog(TagPrint, fmt.Sprint(a...))
}

func Printf(format string, a ...any) {
	WriteLog(TagPrint, fmt.Sprintf(format, a...))
}

func Println(format string, a ...any) {
	WriteLog(TagPrint, fmt.Sprintf(format, a...))
}

func Panic(a ...any) {
	s := fmt.Sprint(a...)
	WriteLog(TagPanic, s)
	panic(s)
}

func Panicf(format string, a ...any) {
	s := fmt.Sprintf(format, a...)
	WriteLog(TagPanic, s)
	panic(s)
}

func Panicln(a ...any) {
	s := fmt.Sprint(a...)
	WriteLog(TagPanic, s)
	panic(s)
}

var bytesPool = sync.Pool{New: func() any {
	b := make([]byte, 0, 1024)
	return &b
}}

// WriteLog write log data
func WriteLog(tag, s string) {
	t := time.Now()

	buf := bytesPool.Get().(*[]byte)

	year, month, day := t.Date()
	appendNumber(buf, year, 4)
	*buf = append(*buf, '/')
	appendNumber(buf, int(month), 2)
	*buf = append(*buf, '/')
	appendNumber(buf, day, 2)
	*buf = append(*buf, ' ')

	hour, min, sec := t.Clock()
	appendNumber(buf, hour, 2)
	*buf = append(*buf, ':')
	appendNumber(buf, min, 2)
	*buf = append(*buf, ':')
	appendNumber(buf, sec, 2)
	*buf = append(*buf, '.')
	appendNumber(buf, t.Nanosecond()/1e6, 3)

	if len(instance) > 0 {
		*buf = append(*buf, ' ', '[')
		*buf = append(*buf, instance...)
		*buf = append(*buf, ']')
	}

	*buf = append(*buf, ' ')
	*buf = append(*buf, tag...)

	if flag&Lfile != 0 {
		*buf = append(*buf, ' ', '[')
		pc, fileName, line, callerOk := runtime.Caller(2)
		if callerOk {
			for i := len(fileName) - 1; i > 0; i-- {
				if fileName[i] == '/' {
					fileName = fileName[i+1:]
					break
				}
			}
			*buf = append(*buf, fileName...)
			if flag&Lfunc != 0 {
				funcName := runtime.FuncForPC(pc).Name() // main.(*MyStruct).foo

				for i := len(funcName) - 1; i > 0; i-- {
					if funcName[i] == '.' {
						funcName = funcName[i+1:]
						break
					}
				}

				*buf = append(*buf, ':')
				*buf = append(*buf, funcName...)
			}
			*buf = append(*buf, ':')
			appendNumber(buf, line, -1)
		} else {
			*buf = append(*buf, '?')
		}

		*buf = append(*buf, ']')
	}

	*buf = append(*buf, ' ')
	*buf = append(*buf, s...)
	if s == "" || s[len(s)-1] != '\n' {
		*buf = append(*buf, '\n')
	}

	_, _ = output.Write(*buf)

	*buf = (*buf)[:0]
	bytesPool.Put(buf)
}

// Cheap integer to fixed-width decimal ASCII. Give a negative width to avoid zero-padding.
func appendNumber(buf *[]byte, i, wid int) {
	// Assemble decimal in reverse order.
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	// i < 10
	b[bp] = byte('0' + i)
	*buf = append(*buf, b[bp:]...)
}
