package conv

import (
	"strings"
	"testing"
)

func cmpSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TestIntsCorrectlyConvertsStringValues(t *testing.T) {
	input := []string{"2", "4", "6", "3", "5", "7"}
	expected := []int{2, 4, 6, 3, 5, 7}
	actual, err := Ints(input...)

	if err != nil {
		t.Error(err)
	}

	if !cmpSlices(actual, expected) {
		t.Errorf("actual differs from expected\ngot: %+v\nexpected: %+v",
			actual, expected)
	}
}

func TestIntsCorrectlyAllocates0CapSlice(t *testing.T) {
	numbers, err := Ints()

	if err != nil {
		t.Error(err)
	}

	if cap(numbers) != 0 && len(numbers) != 0 {
		t.Errorf("resulting slice should be empty\ngot: %+v", numbers)
	}
}

func TestIntsFailsWithNoIntegerStringValue(t *testing.T) {
	input := []string{"2", "3", "a"}
	actual, err := Ints(input...)

	if err == nil {
		t.Errorf("err should not be nil\noutput: %+v", actual)
	}

	if !strings.Contains(err.Error(), "at index `2`") {
		t.Errorf("function failed unexpectedly\ngot: %v", err)
	}
}
