package validate

import (
	"testing"
)

func TestValidadeSequence(t *testing.T) {

	validator := CreateValidator(ValidatorAssync)
	cases := []struct {
		should     bool
		errMessage string
		scenario   []string
	}{
		{
			errMessage: "Should return true because there is a vertical sequence",
			should:     true,
			scenario: []string{
				"DUHHDB",
				"DUBUHD",
				"UBUUHU",
				"BHBDHH",
				"UDBDUH",
				"UHDDDD"},
		},
		{
			errMessage: "Should return true because there is a horizontal sequence",
			should:     true,
			scenario: []string{
				"DUHHDB",
				"DUBUHD",
				"UBUUHH",
				"BHBDHH",
				"UDBDUH",
				"UHDDHH"},
		},
		{
			errMessage: "Should return true because there is a left to right diagonal sequence",
			should:     true,
			scenario: []string{
				"DUHHDB",
				"DDBUHD",
				"UBDUHU",
				"BUBDHH",
				"UDHDUH",
				"UHDUHD"},
		},
		{
			errMessage: "Should return true because there is a right to left diagonal sequence",
			should:     true,
			scenario: []string{
				"DUHHDB",
				"DUBUBD",
				"UBUBHH",
				"BHBDDH",
				"UDBHUH",
				"UHHDHD"},
		},

		{
			errMessage: "Should return false because there is no valid sequence ",
			should:     false,
			scenario: []string{
				"DUHHDB",
				"DUBUHD",
				"UBUUHU",
				"BHBDHH",
				"UDBDUH",
				"UHDDHD"},
		},
	}

	for _, cas := range cases {
		hasSequence := validator.ValidateSequence(cas.scenario)

		if hasSequence != cas.should {
			t.Error(cas.errMessage)
		}
	}

}

func TestCheckHorizontal(t *testing.T) {
	str := "UBYBYBDDDD"

	hasSequence := checkHorizontal(str)

	if !hasSequence {
		t.Error("Expect true, got - ", hasSequence)
	}
}

func TestCheckVertica(t *testing.T) {
	str := []string{"UBCDDDD", "UBCDDDD", "BBCDDDD", "UBCDDDD", "UBCDDDD", "UBCDDDD", "UBCDDDD"}

	hasSequence := checkVertical(str, 0)

	if !hasSequence {
		t.Error("Expect true, got - ", hasSequence)
	}
}

func TestCheckDiagonalLeftToRight(t *testing.T) {
	str := []string{
		"DUHBHBD",
		"DUBUHDB",
		"UDUUHBH",
		"BUBDBHU",
		"DDUBUBB",
		"UDBUUHH",
		"UDBUUHH"}

	hasSequence := checkDiagonalLeftToRight(str)

	if !hasSequence {
		t.Error("Expect false, got - ", hasSequence)
	}
}

func TestCheckDiagonalRightToLeft(t *testing.T) {
	str := []string{
		"DUHBHBD",
		"DUBUHDB",
		"UDDUHBH",
		"BUBDBHU",
		"DUHBUBB",
		"UDBUUHH",
		"UDBBUHH"}

	hasSequence := checkDiagonalRightToLeft(str)

	if !hasSequence {
		t.Error("Expect false, got - ", hasSequence)
	}
}
