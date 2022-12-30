package entity

import (
	"testing"
)

func TestCannotCreateDateInSecondsWithEmptyDate(t *testing.T) {
	date := ""

	_, err := NewDateInSeconds(date)

	if err == nil {
		t.Fatalf("Date designated by the string below is empty, but no error was identified\nString: %s", date)
	}
}

func TestCannotCreateDateInSecondsWithSpacesDate(t *testing.T) {
	date := "   "

	_, err := NewDateInSeconds(date)

	if err == nil {
		t.Fatalf("Date designated by the string below contains only spaces, but no error was identified\nString: %s", date)
	}
}

func TestCannotCreateDateInSecondsWithInvadidDateFormat(t *testing.T) {
	date := "Mon Jan _2 15:04:05 2006"

	_, err := NewDateInSeconds(date)

	if err == nil {
		t.Fatalf("Date designated by the string below is in invalid format, but no error was identified\nString: %s", date)
	}
}

func TestCanCreateDateInSeconds(t *testing.T) {
	date := "2022-12-02"

	_, err := NewDateInSeconds(date)

	if err != nil {
		t.Fatalf("Date designated by the string below is in the valid format, but an error was identified\nString: %s", date)
	}
}

func TestCanReturnDateInSecondsString(t *testing.T) {
	date := "2022-12-02"
	ds, _ := NewDateInSeconds(date)

	dsString := ds.String()

	if dsString != date {
		t.Fatalf("The date string should be the same as the one provided\nDate String: %s\nDate provided: %s", dsString, date)
	}
}
