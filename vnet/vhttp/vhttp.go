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
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/vogo/vogo/vio/vioutil"
)

const (
	XForwardedFor = "X-Forwarded-For"
	XRealIP       = "X-Real-IP"

	// DefaultDownloadTimeout 2 min.
	DefaultDownloadTimeout = 120 * time.Second
)

const (
	HeaderContentType = "content-type"
	ContentTypeJSON   = "application/json"

	DefaultMaxIdleConns        = 32
	DefaultMaxIdleConnsPerHost = 8
	DefaultMaxConnsPerHost     = 64
	DefaultIdleConnTimeout     = time.Second * 8

	DefaultRequestTimeout = time.Second * 16
)

var jsonContentTypeHeader = map[string]string{
	HeaderContentType: ContentTypeJSON,
}

var (
	ErrHTTPStatusNotOK = errors.New("http status not ok")
	ErrHTTPFail        = errors.New("http failed")
)

//nolint:exhaustivestruct // ignore this
var defaultHTTPClient = &http.Client{
	Transport: &http.Transport{
		MaxIdleConns:        DefaultMaxIdleConns,
		MaxIdleConnsPerHost: DefaultMaxIdleConnsPerHost,
		MaxConnsPerHost:     DefaultMaxConnsPerHost,
		IdleConnTimeout:     DefaultIdleConnTimeout,
	},
	Timeout: DefaultRequestTimeout,
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filePath, rawURL string, timeout time.Duration) error {
	u, err := url.Parse(rawURL)
	if err != nil {
		return err
	}

	resp, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusNotFound:
		return fmt.Errorf("%w: download url %s", ErrHTTPFail, rawURL)
	default:
		buf := make([]byte, 1024)
		result := ""

		if n, readErr := resp.Body.Read(buf); n > 0 && readErr == nil {
			result = string(buf[:n])
		}

		return fmt.Errorf("%w: [%d]%s", ErrHTTPFail, resp.StatusCode, result)
	}

	if timeout == 0 {
		timeout = DefaultDownloadTimeout
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("%w: response status %d", ErrHTTPFail, resp.StatusCode)
	}

	return body, err
}

// IsConnectionError is http connection error.
func IsConnectionError(err error) bool {
	if errors.Is(err, http.ErrServerClosed) ||
		errors.Is(err, net.ErrWriteToConnected) {
		return true
	}

	//nolint:errorlint // ignore this
	if _, ok := err.(*net.OpError); ok {
		return true
	}

	//nolint:errorlint // ignore this
	if _, ok := err.(*url.Error); ok {
		return true
	}

	return false
}

func ParseGet(urlAddr string, headers map[string]string, obj any) error {
	return parseJSONResponse(http.MethodGet, urlAddr, headers, nil, obj)
}

func ParsePost(urlAddr string, headers map[string]string, body, obj any) error {
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
		headers[HeaderContentType] = ContentTypeJSON
	}

	return parseJSONResponse(http.MethodPost, urlAddr, headers, data, obj)
}

func parseJSONResponse(method, urlAddr string, headers map[string]string, body io.Reader, obj any) error {
	req, err := http.NewRequest(method, urlAddr, body)
	if err != nil {
		return err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := defaultHTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%w, status: %d, body: %s", ErrHTTPStatusNotOK, resp.StatusCode, b)
	}

	if jsonErr := json.Unmarshal(b, obj); jsonErr != nil {
		return jsonErr
	}

	return nil
}
