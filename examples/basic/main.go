package main

import (
	"fmt"
	box "github.com/afeiship/go-box"
)

func main() {
	fmt.Println("Basic Box Example:")
	fmt.Println("=================")

	// Simple text box
	lines := []string{
		"Hello, World!",
		"This is a simple box",
		"with multiple lines",
	}
	box.PrintBox(lines)

	fmt.Println("\nBox with colored text:")
	fmt.Println("=====================")

	// Box with ANSI color codes
	coloredLines := []string{
		"\x1b[31mRed text\x1b[0m",
		"\x1b[32mGreen text\x1b[0m",
		"\x1b[34mBlue text\x1b[0m",
		"Normal text",
	}
	box.PrintBox(coloredLines)

	fmt.Println("\nBox with Unicode characters:")
	fmt.Println("============================")

	// Box with Unicode characters
	unicodeLines := []string{
		"🌟 Unicode Stars 🌟",
		"Café ☕ Restaurant",
		"🎉 Party Time 🎉",
	}
	box.PrintBox(unicodeLines)
}