# golang-learning

记录一下日常学习golang

# 基础(1周):
- 基础语法
- 变量和常量
- 基础数据类型
- 运算符
- 流程控制
- 数组
- 切片
- map
- 函数

## 总结:
- for range 中不需要的 index/value 用 `_` 占位符去掉
- 修改字符串先转 `[]rune` 或 `[]byte`，修改后再转回 string（重新分配内存）
- 切片扩容：新容量 >= oldCap*2，实际由 growslice 计算（不同元素类型策略不同）
- map 直接赋值指向同一个底层，深拷贝需遍历：`for k,v := range old { new[k]=v }`
- defer 执行顺序：后进先出 (LIFO)，常用于资源清理
- 函数支持多返回值，常返回 (result, error)
- new(T) 返回 *T，make(T) 返回 T（仅用于 slice、map、chan）

# 深入(2周):
- 指针
- 接口
- 结构体与方法
- 包
- JSON 处理
- error
- 依赖管理
- 反射
- 并发
- 并发错误处理
- 网络编程
- 单元测试
- 泛型
- 迭代器

## 总结:
- 指针：`*T` 是指针类型，`&` 取地址，`*` 解引用，常用于避免大对象拷贝
- 结构体：值类型，方法可绑定值接收者或指针接收者（指针可修改原值）
- 接口：`interface{}` 空接口可接收任意类型，底层 eface/iface 结构
- 接口断言：`value, ok := i.(Type)` 用于类型判断
- JSON：`json:"tag"` 定义字段名，`omitempty` 忽略零值，`-` 跳过字段
- error：本质是接口，`errors.New()` 或 `fmt.Errorf()` 创建，支持错误包装 `%w`
- 反射：`reflect.TypeOf()` 获取类型，`reflect.ValueOf()` 获取值，`Kind()` 区分底层类型
- goroutine：`go f()` 启动轻量级线程，由 G-M-P 模型调度
- sync.Mutex：保证并发安全，`Lock()`/`Unlock()` 成对调用，常用 defer Unlock
- sync.WaitGroup：`Add(n)`、`Done()`、`Wait()` 等待一组 goroutine 完成
- errgroup：管理一组 goroutine，任一错误则取消全部，context 传播取消信号
- channel：`make(chan T)` 创建，`close()` 关闭，发送/接收都会阻塞（带缓冲例外）
- select：多 channel 同时监听，`default` 处理非阻塞，`for-select` 常见模式
- TCP：`net.Listen()` 监听，`conn.Accept()` 接受连接，`conn.Read/Write` 读写数据
- 单元测试：文件名 `_test.go`，函数名 `TestXxx(t *testing.T)`，`go test` 运行
- 泛型：`func F[T any](v T) T`，`T ~int` 类型约束，`comparable` 可比较约束

# 常用标准库(ing):
