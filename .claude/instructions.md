# Claude Code 项目配置

## 角色定位

> 详细角色提示词请查看：`.claude/backend-mentor.md`

---

## 学员背景

- 23岁，2024年6月本科毕业，专业生物工程（非CS）
- 工作：泳池清洁机器人 C++ 中台开发，1.5-2年经验
- 技术栈：C++17、ROS、行为树
- 目标：转到互联网做后端工程师（纯后端或后端+AI 都可）

---

## 学习路线

> 时间窗口：现在 → 年后面试（约 2 个月）
> 目标：达到 Go 后端面试要求，争取跳槽互联网
> B 计划：跳槽失败则转正 + 备考

详细学习路线请查看：`.claude/learning-roadmap.md`

---

## 代码风格

- 使用 Go 的惯用写法
- 变量名使用驼峰命名
- 添加必要的中文注释

---

## Git Commit 规范（Conventional Commits）

### 基本格式
```
type(scope): description
```

### Type 类型

| 前缀 | 含义 | 示例 |
|------|------|------|
| `feat:` | 新功能 | `feat: 添加用户登录接口` |
| `fix:` | 修复 bug | `fix: 修复内存泄漏问题` |
| `docs:` | 文档改动 | `docs: 更新 README` |
| `style:` | 代码格式（不影响逻辑） | `style: 统一缩进` |
| `refactor:` | 重构（不修bug也不加功能） | `refactor: 重构配置加载逻辑` |
| `perf:` | 性能优化 | `perf: 优化查询速度` |
| `test:` | 测试相关 | `test: 添加单元测试` |
| `chore:` | 构建/工具链变动 | `chore: 升级依赖版本` |
| `ci:` | CI/CD 配置 | `ci: 修改 GitHub Actions` |
| `revert:` | 回滚提交 | `revert: 撤销上次提交` |

### 示例
```bash
feat(auth): 添加 JWT 认证
fix(db): 修复连接池泄漏
refactor(api): 重构路由注册
```

---

## 交互偏好

- 解释代码时请简洁明了
- 修改代码前先说明改动内容
- 优先使用中文交流
- 每天学完后帮助总结和规划
