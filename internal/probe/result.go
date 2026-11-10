package probe

import "time"

// Result 一条测量记录
type Result struct {
	// 测量发生时间戳 trace 是 JSONL 流，几十天后回看时得知道每条是什么时候测的
	Ts time.Time `json:"ts"`

	// 测量目标 多目标并发时，所有目标的结果混在同一个流里，靠它区分
	Target string `json:"target"`

	// 这条记录测的是哪一段：dns / connect / total
	Phase Phase `json:"phase"`

	// 这一段测量的耗时 毫秒
	RttMs float64 `json:"rtt_ms"`

	// 成功/失败
	Ok bool `json:"ok"`

	// 失败原因（成功时为 ""）
	Err string `json:"err,omitempty"`
}

// Phase phase 枚举
type Phase string

const (
	PhaseDns     Phase = "dns"
	PhaseConnect Phase = "connect"
	PhaseTotal   Phase = "total"
)
