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

package vjava

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/vogo/vogo/vos"
)

var ErrJavaHomeNotFound = errors.New("can't find java home")

func GetJavaHome(pid int) (string, error) {
	if !vos.PidExist(pid) {
		return "", fmt.Errorf("%w: pid %d", os.ErrNotExist, pid)
	}

	fullCommand := fmt.Sprintf(`lsof -p %d \
|grep "/bin/java" \
|awk '{print $9}' \
|xargs ls -l \
|awk '{if($1~/^l/){print $11}else{print $9}}' \
|xargs ls -l \
|awk '{if($1~/^l/){print $11}else{print $9}}'`, pid)

	result, err := vos.SingleCommandResult(fullCommand)
	if err != nil {
		return "", err
	}

	if result == "" {
		return "", ErrJavaHomeNotFound
	}

	const javaBinSuffix = "/bin/java"
	if strings.HasSuffix(result, javaBinSuffix) {
		return result[:len(result)-len(javaBinSuffix)], nil
	}

	return "", fmt.Errorf("%w: path %s", ErrJavaHomeNotFound, result)
}

// ReadAllJavaProcessEnv read all java process env.
func ReadAllJavaProcessEnv() []map[string]string {
	var processes []map[string]string

	result, err := vos.Shell("ps -o pid,cmd -e |grep java |grep -v grep")
	if err != nil {
		return nil
	}

	lines := bytes.Split(result, []byte{'\n'})

	if len(lines) == 0 {
		return nil
	}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		for line[0] == ' ' {
			line = line[1:]
		}

		index := bytes.Index(line, []byte{' '})
		if index <= 0 {
			continue
		}

		pid := line[:index]
		javaProc := line[index+1:]

		if len(javaProc) > 0 {
			for javaProc[0] == ' ' {
				javaProc = javaProc[1:]
			}

			procEnv := vos.ReadProcEnv(pid)
			procEnv["java_process"] = string(javaProc)

			processes = append(processes, procEnv)
		}
	}

	return processes
}
