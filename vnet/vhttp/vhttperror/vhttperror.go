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

package vhttperror

import "net/http"

const (
	CodeOK                 = 0
	CodeUnknownErr         = 10
	CodeUnauthenticatedErr = 20
	CodeUnauthorizedErr    = 21
	CodeForbiddenErr       = 22
	CodeBadRequestErr      = 100
	CodeNotFoundErr        = 101
	CodeArgRequiredErr     = 102
	CodeValueInvalidErr    = 103
)

var (
	ErrBadRequest      = NewStatusCodeError(http.StatusBadRequest, CodeBadRequestErr, "forbidden")
	ErrNotFound        = NewStatusCodeError(http.StatusNotFound, CodeNotFoundErr, "not found")
	ErrArgRequired     = NewStatusCodeError(http.StatusBadRequest, CodeArgRequiredErr, "arg required")
	ErrValueInvalid    = NewStatusCodeError(http.StatusBadRequest, CodeValueInvalidErr, "value invalid")
	ErrUnauthenticated = NewStatusCodeError(http.StatusUnauthorized, CodeUnauthenticatedErr, "unauthenticated")
	ErrUnauthorized    = NewStatusCodeError(http.StatusUnauthorized, CodeUnauthorizedErr, "unauthorized")
	ErrForbidden       = NewStatusCodeError(http.StatusForbidden, CodeForbiddenErr, "forbidden")
)

type Coder interface {
	Code() int
}

type StatusState interface {
	Status() int
}

type CodeError interface {
	error
	Coder
}
type StatusCodeError interface {
	CodeError
	StatusState
}

func NewCodeError(code int, err string) CodeError {
	return &codeError{c: code, m: err}
}

type codeError struct {
	c int
	m string
}

func (e *codeError) Code() int {
	return e.c
}

func (e *codeError) Error() string {
	return e.m
}

func NewStatusCodeError(status, code int, err string) CodeError {
	return &statusCodeError{s: status, c: code, m: err}
}

type statusCodeError struct {
	s int
	c int
	m string
}

func (e *statusCodeError) Code() int {
	return e.c
}

func (e *statusCodeError) Status() int {
	return e.s
}

func (e *statusCodeError) Error() string {
	return e.m
}
