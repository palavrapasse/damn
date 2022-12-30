package entity

import (
	"strings"
	"testing"
)

func TestCannotCreateEmptyPassword(t *testing.T) {
	password := ""

	_, err := NewPassword(password)

	if err == nil {
		t.Fatalf("Password designated by the string below is empty, but no error was identified\nString: %s", password)
	}
}

func TestCanCreatePassword(t *testing.T) {
	password := strings.Repeat("x", 30)

	_, err := NewPassword(password)

	if err != nil {
		t.Fatalf("Password designated by the string below matches 30 characters, but an error was identified\nString: %s", password)
	}
}

func TestCanCreatePasswordAndTrimsSpaces(t *testing.T) {
	password := " password    "

	p, _ := NewPassword(password)

	if len(p) == len(password) {
		t.Fatalf("Original password string contains unneeded spaces, and should be trimmed, but output password still contains those spaces")
	}
}
