# weakscope · 完整功能实现目标（全景清单）

Generated on 2026-11-10 ｜ 版本演进：v0.1 → v1.0（28 周）
定位：**弱网链路听诊器 —— probe 探测 · replay 回放 · diagnose AI 诊断**，for streaming links on weak networks.
配套：《学习路线总纲》（为什么/红线）《weakscope-分阶段学习路径》（怎么学）《weakscope-M0-功能与进度》（当前阶段作战）
用法：建仓后拷入 `docs/plan/features.md`，每完成一项打勾。**视图二（版本能力声明）对外可见，写进 README 的 Roadmap。**

---

## 视图一：功能矩阵（按模块，标注里程碑与验收标准）

### 模块 A · 探针引擎（probe）

- [ ] A1 TCP RTT 探针：DialContext 计时、超时分类 ｜ M0 ｜ 验收：对已知 host:port 输出稳定 rtt_ms
- [ ] A2 DNS 分离计时：phase=dns / connect / total 三段 ｜ M0 ｜ 验收：坏 DNS 域名只报 dns 错误
- [ ] A3 结果结构 Result{ts,target,phase,rtt_ms,ok,err} + JSON tag ｜ M0
- [ ] A4 表驱动测试 + 可注入 Dialer 小接口 ｜ M0 ｜ 验收：测试不依赖真实网络
- [ ] A5 采样循环：Ticker + duration 控制 + signal 优雅退出 ｜ M0
- [ ] A6 多目标并发：errgroup + channel 汇聚 ｜ M0 ｜ 验收：`-race` 干净
- [ ] A7 ICMP 探针（unprivileged datagram，Linux 可选）｜ M3 可选
- [ ] A8 SRT 探针：cgo 绑 libsrt，build tag 隔离，对齐 srt_bstats 口径（丢包/带宽/积压）｜ M1 ｜ 验收：与 srt 侧统计口径一致
- [ ] A9 QUIC 探针：quic-go，握手/RTT/丢包 ｜ M3 ｜ 验收：对 quic 目标出三指标

### 模块 B · 统计与指标（stats）

- [ ] B1 定长滑窗（slice 实现）｜ M0
- [ ] B2 min / max / mean ｜ M0
- [ ] B3 百分位 p50/p95/p99（nearest-rank）｜ M0
- [ ] B4 EWMA（RFC 6298 SRTT 公式，α 可传参）｜ M0 ｜ 验收：公式与 RFC 推导一致（表驱动用例含手算值）
- [ ] B5 jitter（RFC 3550 §6.3.1 连续差值指数平均）｜ M0
- [ ] B6 trendline BWE 估计器 · Go 版（GCC 同族：延迟梯度 + 过载判定 + AIMD）｜ M1 ｜ 验收：合成反馈下升降方向 100% 正确
- [ ] B7 trendline BWE · Rust 重写版（criterion bench）｜ M2 ｜ 验收：与 Go 版同输入输出一致 + bench 对照数据
- [ ] B8 stats 全模块基准（testing.B + -benchmem + pprof 剖析存档）｜ M0

### 模块 C · trace 存档与互操作

- [ ] C1 JSONL 落盘格式 v1（jq 可直接消费）｜ M0 ｜ 验收：`jq .[0] trace.jsonl` 出结构
- [ ] C2 mahimahi 单列 trace 格式互转器（导入官方 Verizon-LTE 等 trace）｜ M1 ｜ 验收：官方 trace 转入后回放结果合理
- [ ] C3 trace 元数据（目标/时间/环境/工具版本），版本化 schema ｜ M1

### 模块 D · 弱网回放（replay）

- [ ] D1 userspace UDP 代理：延迟注入 + 令牌桶限速（按带宽曲线）｜ M1 ｜ 验收：无 root 运行；目标吞吐贴曲线
- [ ] D2 丢包模型：均匀丢包 + Gilbert-Elliott 突发丢包 ｜ M1 ｜ 验收：注入 5% 实测 4.5-5.5%
- [ ] D3 TCP 回放（语义差异文档化：终结重传，RTT 口径与 netem 不同）｜ M1
- [ ] D4 回放配置（--loss/--delay/--bw-profile/--ge-burst）｜ M1

### 模块 E · 对照实验（experiment）

- [ ] E1 实验驱动脚本：对照组固定码率 vs 实验组 BWE 反馈调参推流（ffmpeg，经同一回放代理到 SRS）｜ M1 ｜ 验收：同一 trace 重复 10 组可复现
- [ ] E2 指标采集：卡顿次数/码率轨迹/恢复时间/p95 延迟（ffprobe）｜ M1
- [ ] E3 实验报告：≥3 张数据图 ｜ M1 ｜ 验收：博文③直接引用
- [ ] E4 ffmpeg 运行时改码率 ADR 归档并实现所选路径 ｜ M1

### 模块 F · 报告（report）

- [ ] F1 终端表格输出（JSONL → 表格）｜ M0
- [ ] F2 HTML 单文件对比报告（两次 trace/两个实验并排 + 图）｜ M3 ｜ 验收：单文件无外部依赖可离线打开

### 模块 G · AI 诊断层（Python）

- [ ] G1 语料工程：RFC 3550/6298/9000、GCC draft、SRT/SRS 文档清洗 + 章节感知分块 + 元数据 ｜ M0.5 ｜ 验收：语料零公司资产
- [ ] G2 混合检索：SQLite FTS5（BM25）+ sqlite-vec（向量）+ RRF 融合 ｜ M0.5
- [ ] G3 Rerank（bge-reranker）+ 引用溯源（[RFC3550 §6.3.1] 式引用回显原文）｜ M0.5 ｜ 验收：答案每条关键论断有引用
- [ ] G4 `ask` 命令端到端 ｜ M0.5 ｜ 验收："SRT 的 TSBPD 是什么"返回带引用正确答案
- [ ] G5 RAG 评测：50 条 golden QA + recall@k + 引用命中率 + CI 门禁 ｜ M0.5 ｜ 验收：recall@5 ≥ 0.8，评测不过不许合并
- [ ] G6 诊断 Agent：trace → 异常检测（规则基线）→ 根因假设 → 查证（RAG 引用 + 建议重放命令）→ 证据链报告 ｜ M2 ｜ 验收：异常注入集 ≥10 例根因命中 ≥80%；**只建议命令不执行**
- [ ] G7 MCP server：diagnose_link / read_trace / run_experiment 三工具 ｜ M2 ｜ 验收：Claude 或 Cursor 真实调用成功
- [ ] G8 诊断评测集与回归 ｜ M2
- [ ] G9 LLM 调用观测/成本看板（延迟与成本逐次记录的汇总呈现）｜ M3
- [ ] G10 （远期可选）LLM 驱动 ABR 参数搜索：只输出建议参数组，不执行 ｜ M3+

### 模块 H · CLI 与工程化

- [ ] H1 cobra 子命令骨架：probe / replay / experiment / ask / diagnose / report ｜ M0 占位，后续逐版补实
- [ ] H2 flag 设计统一（--target/--interval/--duration/--timeout）+ 退出码规范（0 成功/2 参数错/3 探测失败）｜ M0-M1
- [ ] H3 slog JSON 日志（stderr）与结果流（stdout JSONL）分离 ｜ M0
- [ ] H4 Makefile（lint/test/build/race/eval）+ golangci-lint ｜ M0
- [ ] H5 CI：test+lint+race，linux/macOS × amd64/arm64 矩阵 ｜ M0
- [ ] H6 goreleaser 发布流水线（4 平台二进制 + tag 自动 Release）｜ M0 snapshot → M1 正式
- [ ] H7 ADR 目录（docs/adr/：Go↔Python 边界、ffmpeg 码率路径、FFI 边界）｜ M0.5 起
- [ ] H8 README：中文 v0（M0）→ 英文版（M3）
- [ ] H9 LEARNING.md 全程公开学习日志 ｜ M0 起每天一行

## 视图二：版本能力演进（对外 Roadmap，写进 README）

| 版本 | 里程碑 | 发布时用户能干什么 |
|---|---|---|
| **v0.1** | M0（第4周） | 对任意 `tcp://host:port` 持续 RTT 体检：DNS/握手分离计时、EWMA/百分位/jitter 统计、JSONL 留档 + 终端报告 |
| **v0.2** | M0.5（第8周） | 加一个"弱网知识库"：自然语言提问，得到带 RFC 引用的答案；项目自带评测门禁 |
| **v0.3** | M1（第14周） | 弱网可复现：导入/录制 trace，用无 root 代理回放注入延迟丢包限速；跑"固定码率 vs BWE 调参"对照实验出报告 |
| **v1.0-rc** | M2（第20周） | AI 读心电图：喂一条 trace 自动出根因诊断报告（证据链+复现建议）；diagnose/read_trace/run_experiment 可被任何 MCP 客户端调用 |
| **v1.0** | M3（第28周） | QUIC 探针、HTML 对比报告、双语 README、预编译多平台二进制、社区病例提交流程 |

## 视图三：明确不做（防 scope creep，对齐总纲红线）

- ❌ ABR 控制器成品（只做实验驱动脚本，不做生产级码率控制）
- ❌ SFU / 流媒体服务器 / 转码 pipeline / 编解码器（那是 SRS/ffmpeg 的地盘）
- ❌ 分布式集群、多机协同（单机 CLI 工具形态）
- ❌ GUI（终端 + 单文件 HTML 封顶）
- ❌ 公司端点/阈值/配置（通用性红线，永不出现在仓库）
- ❌ AI 自动执行任何修复/重放命令（只建议，人执行）
- ❌ Windows 一等公民支持（CI 矩阵不含，issue 来了再说）

## 交付数字（成功标准，非功能项）

- 28 周内 v1.0 发布；累计 ≥6 篇博文（4 篇必交）
- 6-8 个月：≥100 stars 或 ≥1 个外部真实使用者
- 评测：RAG recall@5 ≥0.8 ｜ 诊断命中 ≥80% ｜ 回放丢包误差 ≤0.5pp
- 个人：Go 闭卷自检 13 题 ≥90%（M0）+ 独立完成全流程（M0-M2）

---

*进度记录（每周日更新）：*
- 2026-11-10：清单建立。当前进度：WP0（M0 起步）
