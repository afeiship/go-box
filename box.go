// Package box provides simple ASCII box drawing functionality for command-line applications.
package box

import (
	"fmt"
	"strings"
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

	// Calculate display width for each character
	width := 0
	for _, r := range cleaned {
		if r < 32 || r == 127 {
			// Skip control characters
			continue
		}

		// Calculate character width based on Unicode properties
		charWidth := getRuneWidth(r)
		width += charWidth
	}
	return width
}

// getRuneWidth returns the display width of a Unicode rune.
func getRuneWidth(r rune) int {
	// Handle common wide characters and emoji
	switch {
	case r == 0x2705 || // ✅ White heavy check mark
		r == 0x26A0 || // ⚠️ Warning sign (without variation selector)
		r == 0x274C || // ❌ Cross mark
		r == 0x2611 || // ☑️ Ballot box with check
		r == 0x2714 || // ✔️ Heavy check mark
		r == 0x2713 || // ✓ Check mark
		r == 0x1F4E6 || // 📦 Package
		r == 0x1F9EA || // 🧪 Test tube
		r == 0x1F680 || // 🚀 Rocket
		r == 0x1F527 || // 🔧 Wrench
		r == 0x1F4A9 || // 💩 Pile of poo
		r == 0x1F47E || // 👾 Alien monster
		r == 0x1F47D || // 👽 Extraterrestrial alien
		r >= 0x1100 && (r <= 0x115F || r == 0x2329 || r == 0x232A ||
			(r >= 0x2E80 && r <= 0x303F) ||
			(r >= 0x3040 && r <= 0x3247 && r != 0x303F) ||
			(r >= 0x3250 && r <= 0x4DBF) ||
			(r >= 0x4E00 && r <= 0xA4C6) ||
			(r >= 0xA960 && r <= 0xA97C) ||
			(r >= 0xAC00 && r <= 0xD7A3) ||
			(r >= 0xF900 && r <= 0xFAFF) ||
			(r >= 0xFE10 && r <= 0xFE19) ||
			(r >= 0xFE30 && r <= 0xFE6F) ||
			(r >= 0xFF00 && r <= 0xFF60) ||
			(r >= 0xFFE0 && r <= 0xFFE6) ||
			(r >= 0x1F000 && r <= 0x1FAFF) || // Emoji and symbols
			(r >= 0x1F200 && r <= 0x1F2FF) || // Enclosed characters
			(r >= 0x1F300 && r <= 0x1F5FF) || // Misc symbols and pictographs
			(r >= 0x1F600 && r <= 0x1F64F) || // Emoticons
			(r >= 0x1F680 && r <= 0x1F6FF) || // Transport and map symbols
			(r >= 0x1F700 && r <= 0x1F77F) || // Alchemical symbols
			(r >= 0x1F780 && r <= 0x1F7FF) || // Geometric shapes extended
			(r >= 0x1F800 && r <= 0x1F8FF) || // Supplemental arrows-C
			(r >= 0x1F900 && r <= 0x1F9FF) || // Supplemental symbols and pictographs
			(r >= 0x1FA00 && r <= 0x1FA6F) || // Chess symbols
			(r >= 0x1FA70 && r <= 0x1FAFF) || // Symbols and pictographs extended-A
			(r >= 0x20000 && r <= 0x2FFFD) ||
			(r >= 0x30000 && r <= 0x3FFFD)):
		return 2
	// Handle emoji variation selectors
	case r == 0xFE0F: // Variation selector-16 (emoji presentation)
		return 0
	default:
		return 1
	}
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
		BorderStyle: "round", // round, double, single
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
	case "single":
		hBorder, vBorder = "─", "│"
		tlCorner, trCorner, blCorner, brCorner = "┌", "┐", "└", "┘"
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
