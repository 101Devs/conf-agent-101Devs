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

package conf_reload

import (
	"context"
	"math/rand"
	"time"

	"github.com/baidu/conf-agent/conf_reload/file_store"
	"github.com/baidu/conf-agent/conf_reload/prober"
	"github.com/baidu/conf-agent/conf_reload/trigger"
	"github.com/baidu/conf-agent/config"
	"github.com/baidu/conf-agent/xlog"
)

type Reloader struct {
	// Name is the name of reloader
	Name string
	// ReloadInterval is the interval reloader try to reload
	ReloadInterval time.Duration

	prober    *prober.Prober
	trigger   *trigger.Trigger
	fileStore *file_store.FileStore