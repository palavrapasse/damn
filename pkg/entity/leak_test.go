package entity

import (
	"strings"
	"testing"
)

func TestCannotCreateLeakWithContextEmpty(t *testing.T) {
	dateInSeconds := DateInSeconds(10)
	context := ""

	_, err := NewLeak(context, dateInSeconds)

	if err == nil {
		t.Fatalf("Context designated by the string below is empty, but no error was identified\nString: %s", context)
	}
}

func TestCannotCreateLeakWithSpacesContext(t *testing.T) {
	dateInSeconds := DateInSeconds(10)
	context := "   "

	_, err := NewLeak(context, dateInSeconds)

	if err == nil {
		t.Fatalf("Context designated by the string below contains only spaces, but no error was identified\nString: %s", context)
	}
}

func TestCannotCreateLeakWithContextThatExceeds130Characters(t *testing.T) {
	dateInSeconds := DateInSeconds(10)
	context := strings.Repeat("x", 131)

	_, err := NewLeak(context, dateInSeconds)

	if err == nil {
		t.Fatalf("Context designated by the string below exceeds 130 characters, but no error was identified\nString: %s", context)
	}
}

func TestCanCreateLeakWithContextThatMatches130Characters(t *testing.T) {
	dateInSeconds := DateInSeconds(10)
	context := strings.Repeat("x", 130)

	_, err := NewLeak(context, dateInSeconds)

	if err != nil {
		t.Fatalf("Context designated by the string below matches 130 characters, but an error was identified\nString: %s", context)
	}
}

func TestCanCreateLeakWithContextThatDoesNotExceed130Characters(t *testing.T) {
	dateInSeconds := DateInSeconds(10)
	context := strings.Repeat("x", 129)

	_, err := NewLeak(context, dateInSeconds)

	if err != nil {
		t.Fatalf("Context designated by the string below does not exceed 130 characters, but an error was identified\nString: %s", context)
	}
}

func TestCanCreateLeakAndTrimsContextSpaces(t *testing.T) {
	dateInSeconds := DateInSeconds(10)
	context := " context    "

	leak, _ := NewLeak(context, dateInSeconds)

	if len(leak.Context) == len(context) {
		t.Fatalf("Original context string contains unneeded spaces, and should be trimmed, but output summary still contains those spaces")
	}
}
