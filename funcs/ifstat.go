package funcs

import (
	"log"

	"github.com/open-falcon/falcon-plus/common/model"
	"github.com/shirou/gopsutil/net"
)

func NetMetrics() []*model.MetricValue {
	netIfs, err := net.IOCounters(true)
	if err != nil {
		log.Printf("failed to get iface io counters, err: %v\n", err)
		return []*model.MetricValue{}
	}

	res := netIfs[0]

	return []*model.MetricValue{
		CounterValue("net.if.bytes.send", res.BytesSent),
		CounterValue("net.if.bytes.recv", res.BytesRecv),
		CounterValue("net.if.packets.send", res.PacketsSent),
		CounterValue("net.if.packets.recv", res.PacketsRecv),
		CounterValue("net.if.err.in", res.Errin),
		CounterValue("net.if.err.out", res.Errout),
		CounterValue("net.if.drop.in", res.Dropin),
		CounterValue("net.if.drop.out", res.Dropout),
		CounterValue("net.if.fifo.in", res.Fifoin),
		CounterValue("net.if.fifo.out", res.Fifoout),
	}
}
