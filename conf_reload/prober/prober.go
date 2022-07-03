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

package prober

import (
	"context"
	"time"

	"github.com/baidu/conf-agent/config"
)

type FetchFileResult struct {
	Name    string
	Version string
	Content []byte
}

type commonConfig struct {
	BFECluster string

	ConfTaskHeaders map[string]string
	ConfTaskTimeout time.Duration
}

type Task interface {
	FetchConfFiles(ctx context.Context) ([]*FetchFileResult, error)
}

type Prober struct {
	tasks []Task
}

func (prober *Prober) Probe(ctx context.Context) ([]*FetchFileResult, error) {
	r