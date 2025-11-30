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
	opts.BorderStyle = "double"  // round, single, double
	opts.Padding = 1
	opts.Indent = 2
	box.PrintBoxWithOptions(lines, opts)
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
```