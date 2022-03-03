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
	"testing"

	"github.com/vogo/vogo/vstrings"

	"github.com/stretchr/testify/assert"
)

func TestGetRouteInterfaces(t *testing.T) {
	data := `Iface	Destination	Gateway 	Flags	RefCnt	Use	Metric	Mask		MTU	Window	IRTT
eth0	00000000	FEFF680A	0003	0	0	100	00000000	0	0	0
eth0	0000680A	00000000	0001	0	0	100	0000FFFF	0	0	0
eth0	0000FEA9	00000000	0001	0	0	1002	0000FFFF	0	0	0
eth0	FEA9FEA9	CB66680A	0007	0	0	100	FFFFFFFF	0	0	0
docker0	000011AC	00000000	0001	0	0	0	0000FFFF	0	0	0`

	faces := parseRouteInterfaces([]byte(data))

	assert.Equal(t, 2, len(faces))
	assert.True(t, vstrings.ContainsIn(faces, "eth0"))
	assert.True(t, vstrings.ContainsIn(faces, "docker0"))
}
