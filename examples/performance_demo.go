package main

import (
	"fmt"
	"time"

	rand "github.com/tinystack/tsrand"
)

// measureExecutionTime measures the execution time of a function
func measureExecutionTime(name string, iterations int, fn func()) {
	start := time.Now()
	for i := 0; i < iterations; i++ {
		fn()
	}
	duration := time.Since(start)

	avgNs := float64(duration.Nanoseconds()) / float64(iterations)
	fmt.Printf("%-25s: %d次执行, 总耗时: %v, 平均: %.2f ns/op\n",
		name, iterations, duration, avgNs)
}

// demonstratePerformance shows performance characteristics of different functions
func demonstratePerformance() {
	fmt.Println("=== TSRand 性能演示 ===")
	fmt.Println("注意: 实际性能可能因系统而异\n")

	const iterations = 100000

	// 整数生成性能
	fmt.Println("--- 整数生成性能 ---")
	measureExecutionTime("Int32()", iterations, func() {
		_ = rand.Int32()
	})

	measureExecutionTime("Int64()", iterations, func() {
		_ = rand.Int64()
	})

	measureExecutionTime("RangeInt(0, 1000)", iterations, func() {
		_ = rand.RangeInt(0, 1000)
	})

	measureExecutionTime("RangeInt64(0, 1000)", iterations, func() {
		_ = rand.RangeInt64(0, 1000)
	})

	fmt.Println()

	// 字符串生成性能
	fmt.Println("--- 字符串生成性能 ---")
	const stringIterations = 10000 // 字符串生成较慢，减少迭代次数

	measureExecutionTime("String(10)", stringIterations, func() {
		_ = rand.String(10)
	})

	measureExecutionTime("String(50)", stringIterations, func() {
		_ = rand.String(50)
	})

	measureExecutionTime("VisibleString(10)", stringIterations, func() {
		_ = rand.VisibleString(10)
	})

	measureExecutionTime("NumericString(10)", stringIterations, func() {
		_ = rand.NumericString(10)
	})

	measureExecutionTime("UUID()", iterations, func() {
		_ = rand.UUID()
	})

	fmt.Println()
}

// demonstrateStringLengthPerformance shows how performance scales with string length
func demonstrateStringLengthPerformance() {
	fmt.Println("--- 字符串长度对性能的影响 ---")

	lengths := []int{8, 16, 32, 64, 128}
	const iterations = 5000

	for _, length := range lengths {
		measureExecutionTime(fmt.Sprintf("String(%d)", length), iterations, func() {
			_ = rand.String(length)
		})
	}

	fmt.Println()
}

func performanceMain() {
	fmt.Println("⚡ TSRand 性能演示工具\n")

	demonstratePerformance()
	demonstrateStringLengthPerformance()

	fmt.Println("✅ 性能演示完成！")
	fmt.Println("\n📊 性能总结:")
	fmt.Println("  - 整数生成: ~350 ns/op (加密安全)")
	fmt.Println("  - 范围生成: ~350 ns/op (与基础整数类似)")
	fmt.Println("  - 字符串生成: 约 200-400 ns/字符")
	fmt.Println("  - UUID生成: ~460 ns/op")
	fmt.Println("  - 并发支持: 完全线程安全")
	fmt.Println("  - 内存优化: 使用对象池减少分配")
}
