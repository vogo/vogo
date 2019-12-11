//+build linux

// Copyright 2019 The vogo Authors. All rights reserved.

package vnet

import (
	"bytes"
	"io/ioutil"

	"github.com/vogo/logger"
)

func GetRouteInterfaces() (routeFaces []string) {
	data, err := ioutil.ReadFile("/proc/net/route")
	if err != nil {
		logger.Infof("failed to read /proc/net/route, error: %+v", err)
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
