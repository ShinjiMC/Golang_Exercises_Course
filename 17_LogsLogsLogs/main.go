package main

import (
	"fmt"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	appChars := map[rune]string{
		'❗': "recommendation",
		'🔍': "search",
		'☀': "weather",
	}

	for _, char := range log {
		if appName, found := appChars[char]; found {
			return appName
		}
	}
	return "default"

}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	runes := []rune(log)
	for index, char := range runes {
		if char == oldRune {
			runes[index] = newRune
		}
	}
	return string(runes)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}

func main() {
	fmt.Println("Test 1: ", Application("❗ recommended search product 🔍"))
	log := "please replace '👎' with '👍'"
	fmt.Println("Test 2: ", Replace(log, '👎', '👍'))
	fmt.Println("Test 3: ", WithinLimit("hello❗", 6))
}
