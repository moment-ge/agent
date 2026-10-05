# ICMP packet-count extension

This fork preserves Nezha's ICMP timing and success/latency semantics. Completed
probes also report the actual packets sent/received in `TaskResult.Data`:

```json
{"nezha_icmp_v1":{"sent":5,"received":4}}
```

The companion `moment-ge/nezha` endpoint computes rolling loss per server from
these counters. Errors without trustworthy packet statistics remain unknown;
legacy dashboards continue to use the normal success and delay fields.

Run `go test ./cmd/agent -run TestICMPStatistics`.

The X-STATUS deployment disables automatic/forced agent updates so that an
upstream binary cannot silently remove packet counters. Upgrade using the
checksummed releases of this fork.
