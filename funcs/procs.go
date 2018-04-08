// Copyright 2018 Steven Lee <geekerlw.gmail.com>
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

package funcs

import (
	"log"
	"strings"

	"github.com/geekerlw/falcon-agent/g"
	"github.com/open-falcon/falcon-plus/common/model"
	gp "github.com/shirou/gopsutil/process"
)

func isProc(p *gp.Process, m map[int]string) bool {
	for key, val := range m {
		if key == 1 {
			// name
			name, _ := p.Name()
			if val == name {
				return true
			}
		} else if key == 2 {
			// cmdline
			cmdline, _ := p.Cmdline()
			if strings.Contains(cmdline, val) {
				return true
			}
		}
	}
	return false
}

func ProcMetrics() (L []*model.MetricValue) {
	reportProcs := g.ReportProcs()
	sz := len(reportProcs)
	if sz == 0 {
		return
	}

	ps, err := gp.Processes()
	if err != nil {
		log.Printf("failed to enum all processes: %v\n", err)
		return
	}

	pslen := len(ps)

	for tags, m := range reportProcs {
		cnt := 0
		var cpuTotal float64
		cpuTotal = 0.0
		var memTotal float32
		memTotal = 0.0

		for i := 0; i < pslen; i++ {
			if isProc(ps[i], m) {
				cnt++
				cpu, _ := ps[i].CPUPercent()
				mem, _ := ps[i].MemoryPercent()
				cpuTotal += cpu
				memTotal += mem
			}
		}

		L = append(L, GaugeValue(g.PROC_NUM, cnt, tags))
		L = append(L, GaugeValue("proc.cpu.percent", cpuTotal, tags))
		L = append(L, GaugeValue("proc.mem.percent", memTotal, tags))
	}
	return
}
