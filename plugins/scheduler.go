// Copyright 2017 Xiaomi, Inc.
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

package plugins

import (
	"context"
	"encoding/json"
	"log"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/geekerlw/falcon-agent/g"
	"github.com/open-falcon/falcon-plus/common/model"
	"github.com/toolkits/file"
)

type PluginScheduler struct {
	Ticker *time.Ticker
	Plugin *Plugin
	Quit   chan struct{}
}

func NewPluginScheduler(p *Plugin) *PluginScheduler {
	scheduler := PluginScheduler{Plugin: p}
	scheduler.Ticker = time.NewTicker(time.Duration(p.Cycle) * time.Second)
	scheduler.Quit = make(chan struct{})
	return &scheduler
}

func (this *PluginScheduler) Schedule() {
	go func() {
		for {
			select {
			case <-this.Ticker.C:
				PluginRun(this.Plugin)
			case <-this.Quit:
				this.Ticker.Stop()
				return
			}
		}
	}()
}

func (this *PluginScheduler) Stop() {
	close(this.Quit)
}

func PluginRun(plugin *Plugin) {

	timeout := plugin.Cycle*1000 - 500
	fpath := filepath.Join(g.Config().Plugin.Dir, plugin.FilePath)

	if !file.IsExist(fpath) {
		log.Println("no such plugin:", fpath)
		return
	}

	debug := g.Config().Debug
	if debug {
		log.Println(fpath, "running...")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Millisecond)
	defer cancel() // The cancel should be deferred so resources can be clean up

	// create the command with our context
	cmd := exec.CommandContext(ctx, fpath)

	if debug {
		log.Println("plugin started:", fpath)
	}

	stdout, err := cmd.Output()

	if ctx.Err() == context.DeadlineExceeded {
		log.Println("[INFO] timeout and kill process", fpath, "successfully")
		return
	}

	if err != nil {
		log.Println("[ERROR] exec plugin", fpath, "fail. error:", err)
		return
	}

	// exec successfully
	if len(stdout) == 0 {
		if debug {
			log.Println("[DEBUG] stdout of", fpath, "is blank")
		}
		return
	}

	var metrics []*model.MetricValue
	err = json.Unmarshal(stdout, &metrics)
	if err != nil {
		log.Printf("[ERROR] json.Unmarshal stdout of %s fail. error:%s stdout: \n%s\n", fpath, err, string(stdout[:]))
		return
	}

	g.SendToTransfer(metrics)
}
