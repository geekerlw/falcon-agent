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
	"sync"

	"github.com/open-falcon/falcon-plus/common/model"
	"github.com/shirou/gopsutil/cpu"
)

var (
	cpuTime    [2]cpu.TimesStat
	cpuPercent float64
	cpuLock    = new(sync.RWMutex)
)

func CpuStateUpdate() error {
	c, errc := cpu.Times(false)
	if errc != nil {
		return errc
	}

	p, errp := cpu.Percent(0, false)
	if errp != nil {
		return errp
	}

	cpuLock.Lock()
	defer cpuLock.Unlock()

	cpuTime[0] = cpuTime[1]
	cpuTime[1] = c[0]

	cpuPercent = p[0]

	return nil
}

func cpuUser() float64 {
	cpuLock.RLock()
	defer cpuLock.RUnlock()
	return cpuTime[1].User - cpuTime[0].User
}

func cpuSystem() float64 {
	cpuLock.RLock()
	defer cpuLock.RUnlock()
	return cpuTime[1].System - cpuTime[0].System
}

func cpuIdle() float64 {
	cpuLock.RLock()
	defer cpuLock.RUnlock()
	return cpuTime[1].Idle - cpuTime[0].Idle
}

func cpuNice() float64 {
	cpuLock.RLock()
	defer cpuLock.RUnlock()
	return cpuTime[1].Nice - cpuTime[0].Nice
}

func cpuIowait() float64 {
	cpuLock.RLock()
	defer cpuLock.RUnlock()
	return cpuTime[1].Iowait - cpuTime[0].Iowait
}

func cpuIrq() float64 {
	cpuLock.RLock()
	defer cpuLock.RUnlock()
	return cpuTime[1].Irq - cpuTime[0].Irq
}

func cpuSoftirq() float64 {
	cpuLock.RLock()
	defer cpuLock.RUnlock()
	return cpuTime[1].Softirq - cpuTime[0].Softirq
}

func cpuSteal() float64 {
	cpuLock.RLock()
	defer cpuLock.RUnlock()
	return cpuTime[1].Steal - cpuTime[0].Steal
}

func cpuGuest() float64 {
	cpuLock.RLock()
	defer cpuLock.RUnlock()
	return cpuTime[1].Guest - cpuTime[0].Guest
}

func cpuGuestNice() float64 {
	cpuLock.RLock()
	defer cpuLock.RUnlock()
	return cpuTime[1].GuestNice - cpuTime[0].GuestNice
}

func cpuStolen() float64 {
	cpuLock.RLock()
	defer cpuLock.RUnlock()
	return cpuTime[1].Stolen - cpuTime[0].Stolen
}

func CpuMetrics() []*model.MetricValue {

	return []*model.MetricValue{
		GaugeValue("cpu.user", cpuUser()),
		GaugeValue("cpu.system", cpuSystem()),
		GaugeValue("cpu.idle", cpuIdle()),
		GaugeValue("cpu.nice", cpuNice()),
		GaugeValue("cpu.iowait", cpuIowait()),
		GaugeValue("cpu.irq", cpuIrq()),
		GaugeValue("cpu.softirq", cpuSoftirq()),
		GaugeValue("cpu.steal", cpuSteal()),
		GaugeValue("cpu.guest", cpuGuest()),
		GaugeValue("cpu.gusenice", cpuGuestNice()),
		GaugeValue("cpu.stolen", cpuStolen()),
		GaugeValue("cpu.used.percent", cpuPercent),
	}

}
