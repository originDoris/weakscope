# weakscope

> **弱网链路听诊器**：probe 探测 · replay 回放 · diagnose AI 诊断 —— for streaming links on weak networks.

## 这是什么

给流媒体弱网链路做"心电图 + AI 读片"的命令行工具：

- **probe**：对 tcp / srt / quic 目标持续探测 RTT / 丢包 / 抖动 / 积压，JSONL 留档
- **replay**：无 root 用户态代理复现现场弱网（延迟 / 丢包 / 带宽曲线）
- **experiment**：同一弱网 trace 下对照验证传输策略（固定码率 vs BWE 调参）
- **ask / diagnose**：弱网知识库问答（带 RFC 引用）与 trace 根因诊断（AI 只建议命令，不执行）

## Roadmap

| 版本 | 能力 |
|---|---|
| v0.1 | TCP RTT 探测 + 统计（EWMA / 百分位 / jitter）+ JSONL 留档 + 终端报告 |
| v0.2 | 弱网知识库问答（RAG，带引用溯源）+ 评测门禁 |
| v0.3 | 弱网回放代理 + 对照实验报告 |
| v1.0-rc | AI trace 根因诊断 + MCP server |
| v1.0 | QUIC 探针 + HTML 对比报告 + 多平台预编译二进制 |

完整功能清单见 [docs/plan/features.md](docs/plan/features.md)。

## 快速开始

开发中（v0.1，预计 2026-12 上旬）。当前进度看板：[docs/plan/m0.md](docs/plan/m0.md)。

## 学习日志

本项目同时是一次公开的"边学边做"（Java 工程师 → Go / Rust / Python AI 工程化），
每日学习日志见 [LEARNING.md](LEARNING.md)。

## License

Apache-2.0
