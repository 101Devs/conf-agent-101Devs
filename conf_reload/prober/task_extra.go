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
	"fmt"
	"net/http"
	"strings"

	"github.com/baidu/conf-agent/config"
	"github.com/baidu/conf-agent/xhttp"
	"github.com/baidu/conf-agent/xlog"
	"github.com/ohler55/ojg/oj"
)

type ExtraFileTask struct {
	config config.ExtraFileTaskConfig

	normalFileTask *NormalFileTask
}

func NewExtraFileTask(c config.ExtraFileTaskConfig) (*ExtraFileTask, error) {
	np, err := NewNormalFileTask(c.NormalFileTaskConfig)
	if err != nil {
		return nil, err
	}

	return &ExtraFileTask{
		config: c,

		normalFileTask: np,
	}, nil
}

func (task *ExtraFileTask) FetchConfFiles(ctx context.Context) ([]*FetchFileResult, error) {
	fileList, err := task.normalFileTask.FetchConfFiles(ctx)
	if err != nil {
		return nil, err
	}

	if len(fileList) == 0 {
		return fileList, err
	}

	// analysis file content, obtain extra files
	extraFiles, err := task.obtainExtraFiles(ctx, fileList[0].Content)
	if err != nil {
		return nil, err
	}

	for remotePath, localPath := range ex