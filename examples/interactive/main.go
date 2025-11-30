package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	box "github.com/afeiship/go-box"
)

func main() {
	fmt.Println("Interactive Box Example:")
	fmt.Println("=======================")
	fmt.Println("Enter lines of text (press Enter on empty line to finish):")

	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for {
		fmt.Print(">")
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			break
		}
		lines = append(lines, line)
	}

	if len(lines) > 0 {
		fmt.Println("\nYour box:")
		fmt.Println("=========")
		box.PrintBox(lines)

		// Show different styles
		fmt.Println("\nWith different styles:")
		fmt.Println("======================")

		styles := []string{"round", "single", "double"}
		caser := cases.Title(language.English)
		for _, style := range styles {
			fmt.Printf("\n%s style:\n", caser.String(style))
			opts := &box.BoxOptions{
				BorderStyle: style,
				Padding:     0,
			}
			box.PrintBoxWithOptions(lines, opts)
		}
	} else {
		fmt.Println("No input provided!")
	}
}