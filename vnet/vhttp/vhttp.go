// Copyright 2019 The vogo Authors. All rights reserved.

package vhttp

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/vogo/vogo/vio/vioutil"

	"github.com/vogo/logger"
)

const (
	XForwardedFor = "X-Forwarded-For"
	XRealIP       = "X-Real-IP"

	// 2 min
	DefaultTimeout = 120 * time.Second
)

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filePath, rawURL string, timeout time.Duration) error {
	u, err := url.Parse(rawURL)
	if err != nil {
		return err
	}

	// Get the data
	resp, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
	case 404:
		return fmt.Errorf("file not found: %s", rawURL)
	default:
		buf := make([]byte, 1024)
		result := ""

		if n, readErr := resp.Body.Read(buf); n > 0 && readErr == nil {
			result = string(buf[:n])
		}

		return fmt.Errorf("download failed, status code: %d, result: %s", resp.StatusCode, result)
	}

	logger.Infof("download %s to %s", rawURL, filePath)

	if timeout == 0 {
		timeout = DefaultTimeout
	}

	return vioutil.WriteDataToFile(filePath, resp.Body, timeout)
}

// RemoteIP http remote ip address
func RemoteIP(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get(XRealIP); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get(XForwardedFor); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}

// IP2long convert IPv4 to uint32
func IP2long(ipstr string) uint32 {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0
	}

	ip = ip.To4()

	return binary.BigEndian.Uint32(ip)
}

// Get url response
func Get(rawURL string) ([]byte, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(u.String())

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("response status: %d", resp.StatusCode)
	}

	return body, err
}

// IsConnectionError is http connection error
func IsConnectionError(err error) bool {
	switch err {
	case http.ErrServerClosed, net.ErrWriteToConnected:
		return true
	default:
		if _, ok := err.(*net.OpError); ok {
			return true
		}

		if _, ok := err.(*url.Error); ok {
			return true
		}

		return false
	}
}
