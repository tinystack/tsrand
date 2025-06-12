# TSRand

[![English](https://img.shields.io/badge/Language-English-blue.svg)](README.md) [![中文](https://img.shields.io/badge/Language-中文-red.svg)](README_CN.md)

一个用于 Go 语言的密码学安全随机数生成库，提供线程安全的整数、字符串和 UUID 随机生成器，具有可靠的降级机制。

[![Go Report Card](https://goreportcard.com/badge/github.com/tinystack/tsrand)](https://goreportcard.com/report/github.com/tinystack/tsrand)
![Go Version](https://img.shields.io/badge/go%20version-%3E=1.18-61CFDD.svg?style=flat-square)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/tinystack/tsrand)](https://pkg.go.dev/mod/github.com/tinystack/tsrand)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

## ✨ 特性

- 🔐 **密码学安全**：使用 `crypto/rand` 作为主要源，`math/rand` 作为降级备选
- 🚀 **高性能**：通过对象池和高效算法优化（整数生成约 350 ns/op）
- 🧵 **线程安全**：所有函数支持并发访问，无需额外锁机制
- 🎯 **丰富的 API**：针对不同使用场景的全面函数集
- 🌍 **Unicode 支持**：通过 `CustomString` 完全支持国际字符
- 🛡️ **视觉安全**：`VisibleString` 排除易混淆字符（0, O, I, l, 1）
- ⚡ **内存优化**：对象池技术减少 GC 压力
- 🔄 **可靠降级**：自动降级机制确保生成永不失败

## 📦 安装

```bash
go get -u github.com/tinystack/tsrand
```

## 🚀 快速开始

```go
package main

import (
    "fmt"
    "github.com/tinystack/tsrand"
)

func main() {
    // 生成随机整数
    fmt.Printf("Int32: %d\n", rand.Int32())
    fmt.Printf("Range: %d\n", rand.RangeInt(1, 100))

    // 生成随机字符串
    fmt.Printf("String: %s\n", rand.String(10))
    fmt.Printf("Visible: %s\n", rand.VisibleString(8))
    fmt.Printf("UUID: %s\n", rand.UUID())
}
```

## 📚 API 参考

### 整数生成

| 函数       | 描述                               | 示例            |
| ---------- | ---------------------------------- | --------------- |
| `Int32()`  | [0, MaxInt32) 范围内的随机 int32   | `rand.Int32()`  |
| `Int64()`  | [0, MaxInt64) 范围内的随机 int64   | `rand.Int64()`  |
| `Uint32()` | [0, MaxUint32) 范围内的随机 uint32 | `rand.Uint32()` |
| `Uint64()` | [0, MaxUint64) 范围内的随机 uint64 | `rand.Uint64()` |
| `Int()`    | [0, MaxInt) 范围内的随机 int       | `rand.Int()`    |
| `Uint()`   | [0, MaxUint) 范围内的随机 uint     | `rand.Uint()`   |

### 范围生成

| 函数                       | 描述                           | 示例                                    |
| -------------------------- | ------------------------------ | --------------------------------------- |
| `RangeInt(min, max)`       | [min, max) 范围内的随机 int    | `rand.RangeInt(1, 100)`                 |
| `RangeInt64(min, max)`     | [min, max) 范围内的随机 int64  | `rand.RangeInt64(1000, 9999)`           |
| `RangeUint32(min, max)`    | [min, max) 范围内的随机 uint32 | `rand.RangeUint32(10, 50)`              |
| `RangeUint64(min, max)`    | [min, max) 范围内的随机 uint64 | `rand.RangeUint64(100, 1000)`           |
| `RangeIntSafe(min, max)`   | 带错误返回的范围 int           | `n, err := rand.RangeIntSafe(1, 100)`   |
| `RangeInt64Safe(min, max)` | 带错误返回的范围 int64         | `n, err := rand.RangeInt64Safe(1, 100)` |

### 字符串生成

| 函数                            | 描述           | 字符集            | 示例                              |
| ------------------------------- | -------------- | ----------------- | --------------------------------- |
| `String(length)`                | 字母数字字符串 | 0-9, a-z, A-Z     | `rand.String(10)`                 |
| `VisibleString(length)`         | 无混淆字符     | 排除 0,O,I,l,1    | `rand.VisibleString(8)`           |
| `AlphaString(length)`           | 仅字母         | a-z, A-Z          | `rand.AlphaString(10)`            |
| `NumericString(length)`         | 仅数字         | 0-9               | `rand.NumericString(6)`           |
| `LowercaseString(length)`       | 仅小写字母     | a-z               | `rand.LowercaseString(8)`         |
| `UppercaseString(length)`       | 仅大写字母     | A-Z               | `rand.UppercaseString(8)`         |
| `CustomString(charset, length)` | 自定义字符集   | 用户定义          | `rand.CustomString("ABC123", 10)` |
| `UUID()`                        | 标准 UUID v4   | 十六进制 + 连字符 | `rand.UUID()`                     |

## 🎯 使用场景

### 🔐 安全应用

```go
// 密码生成（避免混淆字符）
password := rand.VisibleString(16)

// API 密钥生成
apiKey := fmt.Sprintf("ak_%s_%s",
    rand.LowercaseString(8), rand.String(16))

// 会话令牌
sessionToken := rand.String(32)

// 加密盐值
salt := rand.String(24)
```

### 💻 Web 开发

```go
// 验证码
verifyCode := rand.NumericString(6)

// 短链接 ID（无混淆字符）
shortID := rand.VisibleString(8)

// 用户 ID
userID := fmt.Sprintf("U%s%s",
    rand.UppercaseString(2), rand.NumericString(8))

// 邀请码
inviteCode := fmt.Sprintf("%s-%s",
    rand.UppercaseString(4), rand.UppercaseString(4))
```

### 🎮 游戏开发

```go
// 骰子模拟
dice1 := rand.RangeInt(1, 7)
dice2 := rand.RangeInt(1, 7)

// 随机装备属性
attack := rand.RangeInt(50, 150)
durability := rand.RangeInt(80, 100)

// 玩家 ID
playerID := fmt.Sprintf("P%s%s",
    rand.UppercaseString(2), rand.NumericString(6))

// 地图种子
mapSeed := rand.Int64()
```

### 🌍 国际化支持

```go
// 中文字符
chineseStr := rand.CustomString("天地玄黄宇宙洪荒", 5)

// 希腊字母
greekStr := rand.CustomString("αβγδεζηθικλμν", 8)

// 自定义表情符号
emojiStr := rand.CustomString("😀😁😂🤣😃😄", 3)
```

## 📊 性能表现

| 操作         | 性能       | 内存      | 分配次数      |
| ------------ | ---------- | --------- | ------------- |
| `Int32()`    | ~350 ns/op | 48 B/op   | 3 allocs/op   |
| `Int64()`    | ~344 ns/op | 48 B/op   | 3 allocs/op   |
| `RangeInt()` | ~350 ns/op | 48 B/op   | 3 allocs/op   |
| `String(10)` | ~3.6 μs/op | ~688 B/op | ~27 allocs/op |
| `UUID()`     | ~460 ns/op | 64 B/op   | 2 allocs/op   |

_基准测试运行环境：Intel Core i7-9750H @ 2.60GHz_

### 性能扩展性

- **整数生成**：恒定时间 ~350 ns/op
- **字符串生成**：线性扩展 ~200-400 ns/字符
- **并发访问**：完全线程安全，无性能损失
- **内存优化**：对象池技术减少 GC 压力

## 🔧 高级用法

### 错误处理

```go
// 安全范围函数对无效范围返回错误
if n, err := rand.RangeIntSafe(10, 100); err != nil {
    log.Printf("无效范围: %v", err)
} else {
    fmt.Printf("生成结果: %d", n)
}
```

### 复杂密码生成

```go
// 分层密码，保证各类字符的存在
upper := rand.UppercaseString(4)
lower := rand.LowercaseString(4)
numbers := rand.NumericString(4)
special := rand.CustomString("!@#$%^&*", 2)

// 组合并打乱
parts := []string{upper, lower, numbers, special}
// ... 打乱逻辑 ...
password := strings.Join(parts, "")
```

### 商业应用

```go
// 订单号
orderNum := fmt.Sprintf("ORD%s%s",
    rand.NumericString(8), rand.UppercaseString(4))

// 交易 ID
txnID := fmt.Sprintf("TXN_%s_%s",
    rand.NumericString(10), rand.String(8))

// 产品序列号
serialNum := fmt.Sprintf("%s-%s-%s",
    rand.UppercaseString(4),
    rand.NumericString(4),
    rand.UppercaseString(4))
```

## 🧪 示例代码

查看 [`examples/`](examples/) 目录获取完整示例：

- **[`main.go`](examples/main.go)**：完整功能演示
- **[`performance_demo.go`](examples/performance_demo.go)**：性能分析
- **[`README.md`](examples/README.md)**：详细示例文档

```bash
# 运行完整示例
cd examples
go run main.go

# 运行性能演示
go run performance_demo.go
```

## 🔍 测试

```bash
# 运行所有测试
go test -v

# 运行基准测试
go test -bench=. -benchmem

# 竞态检测测试
go test -race

# 测试覆盖率
go test -cover
```

## 🛡️ 安全性

- **主要源**：`crypto/rand` 提供密码学安全的随机生成
- **降级机制**：当 `crypto/rand` 不可用时自动降级到 `math/rand`
- **线程安全**：所有函数都安全支持并发使用
- **非阻塞**：即使在系统熵池较低时也不会阻塞

## 🎨 设计原则

1. **安全优先**：默认提供密码学安全性
2. **性能优化**：高效的算法和内存管理
3. **开发者友好**：直观的 API 和全面的文档
4. **可靠性**：多重降级机制确保功能可用
5. **可扩展性**：易于通过自定义字符集扩展

## 📋 系统要求

- Go 1.18 或更高版本
- 除 Go 标准库外无外部依赖

## 🤝 贡献

欢迎贡献！请随时提交 Pull Request。对于重大更改，请先开启 issue 讨论您想要更改的内容。

1. Fork 本仓库
2. 创建您的特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交您的更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 🙏 致谢

- Go 团队提供的优秀 `crypto/rand` 和 `math/rand` 包
- 本库的贡献者和用户
- 开源社区的灵感和反馈

## 📞 支持

- 📖 [文档](https://pkg.go.dev/github.com/tinystack/tsrand)
- 🐛 [问题追踪](https://github.com/tinystack/tsrand/issues)
- 💬 [讨论区](https://github.com/tinystack/tsrand/discussions)

---

**用 ❤️ 为 Go 社区打造**
