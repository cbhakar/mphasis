package utils

import "testing"

func TestStrInListStatusValid(t *testing.T) {
	if ok := StrInListStatus("test_string", []string{"test_string", "another_string"}); !ok {
		t.Error("not an expected output")
		t.Fail()
	}
}
func TestStrInListStatusInValid(t *testing.T) {
	if ok := StrInListStatus("no_string", []string{"test_string", "another_string"}); ok {
		t.Error("not an expected output")
		t.Fail()
	}
}

func TestCheckIntValueValid(t *testing.T) {
	if intValue, ok := CheckIntValue("45"); !ok || intValue != 45 {
		t.Error("not an expected output")
		t.Fail()
	}
}
func TestCheckIntValueInValid(t *testing.T) {
	if intValue, ok := CheckIntValue("test"); ok || intValue != 0 {
		t.Error("not an expected output")
		t.Fail()
	}
}
