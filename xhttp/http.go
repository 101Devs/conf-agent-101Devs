
// Copyright (c) 2021 The BFE Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package xhttp

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type HTTPRequest struct {
	Client *http.Client

	Request *http.Request

	Response   *http.Response
	RawContent []byte

	err error
}

func RspCode200Op(h *HTTPRequest) error {
	if statusCode := h.Response.StatusCode; statusCode != 200 {
		return fmt.Errorf("bad StatuCode: %d, Raw: %s", statusCode, h.RawContent)
	}
	return nil
}

type HTTPRequestOp func(*HTTPRequest) error

func RspBodyRawReaderOp(hr *HTTPRequest) error {
	if hr.Response == nil {
		return fmt.Errorf("body is nil")
	}

	hr.RawContent, hr.err = ioutil.ReadAll(hr.Response.Body)
	defer hr.Response.Body.Close()
