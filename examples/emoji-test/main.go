package main

import (
	"fmt"
	box "github.com/afeiship/box"
)

func main() {
	fmt.Println("Emoji Alignment Test:")
	fmt.Println("====================")

	// Test various emoji and Unicode characters
	testCases := [][]string{
		{
			"✅ Success",
			"⚠️ Warning",
			"❌ Error",
		},
		{
			"🌟 Unicode Stars 🌟",
			"Café ☕ Restaurant",
			"🎉 Party Time 🎉",
		},
		{
			"🚀 Launch: 1.2.3",
			"📦 Dependencies: OK",
			"🧪 Tests: 42/42 ✅",
		},
	}

	for i, testCase := range testCases {
		fmt.Printf("\nTest Case %d:\n", i+1)
		box.PrintBox(testCase)
	}
}