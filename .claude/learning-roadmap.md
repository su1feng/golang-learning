# Go 学习路线（完整版）

> 时间窗口：现在 → 年后面试（约 2 个月）
> 目标：达到 Go 后端面试要求，争取跳槽互联网
> B 计划：跳槽失败则转正 + 备考

---

## 📊 当前进度：DAY13/52+ 已完成 25%

```
基础阶段  ████████████████████░░░░  100% | DAY1-9 ✅
进阶阶段  ████████░░░░░░░░░░░░░  40%  | DAY10-17
Go 核心    ░░░░░░░░░░░░░░░░░░░░░  0%   | DAY14-23
工程阶段  ░░░░░░░░░░░░░░░░░░░░░  0%   | DAY24-28
网络框架  ░░░░░░░░░░░░░░░░░░░░░  0%   | DAY29-33
数据库    ░░░░░░░░░░░░░░░░░░░░░  0%   | DAY34-37
缓存中间件  ░░░░░░░░░░░░░░░░░░░░░  0%   | DAY38-41
微服务    ░░░░░░░░░░░░░░░░░░░░░  0%   | DAY42-46
云原生监控  ░░░░░░░░░░░░░░░░░░░░░░  0%   | DAY47-51
项目实战  ░░░░░░░░░░░░░░░░░░░░░  0%   | DAY52+
```

---

## 📅 第一阶段：Go 核心（DAY14-23）约2周

> 面试必考，必须精通

| 天数 | 主题 | 面试频率 | 状态 |
|------|------|----------|------|
| DAY14 | channel 进阶、select、关闭原理、nil channel | ⭐⭐⭐⭐⭐ | 📌 |
| DAY15 | sync 包（Mutex、RWMutex、WaitGroup、Once、Atomic、Pool、Cond、Map） | ⭐⭐⭐⭐⭐ | |
| DAY16 | context 包（原理、超时、取消、传值、valueCtx、cancelCtx） | ⭐⭐⭐⭐⭐ | |
| DAY17 | 并发模式（worker pool、生产者消费者、fan-in/fan-out、ticker、timeout） | ⭐⭐⭐⭐⭐ | |
| DAY18 | 接口深入（nil 接口、interface 底层、类型断言、eface/iface、动态派发） | ⭐⭐⭐⭐ | |
| DAY19 | 反射（reflect）、unsafe、泛型入门 | ⭐⭐⭐ | |
| DAY20 | GC 原理（三色标记、写屏障、STW、触发条件、GC 算法） | ⭐⭐⭐⭐⭐ | |
| DAY21 | 逃逸分析、内存优化、sync.Pool 对象复用 | ⭐⭐⭐⭐ | |
| DAY22 | pprof 性能分析（CPU、内存、goroutine、block、mutex、火焰图） | ⭐⭐⭐⭐⭐ | |
| DAY23 | GOMAXPROCS、goroutine 泄漏排查、调试工具（dlv） | ⭐⭐⭐⭐ | |

---

## 📅 第二阶段：工程与网络（DAY24-28）约1周

| 天数 | 主题 | 面试频率 | 状态 |
|------|------|----------|------|
| DAY24 | Go Module、依赖管理、项目结构（layout）、GOPATH vs GOROOT vs GOBIN | ⭐⭐⭐⭐ | |
| DAY25 | 单元测试（testing、testify、mock、覆盖率、基准测试、表驱动测试） | ⭐⭐⭐⭐ | |
| DAY26 | 错误处理最佳实践、结构化日志（zap、logrus）、错误包装 | ⭐⭐⭐⭐ | |
| DAY27 | 代码规范（gofmt、golint、goimports、staticcheck）、godoc 文档 | ⭐⭐⭐ | |
| DAY28 | net/http 原理（Handler、ServeMux、底层实现、连接池） | ⭐⭐⭐⭐⭐ | |

---

## 📅 第三阶段：Web 框架与协议（DAY29-33）约1周

| 天数 | 主题 | 面试频率 | 状态 |
|------|------|----------|------|
| DAY29 | Gin 框架（路由、中间件、参数绑定、渲染、分组） | ⭐⭐⭐⭐⭐ | |
| DAY30 | JSON 处理、Tag 机制、自定义序列化、jsoniter | ⭐⭐⭐⭐⭐ | |
| DAY31 | HTTP 客户端、超时控制、连接池、重试 | ⭐⭐⭐⭐ | |
| DAY32 | TCP/IP 基础、socket、粘包/拆包、长连接 | ⭐⭐⭐⭐ | |
| DAY33 | gRPC + Protobuf（服务定义、编解码、流式） | ⭐⭐⭐⭐⭐ | |

---

## 📅 第四阶段：数据库（DAY34-37）约1周

| 天数 | 主题 | 面试频率 | 状态 |
|------|------|----------|------|
| DAY34 | MySQL 索引（B+树、聚簇/非聚簇、覆盖索引、最左前缀） | ⭐⭐⭐⭐⭐ | |
| DAY35 | 事务（ACID）、隔离级别、MVCC、锁（行锁、间隙锁、临键锁） | ⭐⭐⭐⭐⭐ | |
| DAY36 | 慢 SQL 优化、分页优化、主从复制、binlog、explain | ⭐⭐⭐⭐⭐ | |
| DAY37 | database/sql + GORM 实战（Hook、Preload、事务、连接池） | ⭐⭐⭐⭐⭐ | |

---

## 📅 第五阶段：缓存中间件（DAY38-41）约1周

| 天数 | 主题 | 面试频率 | 状态 |
|------|------|----------|------|
| DAY38 | Redis 数据类型、底层实现（跳表、SDS、dict、ziplist） | ⭐⭐⭐⭐⭐ | |
| DAY39 | Redis 持久化（RDB/AOF）、主从、哨兵、集群（16384槽） | ⭐⭐⭐⭐⭐ | |
| DAY40 | 缓存穿透/击穿/雪崩、缓存一致性、布隆过滤器 | ⭐⭐⭐⭐⭐ | |
| DAY41 | 分布式锁（setnx、Redlock、Zookeeper）、看门狗 | ⭐⭐⭐⭐⭐ | |

---

## 📅 第六阶段：微服务与分布式（DAY42-46）约1周

| 天数 | 主题 | 面试频率 | 状态 |
|------|------|----------|------|
| DAY42 | 服务注册与发现（Consul、Etcd、Nacos、Eureka） | ⭐⭐⭐⭐ | |
| DAY43 | 消息队列（RabbitMQ、Kafka、RocketMQ）、可靠性投递 | ⭐⭐⭐⭐⭐ | |
| DAY44 | 分布式事务（2PC/3PC、TCC、Saga、本地消息表） | ⭐⭐⭐⭐ | |
| DAY45 | 链路追踪（Jaeger、SkyWalking、OpenTelemetry、trace_id/span） | ⭐⭐⭐⭐ | |
| DAY46 | 限流熔断（Sentinel、Hystrix、令牌桶、滑动窗口） | ⭐⭐⭐⭐⭐ | |

---

## 📅 第七阶段：云原生与监控（DAY47-51）约1周

| 天数 | 主题 | 面试频率 | 状态 |
|------|------|----------|------|
| DAY47 | Docker（镜像、容器、Dockerfile、Compose、cgroups、namespace） | ⭐⭐⭐⭐ | |
| DAY48 | Kubernetes（Pod、Deployment、Service、Ingress、StatefulSet） | ⭐⭐⭐⭐ | |
| DAY49 | Prometheus + Grafana（指标收集、PromQL、可视化、告警） | ⭐⭐⭐⭐ | |
| DAY50 | ELK Stack（Elasticsearch、Logstash、Kibana、Filebeat） | ⭐⭐⭐⭐ | |
| DAY51 | CI/CD（Jenkins、GitLab CI）、灰度发布、蓝绿部署 | ⭐⭐⭐⭐ | |

---

## 📅 第八阶段：项目实战（DAY52+）约2周

### Todo REST API（面试核心）

**技术栈**：Gin + GORM + MySQL + Redis + JWT + gRPC

**功能模块**：
1. 用户认证（注册/登录、JWT、密码加密）
2. CRUD + 分页 + 排序 + 搜索
3. Redis 缓存（热点数据、缓存策略、一致性）
4. 中间件（认证、日志、限流、CORS）
5. gRPC 服务（内网通信）
6. 单元测试（mock、覆盖率）
7. Docker 部署
8. Swagger 文档

---

## 📅 第九阶段：进阶补充（考后/入职后）

| 主题 | 内容 |
|------|------|
| MongoDB | 文档数据库、聚合查询 |
| ClickHouse | 列式数据库、OLAP |
| Nginx | 反向代理、负载均衡 |
| Istio/Service Mesh | 服务网格 |
| Elasticsearch | 全文搜索 |
| Memcached | 内存缓存对比 |
| Thrift | RPC 框架对比 |
| Seata | 分布式事务框架 |
| ShardingSphere | 分库分表 |
| OAuth2.0 | 授权框架 |

---

## 🔥 面试高频题清单

### Go 语言
- goroutine vs 线程
- channel 底层实现
- Mutex 实现（自旋、状态）
- context 用途
- GC 原理（三色标记）
- GMP 调度模型
- interface 底层
- slice 扩容机制
- map 并发安全
- defer 执行顺序
- panic/recover 机制
- 逃逸分析
- sync.Pool 用途
- pprof 使用

### 网络
- TCP 三次握手
- TCP 四次挥手
- HTTP 1.1 vs 2.0 vs 3.0
- HTTPS 握手过程
- Cookie vs Session vs JWT
- gRPC vs HTTP/REST
- Protobuf 编码原理

### 数据库
- 索引原理（B+树）
- 聚簇索引 vs 非聚簇索引
- 事务 ACID
- 隔离级别 + MVCC
- 慢 SQL 优化
- 分页优化
- 主从复制

### Redis
- 5 种数据类型及应用场景
- ZSet 跳表原理
- RDB vs AOF
- 缓存穿透/击穿/雪崩
- 分布式锁
- 缓存一致性

### 分布式
- 分布式锁实现
- 分布式事务（TCC、Saga）
- 服务注册与发现
- 消息队列（可靠性、顺序、去重）
- 限流算法
- 链路追踪

### 算法（LeetCode Hot 100）
- 数组：两数之和、三数之和
- 链表：反转链表、环形链表
- 二叉树：遍历、最大深度
- 动态规划：爬楼梯、打家劫舍
- 哈希表：字母异位词
- 双指针：盛水最多
- 滑动窗口：无重复最长子串
- 二分查找：搜索范围

### 系统设计
- 短链接生成
- 限流器
- 消息队列
- 分布式锁
- 分布式事务
- 负载均衡
- KV 存储
- 推送系统
- 秒杀系统

---

## 📚 学习资源

| 资源 | 用途 | 链接 |
|------|------|------|
| 李文周博客 | Go 系统教程 | https://www.liwenzhou.com/ |
| Go by Example | 快速查语法 | https://gobyexample.com/ |
| draveness | Go 源码分析 | https://draveness.me/golang/ |
| 码农桃花源 | Go 底层原理 | https://www.qcrao.com/ |
| Go 语言中文网 | 社区 | https://studygolang.com/ |
| system-design | 系统设计 | https://github.com/donnat/system-design-primer |
| LeetCode | 算法刷题 | https://leetcode.cn/ |
| Go 官方文档 | 权威参考 | https://go.dev/doc/ |
