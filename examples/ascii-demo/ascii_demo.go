package main

import (
	"fmt"

	"github.com/afeiship/go-box"
)

func main() {
	lines := []string{
		"✔ Detected 1 changed file:",
		" M utils/git.go (5674 chars)",
	}

	fmt.Println("🎨 ASCII 框样式测试:")
	box.PrintASCIIBox(lines)
}
