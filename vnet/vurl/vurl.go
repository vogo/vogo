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

package vurl

import (
	"net"
	"net/http"
	"net/url"

	"github.com/vogo/logger"
)

// IsURLNetError whether net error for url request.
func IsURLNetError(err error) bool {
	// nolint:errorlint // ignore this
	urlErr, ok := err.(*url.Error)
	if !ok {
		return false
	}

	// nolint:errorlint // ignore this
	_, ok = urlErr.Err.(*net.OpError)

	return ok
}

// nolint:exhaustivestruct // ignore this.
var nonRedirectHTTPClient = &http.Client{
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
}

// RedirectURL redirect url.
func RedirectURL(urlAddr string) string {
	for {
		resp, err := nonRedirectHTTPClient.Get(urlAddr)
		if err != nil {
			logger.Debugf("redirect url error: %v", err)

			return urlAddr
		}

		_ = resp.Body.Close()

		if resp.StatusCode != http.StatusFound {
			return urlAddr
		}

		urlAddr = resp.Header.Get("Location")
	}
}
