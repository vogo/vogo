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

package vhttpresp

import (
	"encoding/json"
	"github.com/vogo/vogo/vnet/vhttp/vhttperror"
	"net/http"

	"github.com/vogo/logger"
)

type ResponseBody struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:"data,omitempty"`
}

func Data(w http.ResponseWriter, req *http.Request, code int, data any) {
	Write(w, req, code, "", data)
}

func CodeData(w http.ResponseWriter, req *http.Request, code int, msg string, data any) {
	Write(w, req, code, msg, data)
}

func OK(w http.ResponseWriter, req *http.Request) {
	Write(w, req, vhttperror.CodeOK, "", "ok")
}

func Success(w http.ResponseWriter, req *http.Request, data any) {
	Write(w, req, vhttperror.CodeOK, "", data)
}

func CodeError(w http.ResponseWriter, req *http.Request, code int, err error) {
	CodeMsg(w, req, code, err.Error())
}

func Error(w http.ResponseWriter, req *http.Request, err error) {
	if c, ok := err.(vhttperror.StatusState); ok {
		w.WriteHeader(c.Status())
	}

	code := vhttperror.CodeUnknownErr

	if c, ok := err.(vhttperror.Coder); ok {
		code = c.Code()
	}

	CodeMsg(w, req, code, err.Error())
}

func BadMsg(w http.ResponseWriter, req *http.Request, msg string) {
	CodeMsg(w, req, vhttperror.CodeBadRequestErr, msg)
}

func BadError(w http.ResponseWriter, req *http.Request, err error) {
	BadMsg(w, req, err.Error())
}

func CodeMsg(w http.ResponseWriter, req *http.Request, code int, msg string) {
	Write(w, req, code, msg, nil)
}

func Write(w http.ResponseWriter, req *http.Request, code int, msg string, data any) {
	resp := ResponseBody{
		Code: code,
		Msg:  msg,
		Data: data,
	}

	jsonBytes, err := json.Marshal(resp)
	if err != nil {
		logger.Infof("json marshal error: %+v", err)

		_, _ = w.Write([]byte("internal server error"))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonBytes)
}
