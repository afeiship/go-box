package main

import (
	"fmt"
	box "github.com/afeiship/box"
)

func main() {
	fmt.Println("Character Width Debug:")
	fmt.Println("=======================")

	testStrings := []string{
		"✅ Build successful",
		"🧪 Tests passed: 42/42",
		"📦 Dependencies up to date",
	}

	for _, str := range testStrings {
		width := box.StringWidth(str)
		fmt.Printf("String: '%s'\n", str)
		fmt.Printf("Length: %d, Calculated Width: %d\n", len(str), width)
		fmt.Printf("Rune by rune:\n")
		for i, r := range str {
			runeWidth := box.GetRuneWidth(r)
			fmt.Printf("  [%d] %c (U+%04X) - Width: %d\n", i, r, r, runeWidth)
		}
		fmt.Println()
	}
}