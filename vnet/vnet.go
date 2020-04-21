// Copyright 2019-2020 The vogo Authors. All rights reserved.

package vnet

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/vogo/vogo/vos"
)

// LocalPortExist check whether local port exist
func LocalPortExist(port int) bool {
	if port < 1 {
		return false
	}

	conn, err := net.DialTimeout("tcp", fmt.Sprintf(":%d", port), time.Second)

	if err != nil {
		return false
	}

	if conn != nil {
		_ = conn.Close()
		return true
	}

	return false
}

func LocalIP() (string, error) {
	faces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	var (
		ipv4 string
		ok   bool
	)

	for _, face := range faces {
		if face.Flags&net.FlagUp == 0 {
			continue // interface down
		}

		if face.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}

		addrs, err := face.Addrs()
		if err != nil {
			return "", err
		}

		if ipv4, ok = getExternalIPv4(addrs); ok {
			faceName := strings.ToLower(face.Name)
			if !strings.Contains(faceName, "docker") && !strings.Contains(faceName, "vmware") {
				return ipv4, nil
			}
		}
	}

	if ipv4 != "" {
		return ipv4, nil
	}

	return "", errors.New("network not connected")
}

func getExternalIPv4(addrs []net.Addr) (string, bool) {
	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		}

		if ip == nil || ip.IsLoopback() {
			continue
		}

		ip = ip.To4()

		if ip == nil {
			continue // not an ipv4 address
		}

		return ip.String(), true
	}

	return "", false
}

func ConnectionCount() (established, listen, closeWait int, err error) {
	result, err := vos.Shell(
		`netstat -ant | awk '/ESTABLISHED|LISTEN|CLOSE_WAIT/ {count[$6]++} END { for(s in count) {  printf("%s:%d\n", s, count[s]); }}'`)
	if err != nil {
		return
	}

	counters := strings.Split(string(result), "\n")

	for _, counter := range counters {
		values := strings.Split(counter, ":")
		if len(values) == 2 {
			switch values[0] {
			case "ESTABLISHED":
				established, _ = strconv.Atoi(values[1])
			case "LISTEN":
				listen, _ = strconv.Atoi(values[1])
			case "CLOSE_WAIT":
				closeWait, _ = strconv.Atoi(values[1])
			}
		}
	}

	return
}
