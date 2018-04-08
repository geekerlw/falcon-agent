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

	"github.com/open-falcon/falcon-plus/common/model"
	"github.com/shirou/gopsutil/mem"
)

var (
	swap    *mem.SwapMemoryStat
	virtual *mem.VirtualMemoryStat
)

func memStateUpdate() (*mem.SwapMemoryStat, *mem.VirtualMemoryStat, error) {
	s, errs := mem.SwapMemory()
	if errs != nil {
		log.Printf("failed to update swap memory state\n")
		return nil, nil, errs
	}
	v, errv := mem.VirtualMemory()
	if errv != nil {
		log.Printf("failed to update virtual memory state\n")
		return nil, nil, errv
	}

	return s, v, nil
}

func MemMetrics() []*model.MetricValue {
	swap, virtual, err := memStateUpdate()
	if err != nil {
		log.Printf("failed to get memory info: %v\n", err)
		return []*model.MetricValue{}
	}

	return []*model.MetricValue{
		GaugeValue("mem.swap.total", swap.Total),
		GaugeValue("mem.swap.used", swap.Used),
		GaugeValue("mem.swap.free", swap.Free),
		GaugeValue("mem.swap.used.percent", swap.UsedPercent),
		GaugeValue("mem.total", virtual.Total),
		GaugeValue("mem.available", virtual.Available),
		GaugeValue("mem.used", virtual.Used),
		GaugeValue("mem.free", virtual.Free),
		GaugeValue("mem.used.percent", virtual.UsedPercent),
	}
}
