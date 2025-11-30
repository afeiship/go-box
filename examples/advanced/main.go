package main

import (
	"fmt"
	box "github.com/afeiship/go-box"
)

func main() {
	fmt.Println("Advanced Box Examples:")
	fmt.Println("=====================")

	// Example 1: Box with different border styles
	fmt.Println("\n1. Double Border Style:")
	fmt.Println("=======================")
	doubleBoxOpts := &box.BoxOptions{
		BorderStyle: "double",
	}
	box.PrintBoxWithOptions([]string{
		"Double Border Box",
		"Looks more formal",
		"Good for important messages",
	}, doubleBoxOpts)

	// Example 2: Box with padding
	fmt.Println("\n2. Box with Padding:")
	fmt.Println("===================")
	paddingOpts := &box.BoxOptions{
		Padding:     1,
		BorderStyle: "round",
	}
	box.PrintBoxWithOptions([]string{
		"Box with padding",
		"More breathing room",
	}, paddingOpts)

	// Example 3: Box with indentation
	fmt.Println("\n3. Box with Indentation:")
	fmt.Println("========================")
	indentOpts := &box.BoxOptions{
		Indent:      4,
		BorderStyle: "single",
	}
	box.PrintBoxWithOptions([]string{
		"Indented box",
		"Nested content",
	}, indentOpts)

	// Example 4: Box with all options
	fmt.Println("\n4. Box with All Options:")
	fmt.Println("========================")
	allOpts := &box.BoxOptions{
		Padding:     1,
		BorderStyle: "double",
		Indent:      2,
	}
	box.PrintBoxWithOptions([]string{
		"\x1b[1mStyled Box\x1b[0m",
		"With \x1b[33mpadding\x1b[0m, \x1b[32mindentation\x1b[0m,",
		"And \x1b[31mdouble borders\x1b[0m",
	}, allOpts)
}