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

package vuser

import "fmt"

const (
	EnvValueSplit = ":"
)

func GetUserEnvProfiles() []string {
	userName := GetCurrentUserName()

	files := []string{
		"/etc/bashrc",
		"/etc/profile",
	}

	if userName == "root" {
		return append(files, "/root/.bashrc", "/root/.bash_profile")
	}

	return append(files, fmt.Sprintf("/home/%s/.bashrc", userName),
		fmt.Sprintf("/home/%s/.bash_profile", userName))
}
