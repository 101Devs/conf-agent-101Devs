
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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"

	"github.com/baidu/conf-agent/config"
	"github.com/baidu/conf-agent/xhttp"
	"github.com/baidu/conf-agent/xlog"
)

type NormalFileTask struct {
	config config.NormalFileTaskConfig

	commonConfig commonConfig
}

func NewNormalFileTask(c config.NormalFileTaskConfig) (*NormalFileTask, error) {
	return &NormalFileTask{
		config: c,
		commonConfig: commonConfig{
			BFECluster:      c.BFECluster,
			ConfTaskHeaders: c.ConfTaskHeaders,
			ConfTaskTimeout: c.ConfTaskTimeout,
		},
	}, nil
}

func (task *NormalFileTask) FetchConfFiles(ctx context.Context) ([]*FetchFileResult, error) {
	config := task.config
	fileName := config.ConfFileName

	localVersion, err := loadLocalVersion(path.Join(config.ConfDir, fileName))
	if err != nil {
		return nil, err
	}

	// obtain config data
	raw, err := obtainRemoteConfig(ctx, task.commonConfig, config.ConfAPI, localVersion)
	if err != nil {
		return nil, err
	}

	// if no newer config, conf server will return null
	if raw == nil || string(raw) == `null` {
		return nil, nil
	}

	version, err := calculateVersion(raw)
	if err != nil {
		return nil, err
	}

	return []*FetchFileResult{
		{
			Name:    fileName,
			Version: version,
			Content: raw,
		},
	}, nil