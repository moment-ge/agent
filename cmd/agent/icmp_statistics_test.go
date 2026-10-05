package main

import (
	pb "github.com/nezhahq/agent/proto"
	ping "github.com/prometheus-community/pro-bing"
	"strings"
	"testing"
	"time"
)

func TestICMPStatisticsPreservePartialLoss(t *testing.T) {
	for _, recv := range []int{0, 3, 5} {
		result := new(pb.TaskResult)
		applyICMPStatistics(result, &ping.Statistics{PacketsSent: 5, PacketsRecv: recv, AvgRtt: 123 * time.Millisecond})
		if !strings.Contains(result.Data, `"sent":5`) || result.Successful != (recv > 0) {
			t.Fatalf("bad result: %+v", result)
		}
		if recv > 0 && result.Delay != 123 {
			t.Fatal("incorrect latency units")
		}
	}
	result := new(pb.TaskResult)
	applyICMPStatistics(result, &ping.Statistics{})
	if result.Successful || strings.Contains(result.Data, "nezha_icmp_v1") {
		t.Fatal("invented measurement without sent packets")
	}
}
