package poet

import "testing"

func Test_FieldName_ShouldNotChangeNonReservedKeywords(t *testing.T) {
	cases := []string{
		"hello",
		"world",
		"iff",
		"not_reserved",
		"class_",
	}

	for _, key := range cases {
		t.Run(key, func(t *testing.T) {
			got := FieldName(key)
			if got != key {
				t.Errorf("FieldName(%q) = %q, want %q", key, got, key)
			}
		})
	}
}

func Test_FieldName_ShouldUdateReservedKeywords(t *testing.T) {
	cases := []string{
		"if",
		"class",
	}

	for _, key := range cases {
		t.Run(key, func(t *testing.T) {
			if !IsReserved(key) {
				t.Errorf("%s should be reserved", key)
			}
			got := FieldName(key)
			if IsReserved(got) {
				t.Errorf("FieldName(%s) = %s, should not be reserved", key, got)
			}
		})
	}
}
