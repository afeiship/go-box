// Package box provides simple ASCII box drawing functionality for command-line applications.
package box

import (
	"fmt"
	"strings"
	"github.com/mattn/go-runewidth"
)

// stripANSI removes ANSI escape sequences from a string.
func stripANSI(s string) string {
	var result strings.Builder
	inEscape := false

	for _, r := range s {
		if r == '\x1b' {
			inEscape = true
			continue
		}
		if inEscape && r == 'm' {
			inEscape = false
			continue
		}
		if !inEscape {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// StringWidth calculates the display width of a string, filtering out ANSI color codes.
func StringWidth(s string) int {
	// Remove ANSI escape sequences first
	cleaned := stripANSI(s)

	// Use runewidth package for accurate Unicode width calculation
	return runewidth.StringWidth(cleaned)
}


// PrintBox prints text surrounded by a decorative ASCII box.
// It automatically handles ANSI color codes and calculates proper padding.
func PrintBox(lines []string) {
	if len(lines) == 0 {
		return
	}

	// Calculate maximum width (considering Unicode character display width)
	maxLen := 0
	for _, line := range lines {
		if l := StringWidth(line); l > maxLen {
			maxLen = l
		}
	}

	// Top border
	top := "┌" + strings.Repeat("─", maxLen+2) + "┐"
	fmt.Println(top)

	// Content lines
	for _, line := range lines {
		width := StringWidth(line)
		padding := strings.Repeat(" ", maxLen-width)
		fmt.Printf("│ %s%s │\n", line, padding)
	}

	// Bottom border
	bottom := "└" + strings.Repeat("─", maxLen+2) + "┘"
	fmt.Println(bottom)
}

// BoxOptions contains configuration options for custom box rendering.
type BoxOptions struct {
	Padding     int
	BorderStyle string
	Indent      int
}

// DefaultBoxOptions returns default options for box rendering.
func DefaultBoxOptions() *BoxOptions {
	return &BoxOptions{
		Padding:     0,
		BorderStyle: "round", // round, double, ascii
		Indent:      0,
	}
}

// PrintBoxWithOptions prints text with custom box styling options.
func PrintBoxWithOptions(lines []string, opts *BoxOptions) {
	if len(lines) == 0 {
		return
	}

	if opts == nil {
		opts = DefaultBoxOptions()
	}

	// Add padding if specified
	if opts.Padding > 0 {
		paddedLines := make([]string, 0, len(lines)+2*opts.Padding)
		paddingLine := strings.Repeat(" ", len(lines[0]))
		for i := 0; i < opts.Padding; i++ {
			paddedLines = append(paddedLines, paddingLine)
		}
		paddedLines = append(paddedLines, lines...)
		for i := 0; i < opts.Padding; i++ {
			paddedLines = append(paddedLines, paddingLine)
		}
		lines = paddedLines
	}

	// Calculate maximum width
	maxLen := 0
	for _, line := range lines {
		if l := StringWidth(line); l > maxLen {
			maxLen = l
		}
	}

	// Select border style
	var hBorder, vBorder, tlCorner, trCorner, blCorner, brCorner string
	switch opts.BorderStyle {
	case "double":
		hBorder, vBorder = "═", "║"
		tlCorner, trCorner, blCorner, brCorner = "╔", "╗", "╚", "╝"
	case "ascii":
		hBorder, vBorder = "-", "|"
		tlCorner, trCorner, blCorner, brCorner = "+", "+", "+", "+"
	default: // round
		hBorder, vBorder = "─", "│"
		tlCorner, trCorner, blCorner, brCorner = "┌", "┐", "└", "┘"
	}

	// Add indentation
	indent := strings.Repeat(" ", opts.Indent)

	// Print box
	fmt.Printf("%s%s%s%s\n", indent, tlCorner, strings.Repeat(hBorder, maxLen+2), trCorner)
	for _, line := range lines {
		width := StringWidth(line)
		padding := strings.Repeat(" ", maxLen-width)
		fmt.Printf("%s%s %s%s %s\n", indent, vBorder, line, padding, vBorder)
	}
	fmt.Printf("%s%s%s%s\n", indent, blCorner, strings.Repeat(hBorder, maxLen+2), brCorner)
}

// PrintASCIIBox prints text surrounded by ASCII box characters (+---+).
// This is a convenience function that uses ASCII border style.
func PrintASCIIBox(lines []string) {
	opts := &BoxOptions{
		Padding:     0,
		BorderStyle: "ascii",
		Indent:      0,
	}
	PrintBoxWithOptions(lines, opts)
}
