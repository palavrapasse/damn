package entity

import (
	"strings"
	"testing"
)

func TestCannotCreatePlatformWithContextEmpty(t *testing.T) {
	name := ""

	_, err := NewPlatform(name)

	if err == nil {
		t.Fatalf("Context designated by the string below is empty, but no error was identified\nString: %s", name)
	}
}

func TestCannotCreatePlatformWithContextWithOnlySpaces(t *testing.T) {
	name := "   "

	_, err := NewPlatform(name)

	if err == nil {
		t.Fatalf("Context designated by the string below contains only spaces, but no error was identified\nString: %s", name)
	}
}

func TestCannotCreatePlatformWithContextThatExceeds30Characters(t *testing.T) {
	name := strings.Repeat("x", 37)

	_, err := NewPlatform(name)

	if err == nil {
		t.Fatalf("Context designated by the string below exceeds 30 characters, but no error was identified\nString: %s", name)
	}
}

func TestCanCreatePlatformWithContextThatMatches30Characters(t *testing.T) {
	name := strings.Repeat("x", 30)

	_, err := NewPlatform(name)

	if err != nil {
		t.Fatalf("Context designated by the string below matches 30 characters, but an error was identified\nString: %s", name)
	}
}

func TestCanCreatePlatformWithContextThatDoesNotExceed30Characters(t *testing.T) {
	name := strings.Repeat("x", 29)

	_, err := NewPlatform(name)

	if err != nil {
		t.Fatalf("Context designated by the string below does not exceed 30 characters, but an error was identified\nString: %s", name)
	}
}

func TestCanCreatePlatformAndTrimsContextSpaces(t *testing.T) {
	name := " name    "

	leak, _ := NewPlatform(name)

	if len(leak.Name) == len(name) {
		t.Fatalf("Original name string contains unneeded spaces, and should be trimmed, but output summary still contains those spaces")
	}
}
