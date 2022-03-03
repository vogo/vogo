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

package vhttp

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
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

const (
	HeaderContentType = "content-type"
	ContentTypeJson   = "application/json"
)

var jsonContentTypeHeader = map[string]string{
	HeaderContentType: ContentTypeJson,
}

var ErrHTTPStatusNotOK = errors.New("http status not ok")
var ErrHTTPFail = errors.New("http failed")

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
		return fmt.Errorf("%w: download url %s", ErrHTTPFail, rawURL)
	default:
		buf := make([]byte, 1024)
		result := ""

		if n, readErr := resp.Body.Read(buf); n > 0 && readErr == nil {
			result = string(buf[:n])
		}

		return fmt.Errorf("%w: [%d]%s", ErrHTTPFail, resp.StatusCode, result)
	}

	logger.Infof("download %s to %s", rawURL, filePath)

	if timeout == 0 {
		timeout = DefaultTimeout
	}

	return vioutil.WriteDataToFile(filePath, resp.Body, timeout)
}

// RemoteIP http remote ip address.
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

// IP2long convert IPv4 to uint32.
func IP2long(ipstr string) uint32 {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0
	}

	ip = ip.To4()

	return binary.BigEndian.Uint32(ip)
}

// Get url response.
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
		err = fmt.Errorf("%w: response status %d", ErrHTTPFail, resp.StatusCode)
	}

	return body, err
}

// IsConnectionError is http connection error.
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

func ParseGet(url string, headers map[string]string, obj interface{}) error {
	return parseJsonResponse(http.MethodGet, url, headers, nil, obj)
}

func ParsePost(url string, headers map[string]string, body interface{}, obj interface{}) error {
	var data io.Reader

	if body != nil {
		switch raw := body.(type) {
		case []byte:
			data = bytes.NewReader(raw)
		case string:
			data = strings.NewReader(raw)
		default:
			bytesData, jsonErr := json.Marshal(body)
			if jsonErr != nil {
				return jsonErr
			}
			data = bytes.NewReader(bytesData)
		}
	}

	if headers == nil {
		headers = jsonContentTypeHeader
	} else {
		headers[HeaderContentType] = ContentTypeJson
	}

	return parseJsonResponse(http.MethodPost, url, headers, data, obj)
}

func parseJsonResponse(method, url string, headers map[string]string, body io.Reader, obj interface{}) error {
	req, _ := http.NewRequest(method, url, body)
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	resp, err := (&http.Client{}).Do(req)
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%v, status: %d, body: %s", ErrHTTPStatusNotOK, resp.StatusCode, b)
	}

	if err = json.Unmarshal(b, obj); err != nil {
		return err
	}

	return nil
}
