package main

import "testing"

// you must import the testing package to preform testing scripts.
var tests = []struct {
	name     string
	divided  float32
	divisor  float32
	expected float32
	ieErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0, true},
	{"expect-5", 50.0, 10.0, 5.0, true},
	{"expect-fraction", -1.0, -777.0, 0.0012870013, true},
} // This process fills in the data to test

func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("Got an error", err)
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 0)
	if err != nil {
		t.Error("No errors", err)
	}
}
func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.divided, tt.divisor)
		if tt.ieErr {
			if err == nil {
				t.Error("No error thrown")
			}
		} else {
			if err != nil {
				t.Error("Got an Error, Unexpected", err.Error())
			}
		}
		if got != tt.expected {
			t.Errorf("Expected %f but got %f", tt.expected, got)
		}
	}
}

// go test -coverprofile=coverage.out && go tool cover -html=coverage.out
// Produces an HTML error report of working and non-working functions
