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

package vnet

import (
	"bytes"
	"log"
	"os"
)

func GetRouteInterfaces() (routeFaces []string) {
	data, err := os.ReadFile("/proc/net/route")
	if err != nil {
		log.Printf("failed to read /proc/net/route, error: %+v", err)
		return
	}
	return parseRouteInterfaces(data)
}

func parseRouteInterfaces(data []byte) (routeFaces []string) {
	lines := bytes.Split(data, []byte{'\n'})
	names := make(map[string]struct{})
	for _, line := range lines[1:] {
		index := bytes.IndexByte(line, '\t')
		if index == -1 {
			index = bytes.IndexByte(line, ' ')
		}
		if index > 0 {
			names[string(line[:index])] = struct{}{}
		}
	}
	for name := range names {
		routeFaces = append(routeFaces, name)
	}
	return
}
