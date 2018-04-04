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
