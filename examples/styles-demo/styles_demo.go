package main

import (
	"fmt"
	"github.com/afeiship/go-box"
)

func main() {
	lines := []string{
		"🚀 Rocket launch detected!",
		"✅ Status: Success",
		"📦 Package: dependencies",
		"🔧 Tools: runewidth optimization",
	}

	fmt.Println("🎯 默认圆角框:")
	box.PrintBox(lines)

	fmt.Println("\n💫 ASCII 风格框:")
	box.PrintASCIIBox(lines)

	// 使用自定义选项展示不同样式
	fmt.Println("\n🔷 双线框:")
	opts := &box.BoxOptions{
		Padding:     0,
		BorderStyle: "double",
		Indent:      0,
	}
	box.PrintBoxWithOptions(lines, opts)

	fmt.Println("\n🔸 单线框:")
	opts.BorderStyle = "single"
	box.PrintBoxWithOptions(lines, opts)

	fmt.Println("\n🔹 缩进 + ASCII 风格:")
	opts.BorderStyle = "ascii"
	opts.Indent = 4
	box.PrintBoxWithOptions(lines, opts)

	// 测试包含 ANSI 颜色码的文本
	fmt.Println("\n🌈 包含 ANSI 颜色码:")
	coloredLines := []string{
		"\x1b[31m🔴 红色错误信息\x1b[0m",
		"\x1b[32m🟢 绿色成功信息\x1b[0m",
		"\x1b[33m🟡 黄色警告信息\x1b[0m",
		"\x1b[34m🔵 蓝色信息\x1b[0m",
	}
	box.PrintBox(coloredLines)
}