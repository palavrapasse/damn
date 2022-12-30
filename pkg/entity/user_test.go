package entity

import (
	"strings"
	"testing"
)

func TestCannotCreateUserWithEmailEmpty(t *testing.T) {
	email := ""

	_, err := NewUser(email)

	if err == nil {
		t.Fatalf("Email designated by the string below is empty, but no error was identified\nString: %s", email)
	}
}

func TestCannotCreateUserWithEmailWithOnlySpaces(t *testing.T) {
	email := "   "

	_, err := NewUser(email)

	if err == nil {
		t.Fatalf("Email designated by the string below contains only spaces, but no error was identified\nString: %s", email)
	}
}

func TestCannotCreateUserWithInvalidEmail(t *testing.T) {
	email := "email@"

	_, err := NewUser(email)

	if err == nil {
		t.Fatalf("Email designated by the string below is invalid, but no error was identified\nString: %s", email)
	}
}

func TestCannotCreateUserWithEmailThatExceeds130Characters(t *testing.T) {
	email := strings.Repeat("x", 131) + "@gmail.com"

	_, err := NewUser(email)

	if err == nil {
		t.Fatalf("Email designated by the string below exceeds 130 characters, but no error was identified\nString: %s", email)
	}
}

func TestCanCreateUserWithEmailThatMatches130Characters(t *testing.T) {
	email := strings.Repeat("x", 120) + "@gmail.com"

	_, err := NewUser(email)

	if err != nil {
		t.Fatalf("Email designated by the string below matches 130 characters, but an error was identified\nString: %s", email)
	}
}

func TestCanCreateUserWithEmailThatDoesNotExceed130Characters(t *testing.T) {
	email := strings.Repeat("x", 119) + "@gmail.com"

	_, err := NewUser(email)

	if err != nil {
		t.Fatalf("Email designated by the string below does not exceed 130 characters, but an error was identified\nString: %s", email)
	}
}

func TestCanCreateUserAndTrimsEmailSpaces(t *testing.T) {
	email := " email@gmail.com    "

	leak, _ := NewUser(email)

	if len(leak.Email) == len(email) {
		t.Fatalf("Original email string contains unneeded spaces, and should be trimmed, but output summary still contains those spaces")
	}
}
