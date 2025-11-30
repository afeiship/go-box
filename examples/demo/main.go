package main

import (
	"fmt"
	"time"
	box "github.com/afeiship/go-box"
)

func main() {
	fmt.Println("Box Package Demo")
	fmt.Println("================")

	// Demo 1: Status messages
	fmt.Println("\n📊 Status Messages:")
	statusMessages := []string{
		"✅ Build successful",
		"🧪 Tests passed: 42/42",
		"📦 Dependencies up to date",
	}
	box.PrintBoxWithOptions(statusMessages, &box.BoxOptions{
		BorderStyle: "round",
		Padding:     0,
	})

	time.Sleep(1 * time.Second)

	// Demo 2: Error messages
	fmt.Println("\n❌ Error Messages:")
	errorMessages := []string{
		"\x1b[31m🚨 Critical Error\x1b[0m",
		"Database connection failed",
		"Please check your configuration",
	}
	box.PrintBoxWithOptions(errorMessages, &box.BoxOptions{
		BorderStyle: "single",
		Padding:     1,
		Indent:      2,
	})

	time.Sleep(1 * time.Second)

	// Demo 3: Welcome message
	fmt.Println("\n🎉 Welcome Message:")
	welcomeMessage := []string{
		"\x1b[1m\x1b[36mWelcome to Go Box!\x1b[0m",
		"Create beautiful ASCII boxes",
		"With customizable styles",
		"Supports colors and Unicode 🌟",
	}
	box.PrintBoxWithOptions(welcomeMessage, &box.BoxOptions{
		BorderStyle: "double",
		Padding:     1,
	})

	time.Sleep(1 * time.Second)

	// Demo 4: Menu options
	fmt.Println("\n📋 Menu Options:")
	menuOptions := []string{
		"[1] Create new box",
		"[2] Customize style",
		"[3] Add colors",
		"[4] Exit",
	}
	box.PrintBoxWithOptions(menuOptions, &box.BoxOptions{
		BorderStyle: "round",
		Indent:      4,
	})

	// Demo 5: Code example
	fmt.Println("\n💻 Code Example:")
	codeExample := []string{
		"lines := []string{",
		"  \"Hello, World!\",",
		"  \"This is a box\",",
		"}",
		"box.PrintBox(lines)",
	}
	box.PrintBoxWithOptions(codeExample, &box.BoxOptions{
		BorderStyle: "single",
		Indent:      2,
	})
}