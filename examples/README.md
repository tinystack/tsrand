# TSRand 示例集合

这个目录包含了 TSRand 包的各种使用示例，展示了包的完整功能和最佳实践。

## 示例文件说明

### 📁 main.go - 完整功能演示

这是最全面的示例文件，包含了 TSRand 包的所有功能演示：

**运行方式:**

```bash
cd examples
go run main.go
```

**包含内容:**

- 🔢 **基础整数生成**: Int32, Int64, Uint32, Uint64 等
- 📏 **范围随机数生成**: RangeInt, RangeInt64 及其安全版本
- 🔤 **基础字符串生成**: String, VisibleString, UUID
- 🎭 **高级字符串生成**: AlphaString, NumericString, CustomString 等
- 🛡️ **VisibleString 优势演示**: 避免字符混淆的可视字符串
- 🔐 **密码生成示例**: 简单密码、复杂密码、分层密码生成
- 🎫 **令牌生成示例**: API Key, Session Token, 验证码等
- 🌍 **Unicode 支持示例**: 中文、希腊字母、表情符号
- 🎮 **游戏应用示例**: 骰子模拟、装备属性、玩家 ID
- 💼 **商业应用示例**: 订单号、交易 ID、优惠券代码等

### 📊 performance_demo.go - 性能演示

展示不同函数的性能特征和优化效果：

**运行方式:**

```bash
cd examples
go run performance_demo.go
```

**演示内容:**

- ⚡ 各种函数的执行时间测量
- 📈 字符串长度对性能的影响
- 🔄 并发访问性能测试
- 💾 内存使用效率分析

## 🎯 快速开始指南

### 1. 基础使用

```go
import "github.com/tinystack/tsrand"

// 生成随机整数
n := rand.Int32()
range_n := rand.RangeInt(1, 100)

// 生成随机字符串
s := rand.String(10)
uuid := rand.UUID()
```

### 2. 高级功能

```go
// 避免混淆字符的可视字符串
visible := rand.VisibleString(8)

// 自定义字符集
hex := rand.CustomString("0123456789ABCDEF", 16)

// 特定类型字符串
password := rand.UppercaseString(8)
code := rand.NumericString(6)
```

### 3. 错误处理

```go
// 使用安全版本获得错误信息
if n, err := rand.RangeIntSafe(10, 100); err != nil {
    log.Printf("Error: %v", err)
} else {
    fmt.Printf("Generated: %d", n)
}
```

## 🌟 使用场景

### 🔐 安全应用

- **密码生成**: 使用 `VisibleString` 避免字符混淆
- **API 密钥**: 使用 `String` 或 `CustomString`
- **会话令牌**: 使用长度足够的 `String`
- **加密盐值**: 使用 `String` 或 `CustomString`

### 💻 Web 开发

- **验证码**: 使用 `NumericString` 或 `VisibleString`
- **短链接 ID**: 使用 `VisibleString` 避免混淆
- **用户 ID**: 组合 `UppercaseString` 和 `NumericString`
- **邀请码**: 格式化的组合生成

### 🎮 游戏开发

- **玩家 ID**: 组合字母和数字
- **随机事件**: 使用 `RangeInt` 控制概率
- **装备属性**: 使用范围生成随机属性值
- **地图种子**: 使用 `Int64` 生成地图种子

### 📊 数据处理

- **采样**: 使用 `RangeInt` 进行随机采样
- **测试数据**: 生成各种类型的测试数据
- **负载测试**: 生成随机负载模式
- **A/B 测试**: 随机分组用户

## ⚠️ 注意事项

### 安全性

- TSRand 使用 `crypto/rand` 作为主要随机源，提供加密级别的安全性
- 当 `crypto/rand` 不可用时，自动降级到 `math/rand`
- 所有函数都是线程安全的

### 性能

- 整数生成：约 350 ns/op
- 字符串生成：线性扩展，约 200-400 ns/字符
- UUID 生成：约 460 ns/op
- 使用对象池优化内存分配

### 最佳实践

1. **选择合适的函数**: 根据用途选择最适合的生成函数
2. **错误处理**: 在关键应用中使用 `*Safe` 版本函数
3. **字符集选择**: 用户界面相关的场景使用 `VisibleString`
4. **长度规划**: 根据安全需求选择合适的长度
5. **国际化**: 使用 `CustomString` 支持多语言字符

## 🤝 参与贡献

如果您有新的使用场景或改进建议，欢迎：

1. 提交 Issue 描述您的需求
2. 提交 Pull Request 贡献新的示例
3. 分享您的使用经验

## 📚 更多资源

- [项目主页](https://github.com/tinystack/tsrand)
- [API 文档](https://pkg.go.dev/github.com/tinystack/tsrand)
- [性能测试报告](../performance_test.go)
- [单元测试](../*_test.go)
