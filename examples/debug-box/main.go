package main

import (
	"fmt"
	"strings"

	box "github.com/afeiship/box"
)

func main() {
	fmt.Println("Box Debug:")
	fmt.Println("===========")

	// Test Case 1
	fmt.Println("Test Case 1:")
	testCase1 := []string{
		"✅ Success",
		"⚠️ Warning",
		"❌ Error",
	}

	// Calculate width manually
	fmt.Printf("Manual width calculation:\n")
	maxLen := 0
	for _, line := range testCase1 {
		width := box.StringWidth(line)
		fmt.Printf("'%s' -> width: %d\n", line, width)
		if width > maxLen {
			maxLen = width
		}
	}
	fmt.Printf("Max width: %d\n", maxLen)

	// Show manual padding calculation
	for _, line := range testCase1 {
		width := box.StringWidth(line)
		padding := strings.Repeat(" ", maxLen-width)
		fmt.Printf("'%s' + '%d spaces' -> '%s%s'\n", line, maxLen-width, line, padding)
	}

	fmt.Println("\nActual box output:")
	box.PrintBox(testCase1)

	// Test Case 2
	fmt.Println("\n\nTest Case 2:")
	testCase2 := []string{
		"🌟 Unicode Stars 🌟",
		"Café ☕ Restaurant",
		"🎉 Party Time 🎉",
	}

	fmt.Printf("Manual width calculation:\n")
	maxLen = 0
	for _, line := range testCase2 {
		width := box.StringWidth(line)
		fmt.Printf("'%s' -> width: %d\n", line, width)
		if width > maxLen {
			maxLen = width
		}
	}
	fmt.Printf("Max width: %d\n", maxLen)

	fmt.Println("\nActual box output:")
	box.PrintBox(testCase2)
}
