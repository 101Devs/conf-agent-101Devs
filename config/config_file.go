
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
	"path"
)

type BasicFile struct {
	// BFECluster is the BFECluster of current instance
	BFECluster string `validate:"required"`
	// ReloadIntervalMs is reload interval in ms
	ReloadIntervalMs int `validate:"min=1"`

	// BFEConfDir is the Dir of BFE Conf
	BFEConfDir string `validate:"required"`
	// BFEMonitorPort is the port of BFE Moinitor, agent will access this port to reload BFE
	BFEMonitorPort int `validate:"required"`
	// BFEReloadTimeoutMs is the timeout of reload BFE request
	BFEReloadTimeoutMs int `validate:"min=1"`

	// ConfServer is api server address
	ConfServer string `validate:"min=1"`
	// ConfTaskHeaders will be carry to api server
	// Authorization should be set
	ConfTaskHeaders map[string]string
	// ConfTaskTimeoutMs is the timeout of conf prober request
	ConfTaskTimeoutMs int `validate:"min=1"`

	// ExtraFileSever is Extra File address
	ExtraFileServer string `validate:"min=1"`
	// ExtraFileTaskHeaders will be carry to extra file server
	// Authorization should be set
	ExtraFileTaskHeaders map[string]string
	// ExtraFileTaskTimeoutMs is the timeout of extra file download request
	ExtraFileTaskTimeoutMs int `validate:"min=1"`
}

type ReloaderConfigFile struct {
	name string
	// BFECluster is the BFECluster of current instance, inherit BasicFile.BFECluster as default value
	BFECluster string `validate:"required"`

	// ConfDir is the reloadr conf dir, BasicFile.BFEConfDir join ConfDir is the conf root dir
	// inherit reloader map's key as default value
	ConfDir string `validate:"min=1"`
	// BFEReloadAPI is the reload api of bfe, with /reload/ prefix all the time
	BFEReloadAPI string `validate:"min=1"`

	// optional, inherit BasicConfig if not set
	BFEReloadTimeoutMs int `validate:"min=1"`
	ReloadIntervalMs   int `validate:"min=1"`

	// CopyFiles is the file/directory which will be copy from default conf dir to newer version conf dir
	// many conf can't fetch from conf file, newer version conf dir show inherit them so bfe can startup aftert stop
	CopyFiles []string

	// NormalFileTasks is the list of NormalFileTask
	// NormalFileTask meaning to conf file and  conf api one to one correspondence
	NormalFileTasks []NormalFileTaskConfigFile
	// MultiKeyFileTasks is the list of MultiKeyFileTask
	// MultiKeyFileTask meaning to conf file and  conf api many to one correspondence
	MultiKeyFileTasks []MultiJSONKeyFileTaskConfigFile
	// ExtraFileTasks is the los of ExtraFile
	// ExtraFile meaning to conf file and  conf api one to one correspondence
	// extra files info can be obtained by parse conf file
	ExtraFileTasks []ExtraFileTaskConfigFile
}

type NormalFileTaskConfigFile struct {
	// ConfAPI use to access to obtain conf file info