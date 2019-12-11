// Copyright 2019 The vogo Authors. All rights reserved.

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
