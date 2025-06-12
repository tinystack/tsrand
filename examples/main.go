package main

import (
	"fmt"
	"log"
	"strings"

	rand "github.com/tinystack/tsrand"
)

// demonstrateBasicIntegers shows basic integer generation functions
func demonstrateBasicIntegers() {
	fmt.Println("=== 基础整数生成示例 ===")

	// 生成各种类型的随机整数
	fmt.Printf("Int32:  %d\n", rand.Int32())
	fmt.Printf("Int64:  %d\n", rand.Int64())
	fmt.Printf("Int:    %d\n", rand.Int())
	fmt.Printf("Uint32: %d\n", rand.Uint32())
	fmt.Printf("Uint64: %d\n", rand.Uint64())
	fmt.Printf("Uint:   %d\n", rand.Uint())
	fmt.Println()
}

// demonstrateRangeGeneration shows range-based random generation
func demonstrateRangeGeneration() {
	fmt.Println("=== 范围随机数生成示例 ===")

	// 基本范围生成
	fmt.Printf("RangeInt(1, 100):      %d\n", rand.RangeInt(1, 100))
	fmt.Printf("RangeInt64(1000, 9999): %d\n", rand.RangeInt64(1000, 9999))
	fmt.Printf("RangeUint32(10, 50):   %d\n", rand.RangeUint32(10, 50))
	fmt.Printf("RangeUint64(100, 1000): %d\n", rand.RangeUint64(100, 1000))

	// 使用安全版本（带错误处理）
	if n, err := rand.RangeIntSafe(50, 150); err != nil {
		log.Printf("Error: %v", err)
	} else {
		fmt.Printf("RangeIntSafe(50, 150): %d\n", n)
	}

	// 演示错误处理
	if _, err := rand.RangeIntSafe(100, 50); err != nil {
		fmt.Printf("预期的错误（无效范围）: %v\n", err)
	}

	fmt.Println()
}

// demonstrateBasicStrings shows basic string generation functions
func demonstrateBasicStrings() {
	fmt.Println("=== 基础字符串生成示例 ===")

	// 基本字符串生成
	fmt.Printf("String(20):        %s\n", rand.String(20))
	fmt.Printf("VisibleString(20): %s\n", rand.VisibleString(20))

	// 不同长度的字符串
	fmt.Printf("String(8):         %s\n", rand.String(8))
	fmt.Printf("String(16):        %s\n", rand.String(16))
	fmt.Printf("String(32):        %s\n", rand.String(32))

	// UUID 生成
	fmt.Printf("UUID:              %s\n", rand.UUID())
	fmt.Printf("UUID:              %s\n", rand.UUID())
	fmt.Println()
}

// demonstrateAdvancedStrings shows advanced string generation functions
func demonstrateAdvancedStrings() {
	fmt.Println("=== 高级字符串生成示例 ===")

	// 不同类型的字符串生成
	fmt.Printf("AlphaString(15):     %s\n", rand.AlphaString(15))
	fmt.Printf("NumericString(10):   %s\n", rand.NumericString(10))
	fmt.Printf("LowercaseString(12): %s\n", rand.LowercaseString(12))
	fmt.Printf("UppercaseString(12): %s\n", rand.UppercaseString(12))

	// 自定义字符集
	hexCharset := "0123456789ABCDEF"
	fmt.Printf("CustomString(hex, 16):    %s\n", rand.CustomString(hexCharset, 16))

	binaryCharset := "01"
	fmt.Printf("CustomString(binary, 20): %s\n", rand.CustomString(binaryCharset, 20))

	fmt.Println()
}

// demonstrateVisibleStringAdvantages shows the advantages of VisibleString
func demonstrateVisibleStringAdvantages() {
	fmt.Println("=== VisibleString 优势演示 ===")
	fmt.Println("注意: VisibleString 不包含容易混淆的字符 (0, O, I, l, 1)")

	for i := 0; i < 3; i++ {
		normal := rand.String(15)
		visible := rand.VisibleString(15)
		fmt.Printf("String:        %s\n", normal)
		fmt.Printf("VisibleString: %s\n", visible)
		fmt.Println()
	}
}

// demonstratePasswordGeneration shows practical password generation
func demonstratePasswordGeneration() {
	fmt.Println("=== 实用密码生成示例 ===")

	// 简单密码（仅字母数字）
	simplePassword := rand.String(12)
	fmt.Printf("简单密码 (12位):     %s\n", simplePassword)

	// 可视密码（避免混淆字符）
	visiblePassword := rand.VisibleString(16)
	fmt.Printf("可视密码 (16位):     %s\n", visiblePassword)

	// 复杂密码（包含特殊字符）
	complexCharset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"
	complexPassword := rand.CustomString(complexCharset, 20)
	fmt.Printf("复杂密码 (20位):     %s\n", complexPassword)

	// 分层密码生成
	fmt.Println("\n--- 分层密码生成 ---")
	upper := rand.UppercaseString(4)
	lower := rand.LowercaseString(4)
	numbers := rand.NumericString(4)
	special := rand.CustomString("!@#$%^&*", 2)

	// 随机组合所有部分
	parts := []string{upper, lower, numbers, special}
	shuffled := shuffleStringParts(parts)
	layeredPassword := strings.Join(shuffled, "")
	fmt.Printf("分层密码 (14位):     %s\n", layeredPassword)

	fmt.Println()
}

// shuffleStringParts randomly shuffles string parts (simple demonstration)
func shuffleStringParts(parts []string) []string {
	n := len(parts)
	for i := 0; i < n*2; i++ {
		a := rand.RangeInt(0, n)
		b := rand.RangeInt(0, n)
		parts[a], parts[b] = parts[b], parts[a]
	}
	return parts
}

// demonstrateTokenGeneration shows various token generation patterns
func demonstrateTokenGeneration() {
	fmt.Println("=== 令牌生成示例 ===")

	// API Key 样式
	apiKey := fmt.Sprintf("ak_%s_%s",
		rand.LowercaseString(8),
		rand.String(16))
	fmt.Printf("API Key:           %s\n", apiKey)

	// Session Token
	sessionToken := rand.String(32)
	fmt.Printf("Session Token:     %s\n", sessionToken)

	// 验证码
	verificationCode := rand.NumericString(6)
	fmt.Printf("验证码:             %s\n", verificationCode)

	// 短链接 ID（使用可视字符避免混淆）
	shortLinkId := rand.VisibleString(8)
	fmt.Printf("短链接 ID:         %s\n", shortLinkId)

	// 邀请码
	inviteCode := fmt.Sprintf("%s-%s",
		rand.UppercaseString(4),
		rand.UppercaseString(4))
	fmt.Printf("邀请码:             %s\n", inviteCode)

	fmt.Println()
}

// demonstrateUnicodeSupport shows Unicode character support
func demonstrateUnicodeSupport() {
	fmt.Println("=== Unicode 字符支持示例 ===")

	// 中文字符集
	chineseCharset := "天地玄黄宇宙洪荒日月盈昃辰宿列张寒来暑往秋收冬藏"
	fmt.Printf("中文随机字符串: %s\n", rand.CustomString(chineseCharset, 8))

	// 希腊字母
	greekCharset := "αβγδεζηθικλμνξοπρστυφχψω"
	fmt.Printf("希腊字母字符串: %s\n", rand.CustomString(greekCharset, 10))

	// 表情符号
	emojiCharset := "😀😁😂🤣😃😄😅😆😉😊😋😎😍😘🥰😗😙😚"
	fmt.Printf("表情符号字符串: %s\n", rand.CustomString(emojiCharset, 5))

	fmt.Println()
}

// demonstrateGameApplications shows gaming-related random generation
func demonstrateGameApplications() {
	fmt.Println("=== 游戏应用示例 ===")

	// 模拟骰子
	fmt.Println("--- 骰子模拟 ---")
	for i := 0; i < 5; i++ {
		dice1 := rand.RangeInt(1, 7)
		dice2 := rand.RangeInt(1, 7)
		fmt.Printf("骰子 %d: %d + %d = %d\n", i+1, dice1, dice2, dice1+dice2)
	}

	// 随机装备属性
	fmt.Println("\n--- 随机装备属性 ---")
	weaponNames := []string{"剑", "斧", "弓", "法杖", "匕首"}
	for i := 0; i < 3; i++ {
		weapon := weaponNames[rand.RangeInt(0, len(weaponNames))]
		attack := rand.RangeInt(50, 150)
		durability := rand.RangeInt(80, 100)
		fmt.Printf("武器: %s, 攻击力: %d, 耐久度: %d%%\n", weapon, attack, durability)
	}

	// 玩家ID生成
	fmt.Println("\n--- 玩家ID生成 ---")
	for i := 0; i < 3; i++ {
		playerId := fmt.Sprintf("P%s%s",
			rand.UppercaseString(2),
			rand.NumericString(6))
		fmt.Printf("玩家ID: %s\n", playerId)
	}

	fmt.Println()
}

// demonstrateBusinessApplications shows business-related random generation
func demonstrateBusinessApplications() {
	fmt.Println("=== 商业应用示例 ===")

	// 订单号生成
	orderNumber := fmt.Sprintf("ORD%s%s",
		rand.NumericString(8),
		rand.UppercaseString(4))
	fmt.Printf("订单号:     %s\n", orderNumber)

	// 交易ID
	transactionId := fmt.Sprintf("TXN_%s_%s",
		rand.NumericString(10),
		rand.String(8))
	fmt.Printf("交易ID:     %s\n", transactionId)

	// 优惠券代码
	couponCode := fmt.Sprintf("SAVE%s", rand.UppercaseString(6))
	fmt.Printf("优惠券代码: %s\n", couponCode)

	// 产品序列号
	serialNumber := fmt.Sprintf("%s-%s-%s",
		rand.UppercaseString(4),
		rand.NumericString(4),
		rand.UppercaseString(4))
	fmt.Printf("产品序列号: %s\n", serialNumber)

	// 批次号
	batchNumber := fmt.Sprintf("BATCH_%s_%s",
		rand.NumericString(6),
		rand.UppercaseString(3))
	fmt.Printf("批次号:     %s\n", batchNumber)

	fmt.Println()
}

func main() {
	fmt.Println("🎲 TSRand 完整功能演示\n")
	fmt.Println("这是一个展示 TSRand 包所有功能的综合示例\n")

	// 基础功能演示
	demonstrateBasicIntegers()
	demonstrateRangeGeneration()
	demonstrateBasicStrings()

	// 高级功能演示
	demonstrateAdvancedStrings()
	demonstrateVisibleStringAdvantages()
	demonstratePasswordGeneration()
	demonstrateTokenGeneration()
	demonstrateUnicodeSupport()

	// 应用场景演示
	demonstrateGameApplications()
	demonstrateBusinessApplications()

	fmt.Println("✅ 所有示例演示完成！")
	fmt.Println("\n💡 使用建议:")
	fmt.Println("  - 密码生成: 使用 VisibleString 避免字符混淆")
	fmt.Println("  - API令牌: 使用 String 或 CustomString")
	fmt.Println("  - 验证码: 使用 NumericString")
	fmt.Println("  - 游戏ID: 使用 RangeInt 或组合生成")
	fmt.Println("  - 国际化: 使用 CustomString 支持多语言字符")
}
