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

package vos

import (
	"context"
	"os/exec"
	"strings"

	"github.com/vogo/logger"
)

func ExecShell(fullCommand string) ([]byte, error) {
	logger.Debugf("exec: %s", fullCommand)
	cmd := exec.Command("/bin/sh", "-c", fullCommand)

	return cmd.CombinedOutput()
}

func ExecContext(ctx context.Context, fullCommand string) ([]byte, error) {
	logger.Debugf("exec: %s", fullCommand)
	cmd := exec.CommandContext(ctx, "/bin/sh", "-c", fullCommand)

	return cmd.CombinedOutput()
}

// Shell execute shell without log.
func Shell(fullCommand string) ([]byte, error) {
	cmd := exec.Command("/bin/sh", "-c", fullCommand)

	return cmd.CombinedOutput()
}

// SingleCommandResult exec command and return a single line result.
func SingleCommandResult(fullCommand string) (string, error) {
	output, err := ExecShell(fullCommand)
	if err != nil {
		return "", err
	}

	result := string(output)
	result = strings.Replace(result, "\r", "", -1)
	result = strings.Replace(result, "\n", "", -1)

	return result, nil
}
