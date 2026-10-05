package main

import (
	"fmt"
	pb "github.com/nezhahq/agent/proto"
	ping "github.com/prometheus-community/pro-bing"
)

// Preserve the normal Nezha result while supplying packet counters to dashboards
// that understand this versioned extension. A legacy dashboard ignores Data.
func applyICMPStatistics(result *pb.TaskResult, stat *ping.Statistics) {
	if stat == nil || stat.PacketsSent <= 0 || stat.PacketsRecv < 0 || stat.PacketsRecv > stat.PacketsSent {
		result.Data = "ICMP packet statistics unavailable"
		return
	}
	result.Data = fmt.Sprintf(`{"nezha_icmp_v1":{"sent":%d,"received":%d}}`, stat.PacketsSent, stat.PacketsRecv)
	result.Successful = stat.PacketsRecv > 0
	if result.Successful {
		result.Delay = float32(stat.AvgRtt.Microseconds()) / 1000
	}
}
