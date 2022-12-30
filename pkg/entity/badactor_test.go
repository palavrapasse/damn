package entity

import (
	"strings"
	"testing"
)

func TestCannotCreateBadActorWithEmptyIdentifier(t *testing.T) {
	identifier := ""

	_, err := NewBadActor(identifier)

	if err == nil {
		t.Fatalf("Identifier designated by the string below is empty, but no error was identified\nString: %s", identifier)
	}
}

func TestCannotCreateBadActorWithSpacesIdentifier(t *testing.T) {
	identifier := "   "

	_, err := NewBadActor(identifier)

	if err == nil {
		t.Fatalf("Identifier designated by the string below contains only spaces, but no error was identified\nString: %s", identifier)
	}
}

func TestCannotCreateBadActorWithIdentifierThatExceeds30Characters(t *testing.T) {
	identifier := strings.Repeat("x", 31)

	_, err := NewBadActor(identifier)

	if err == nil {
		t.Fatalf("Identifier designated by the string below exceeds 30 characters, but no error was identified\nString: %s", identifier)
	}
}

func TestCanCreateBadActorWithIdentifierThatMatches30Characters(t *testing.T) {
	identifier := strings.Repeat("x", 30)

	_, err := NewBadActor(identifier)

	if err != nil {
		t.Fatalf("Identifier designated by the string below matches 30 characters, but an error was identified\nString: %s", identifier)
	}
}

func TestCanCreateBadActorWithIdentifierThatDoesNotExceed30Characters(t *testing.T) {
	identifier := strings.Repeat("x", 20)

	_, err := NewBadActor(identifier)

	if err != nil {
		t.Fatalf("Identifier designated by the string below does not exceed 30 characters, but an error was identified\nString: %s", identifier)
	}
}

func TestCanCreateBadActorAndTrimsIdentifierSpaces(t *testing.T) {
	identifier := " identifier    "

	ba, _ := NewBadActor(identifier)

	if len(ba.Identifier) == len(identifier) {
		t.Fatalf("Original identifier string contains unneeded spaces, and should be trimmed, but output summary still contains those spaces")
	}
}
