# Go-Box AI Usage Guide

## Overview

**go-box** is a Go library for creating ASCII/Unicode boxes in command-line applications. It provides simple APIs to draw decorative boxes around text with support for Unicode characters, emojis, ANSI color codes, and multiple border styles.

## Key Features

- 🎨 **Multiple Border Styles**: Round (Unicode), Double (Unicode), ASCII
- 🌍 **Unicode Support**: Proper width calculation for emojis and CJK characters
- 🎨 **ANSI Color Support**: Automatic filtering of color codes for perfect alignment
- ⚙️ **Customizable Options**: Padding, indentation, border styles
- 📏 **Accurate Width Calculation**: Uses `github.com/mattn/go-runewidth` for precise display width

## Core API Functions

### 1. `PrintBox(lines []string)`
The simplest way to create a box with default round style borders.

```go
import "github.com/afeiship/go-box"

lines := []string{
    "Hello, World!",
    "This is a default box",
}
box.PrintBox(lines)
```

### 2. `PrintASCIIBox(lines []string)`
Creates a box using pure ASCII characters for maximum compatibility.

```go
lines := []string{
    "ASCII compatible box",
    "Works on any terminal",
}
box.PrintASCIIBox(lines)
```

### 3. `PrintBoxWithOptions(lines []string, opts *BoxOptions)`
Full control over box appearance with custom options.

```go
opts := &box.BoxOptions{
    Padding:     1,           // Inner padding lines
    BorderStyle: "double",    // "round", "double", "ascii"
    Indent:      4,           // Left indentation spaces
}
box.PrintBoxWithOptions(lines, opts)
```

### 4. `StringWidth(s string) int`
Utility function to calculate display width of a string, automatically filtering ANSI codes.

```go
width := box.StringWidth("🚀 Hello") // Returns 10 (2 for emoji + 8 for text)
```

## Configuration Options

### BoxOptions Structure
```go
type BoxOptions struct {
    Padding     int    // Number of empty padding lines inside the box
    BorderStyle string // Border style: "round", "double", "ascii"
    Indent      int    // Number of spaces to indent the entire box
}
```

### Border Styles

1. **"round"** (Default) - Unicode round corners
   ```
   ┌─────────────────┐
   │ Round corners   │
   └─────────────────┘
   ```

2. **"double"** - Unicode double lines
   ```
   ╔═════════════════╗
   ║ Double lines    ║
   ╚═════════════════╝
   ```

3. **"ascii"** - Pure ASCII characters
   ```
   +-----------------+
   | ASCII only      |
   +-----------------+
   ```

## Usage Patterns for AI

### 1. Error Messages
```go
func showError(message string) {
    lines := []string{
        "\x1b[31m❌ ERROR\x1b[0m",
        message,
    }
    box.PrintBoxWithOptions(lines, &box.BoxOptions{
        BorderStyle: "double",
        Padding:     0,
    })
}
```

### 2. Success Messages
```go
func showSuccess(message string) {
    lines := []string{
        "\x1b[32m✅ SUCCESS\x1b[0m",
        message,
    }
    box.PrintBox(lines)
}
```

### 3. Info Boxes with Padding
```go
func showInfo(title, content string) {
    lines := []string{
        fmt.Sprintf("\x1b[34mℹ️  %s\x1b[0m", title),
        "",
        content,
    }
    box.PrintBoxWithOptions(lines, &box.BoxOptions{
        Padding:     1,
        BorderStyle: "round",
    })
}
```

### 4. Menu Options
```go
func showMenu(options []string) {
    lines := make([]string, len(options)+1)
    lines[0] = "\x1b[1m📋 Select an option:\x1b[0m"
    for i, opt := range options {
        lines[i+1] = fmt.Sprintf("  %d. %s", i+1, opt)
    }
    box.PrintBoxWithOptions(lines, &box.BoxOptions{
        Indent:      2,
        BorderStyle: "ascii",
    })
}
```

### 5. Nested Content
```go
func showNestedContent() {
    outer := &box.BoxOptions{Indent: 0, BorderStyle: "double"}
    inner := &box.BoxOptions{Indent: 4, BorderStyle: "round"}

    // Outer box
    outerLines := []string{
        "📦 Package Information",
        "",
        "Dependencies and version info:",
    }
    box.PrintBoxWithOptions(outerLines, outer)

    // Inner nested box
    innerLines := []string{
        "go-box v1.0.7",
        "runewidth v0.0.19",
    }
    box.PrintBoxWithOptions(innerLines, inner)
}
```

## Best Practices

### 1. Handle Empty Input
```go
func safeBoxPrint(lines []string) {
    if len(lines) == 0 {
        box.PrintBox([]string{"(empty)"})
        return
    }
    box.PrintBox(lines)
}
```

### 2. Use Consistent Styling
```go
type Theme struct {
    ErrorStyle   *box.BoxOptions
    SuccessStyle *box.BoxOptions
    InfoStyle    *box.BoxOptions
}

var DefaultTheme = Theme{
    ErrorStyle: &box.BoxOptions{BorderStyle: "double"},
    SuccessStyle: &box.BoxOptions{BorderStyle: "round"},
    InfoStyle: &box.BoxOptions{BorderStyle: "ascii"},
}
```

### 3. Combine with ANSI Colors
```go
// Color constants
const (
    ColorRed    = "\x1b[31m"
    ColorGreen  = "\x1b[32m"
    ColorYellow = "\x1b[33m"
    ColorReset  = "\x1b[0m"
)

func showColoredBox(lines []string, colors []string) {
    coloredLines := make([]string, len(lines))
    for i, line := range lines {
        if i < len(colors) {
            coloredLines[i] = colors[i] + line + ColorReset
        } else {
            coloredLines[i] = line
        }
    }
    box.PrintBox(coloredLines)
}
```

## Common Use Cases

### 1. CLI Help Text
```go
func showHelp() {
    lines := []string{
        "\x1b[1m📖 Command Help\x1b[0m",
        "",
        "Usage: myapp [command] [options]",
        "",
        "Commands:",
        "  help     - Show this help message",
        "  version  - Display version info",
        "  run      - Execute the application",
    }
    box.PrintBoxWithOptions(lines, &box.BoxOptions{
        BorderStyle: "double",
        Padding:     0,
    })
}
```

### 2. Progress Indicators
```go
func showProgress(current, total int, message string) {
    percentage := float64(current) / float64(total) * 100
    bar := strings.Repeat("█", int(percentage/5)) + strings.Repeat("░", 20-int(percentage/5))

    lines := []string{
        fmt.Sprintf("\x1b[36m📊 Progress\x1b[0m"),
        fmt.Sprintf("%s %.1f%%", bar, percentage),
        message,
    }
    box.PrintBoxWithOptions(lines, &box.BoxOptions{
        BorderStyle: "round",
        Indent:      2,
    })
}
```

### 3. Data Display Tables
```go
func showTable(headers []string, rows [][]string) {
    // Calculate column widths...
    lines := []string{}

    // Add headers
    headerLine := strings.Join(headers, " | ")
    lines = append(lines, "\x1b[1m"+headerLine+"\x1b[0m")
    lines = append(lines, strings.Repeat("─", len(headerLine)))

    // Add data rows
    for _, row := range rows {
        lines = append(lines, strings.Join(row, " | "))
    }

    box.PrintBoxWithOptions(lines, &box.BoxOptions{
        BorderStyle: "ascii",
        Padding:     0,
    })
}
```

## Integration Tips

### 1. Combine with Other CLI Libraries
go-box works well with popular CLI frameworks like Cobra, Flag, etc.

### 2. Terminal Detection
For better user experience, detect terminal capabilities:
```go
import "golang.org/x/term"

func isSmartTerminal() bool {
    return term.IsTerminal(int(os.Stdout.Fd()))
}

func smartPrint(lines []string) {
    if isSmartTerminal() {
        box.PrintBox(lines)
    } else {
        // Fallback for non-terminal output
        for _, line := range lines {
            fmt.Println(line)
        }
    }
}
```

### 3. Output Redirection
Consider when output is redirected to files:
```go
func printBoxOrPlain(lines []string) {
    // Check if output is going to a terminal
    if term.IsTerminal(int(os.Stdout.Fd())) {
        box.PrintBox(lines)
    } else {
        // Plain output for files/pipes
        for _, line := range lines {
            fmt.Println(line)
        }
    }
}
```

## Troubleshooting

### Common Issues and Solutions

1. **Misaligned borders with ANSI colors**: The library automatically handles this, but ensure you're using proper ANSI escape sequences.

2. **Unicode display issues**: Some terminals may not support Unicode. Use ASCII style as fallback.

3. **Width calculation problems**: The library uses `runewidth` for accurate calculation. If issues persist, they might be terminal-specific.

### Debug Mode
Create a debug function to inspect width calculations:
```go
func debugWidth(s string) {
    cleaned := box.StringWidth(s)
    fmt.Printf("Original: %q\n", s)
    fmt.Printf("Display width: %d\n", cleaned)
    fmt.Printf("Runes: %v\n", []rune(s))
}
```

## Examples Repository

Check the `examples/` directory for comprehensive examples:
- `basic/` - Simple usage patterns
- `advanced/` - Complex configurations
- `styles-demo/` - All border styles
- `emoji-test/` - Unicode and emoji support
- `interactive/` - Interactive CLI applications

## Performance Considerations

- The library is optimized for typical CLI usage (dozens of lines, not thousands)
- ANSI code filtering is O(n) where n is string length
- Memory usage is minimal - only stores necessary strings

## Version Compatibility

- Go 1.24+ required
- Uses `golang.org/x/text` and `github.com/mattn/go-runewidth`
- Cross-platform support (Windows, macOS, Linux)

## Conclusion

go-box provides a simple yet powerful way to enhance CLI applications with decorative text boxes. Its Unicode support, ANSI color handling, and flexible styling options make it ideal for creating professional-looking command-line interfaces.