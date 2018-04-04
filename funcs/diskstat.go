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
	"fmt"
	"log"

	"github.com/open-falcon/falcon-plus/common/model"
	"github.com/shirou/gopsutil/disk"
)

func diskStateUpdate() ([]*disk.UsageStat, error) {
	parts, errp := disk.Partitions(false)
	if errp != nil {
		log.Printf("failed to update disk partition state\n")
		return nil, errp
	}

	var usage []*disk.UsageStat

	for _, part := range parts {
		u, erru := disk.Usage(part.Mountpoint)
		if erru != nil {
			log.Printf("failed to update disk usage state\n")
			return nil, erru
		}
		usage = append(usage, u)
	}
	return usage, nil
}

func DiskMetrics() []*model.MetricValue {
	diskInfo, err := diskStateUpdate()
	if err != nil {
		log.Printf("failed to get disk info: %v\n", err)
		return []*model.MetricValue{}
	}

	var diskStat []*model.MetricValue

	for _, du := range diskInfo {
		tag := fmt.Sprintf("diskpath=%s", du.Path)
		diskStat = append(diskStat, GaugeValue("disk.total", du.Total, tag))
		diskStat = append(diskStat, GaugeValue("disk.free", du.Free, tag))
		diskStat = append(diskStat, GaugeValue("disk.used", du.Used, tag))
		diskStat = append(diskStat, GaugeValue("disk.used.percent", du.UsedPercent, tag))
	}

	return diskStat
}
