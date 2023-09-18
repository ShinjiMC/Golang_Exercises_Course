package main

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	standars := `^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`
	result, ok := regexp.MatchString(standars, text)
	if ok != nil {
		panic("invalid regex: " + standars)
	}
	return result
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[*~=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	counts := 0
	re := regexp.MustCompile(`(?i)".*password.*"`)
	for _, i := range lines {
		if re.MatchString(i) {
			counts++
		}
	}
	return counts
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllLiteralString(text, "")
}

func TagWithUserName(lines []string) []string {
	var strings []string
	re := regexp.MustCompile(`\sUser\s+(\S+)`)
	for _, i := range lines {
		j := re.FindStringSubmatch(i)
		if 2 <= len(j) {
			i = "[USR] " + j[1] + " " + i
		}
		strings = append(strings, i)
	}
	return strings
}

func main() {
	fmt.Println("Test 1: ", IsValidLine("[ERR] A good error here"))
	fmt.Println("Test 2: ", SplitLogLine("section 1<*>section 2<~~~>section 3"))
	lines := []string{
		`[INF] passWord`,
		`"passWord"`,
		`[INF] User saw error message "Unexpected Error" on page load.`,
		`[INF] The message "Please reset your password" was ignored by the user`,
	}
	fmt.Println("Test 3: ", CountQuotedPasswords(lines))
	fmt.Println("Test 4:", RemoveEndOfLineText("[INF] end-of-line23033 Network Failure end-of-line27"))
	result := []string{
		"[WRN] User James123 has exceeded storage space.",
		"[WRN] Host down. User   Michelle4 lost connection.",
		"[INF] Users can login again after 23:00.",
		"[DBG] We need to check that user names are at least 6 chars long.",
	}
	fmt.Println("Test 4:", TagWithUserName(result))
}
