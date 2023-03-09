
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

package config

import (
	"fmt"
	"sort"
	"time"

	"github.com/go-playground/validator"
	"github.com/ohler55/ojg/jp"
)

type Config struct {
	Reloaders []*ReloaderConfig
	Logger    *LoggerConfig
}

type ReloaderConfig struct {
	Name string

	ConfDir        string
	ReloadInterval time.Duration

	Trigger TriggerConfig

	CopyFiles []string

	NormalFileTasks       []*NormalFileTaskConfig
	MultiJSONKeyFileTasks []*MultiJSONKeyFileTaskConfig
	ExtraFileFileTasks    []*ExtraFileTaskConfig
}

type NormalFileTaskConfig struct {
	BFECluster string

	ConfDir      string
	ConfAPI      string
	ConfFileName string

	ConfTaskHeaders map[string]string
	ConfTaskTimeout time.Duration
}

func newNormalFileTaskConfig(cf NormalFileTaskConfigFile, rcf ReloaderConfigFile) *NormalFileTaskConfig {
	return &NormalFileTaskConfig{
		BFECluster: rcf.BFECluster,
		ConfDir:    rcf.ConfDir,

		ConfAPI:      cf.ConfServer + cf.ConfAPI,
		ConfFileName: cf.ConfFileName,
