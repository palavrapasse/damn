package query

import (
	"strings"
	"testing"
)

func TestCannotCreateEmailWithEmptyString(t *testing.T) {
	email := ""

	_, err := NewEmail(email)

	if err == nil {
		t.Fatalf("Email designated by the string below is empty, but no error was identified\nString: %s", email)
	}
}

func TestCannotCreateEmailWithStringWithOnlySpaces(t *testing.T) {
	email := "   "

	_, err := NewEmail(email)

	if err == nil {
		t.Fatalf("Email designated by the string below contains only spaces, but no error was identified\nString: %s", email)
	}
}

func TestCannotCreateEmailWithInvalidEmail(t *testing.T) {
	email := "email@"

	_, err := NewEmail(email)

	if err == nil {
		t.Fatalf("Email designated by the string below is invalid, but no error was identified\nString: %s", email)
	}
}

func TestCannotCreateEmailWithStringThatExceeds130Characters(t *testing.T) {
	email := strings.Repeat("x", 131) + "@gmail.com"

	_, err := NewEmail(email)

	if err == nil {
		t.Fatalf("Email designated by the string below exceeds 130 characters, but no error was identified\nString: %s", email)
	}
}

func TestCanCreateEmailWithStringThatMatches130Characters(t *testing.T) {
	email := strings.Repeat("x", 120) + "@gmail.com"

	_, err := NewEmail(email)

	if err != nil {
		t.Fatalf("Email designated by the string below matches 130 characters, but an error was identified\nString: %s", email)
	}
}

func TestCanCreateEmailWithStringThatDoesNotExceed130Characters(t *testing.T) {
	email := strings.Repeat("x", 119) + "@gmail.com"

	_, err := NewEmail(email)

	if err != nil {
		t.Fatalf("Email designated by the string below does not exceed 130 characters, but an error was identified\nString: %s", email)
	}
}

func TestCreateEmailTrimsSpaces(t *testing.T) {
	email := " email@gmail.com    "

	e, err := NewEmail(email)

	panicOnError(err)

	if len(e) == len(email) {
		t.Fatalf("Original email string contains unneeded spaces, and should be trimmed, but output summary still contains those spaces")
	}
}

func TestCreateEmailUsesLowercase(t *testing.T) {
	email := "eMail@gMAil.Com"
	lowerEmail := strings.ToLower(email)

	e, err := NewEmail(email)

	panicOnError(err)

	if e != Email(lowerEmail) {
		t.Fatalf("Original email contains upper case characters, but output should convert them to lower case")
	}
}
