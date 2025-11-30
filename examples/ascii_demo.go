package main

import (
	"fmt"
	"strings"
)

// 从 box.go 复制必要的函数进行测试
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

func PrintASCIIBox(lines []string) {
	if len(lines) == 0 {
		return
	}

	maxLen := 0
	for _, line := range lines {
		if l := StringWidth(line); l > maxLen {
			maxLen = l
		}
	}

	top := "+" + strings.Repeat("-", maxLen+2) + "+"
	fmt.Println(top)

	for _, line := range lines {
		width := StringWidth(line)
		padding := strings.Repeat(" ", maxLen-width)
		fmt.Printf("| %s%s |\n", line, padding)
	}

	bottom := "+" + strings.Repeat("-", maxLen+2) + "+"
	fmt.Println(bottom)
}

func main() {
	lines := []string{
		"✔ Detected 1 changed file:",
		" M utils/git.go (5674 chars)",
	}

	fmt.Println("🎨 ASCII 框样式测试:")
	PrintASCIIBox(lines)
}