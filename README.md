# go-box
> Simple ASCII box drawing for Go CLI apps.

## installation
```sh
go get -u github.com/afeiship/go-box
```

## usage

```go
package main

import "github.com/afeiship/go-box"

func main() {
	// 🎉 极简使用
	lines := []string{
		"✅ Hello, World!",
		"🚀 Go is awesome!",
	}
	box.PrintBox(lines)

	// 🎨 自定义样式
	opts := box.DefaultBoxOptions()
	opts.BorderStyle = "double"  // round, single, double, ascii
	opts.Padding = 1
	opts.Indent = 2
	box.PrintBoxWithOptions(lines, opts)

	// 📱 ASCII 框样式
	box.PrintASCIIBox(lines)
}
```

**效果预览：**

```
┌──────────────────────┐
│ ✅ Hello, World!     │
│ 🚀 Go is awesome!    │
└──────────────────────┘

  ╔═════════════════════════════╗
  ║                             ║
  ║ ✅ Hello, World!            ║
  ║ 🚀 Go is awesome!           ║
  ║                             ║
  ╚═════════════════════════════╝

+------------------------+
| ✅ Hello, World!     |
| 🚀 Go is awesome!    |
+------------------------+
```

## 🎨 边框样式

go-box 支持多种边框样式：

### 1. **round** (默认)
使用 Unicode 圆角字符，适合大多数现代终端。

```
┌──────────────────────┐
│ ✅ Hello, World!     │
└──────────────────────┘
```

### 2. **single**
使用 Unicode 单线字符。

```
┌──────────────────────┐
│ ✅ Hello, World!     │
└──────────────────────┘
```

### 3. **double**
使用 Unicode 双线字符，更加醒目。

```
╔══════════════════════╗
║ ✅ Hello, World!     ║
╚══════════════════════╝
```

### 4. **ascii** ⭐ 新增
使用纯 ASCII 字符，兼容性最好，适用于任何终端。

```
+------------------------+
| ✅ Hello, World!     |
+------------------------+
```

## 📋 API 参考

### PrintBox(lines []string)
使用默认的 round 样式打印框。

### PrintASCIIBox(lines []string) ⭐ 新增
使用 ASCII 样式打印框的便捷函数。

### PrintBoxWithOptions(lines []string, opts *BoxOptions)
使用自定义选项打印框。

**选项配置：**
```go
type BoxOptions struct {
    Padding     int    // 内边距
    BorderStyle string // 边框样式: "round", "single", "double", "ascii"
    Indent      int    // 缩进空格数
}
```