package BetterConfig_test

import (
	"path"
	"testing"

	"BetterConfig"
)

func Test_GetAll(t *testing.T) {
	err := BetterConfig.InitBetterConfig(
		path.Join("Medtereus", "TestingMyConfigLibrary"),
	)
	if err != nil {
		t.Fatalf("InitBetterConfig() failed: %v", err)
	}

	// Bool
	if got := BetterConfig.GetBool(TEST_BOOL_TRUE); got != BOOL_T {
		t.Errorf("GetBool(%q) = %v, want %v",
			TEST_BOOL_TRUE, got, BOOL_T)
	}

	if got := BetterConfig.GetBool(TEST_BOOL_FALSE); got != BOOL_F {
		t.Errorf("GetBool(%q) = %v, want %v",
			TEST_BOOL_FALSE, got, BOOL_F)
	}

	// Float
	if got := BetterConfig.GetFloat(TEST_FLOAT_POS); got != FLOAT_POS {
		t.Errorf("GetFloat(%q) = %v, want %v",
			TEST_FLOAT_POS, got, FLOAT_POS)
	}

	if got := BetterConfig.GetFloat(TEST_FLOAT_NEG); got != FLOAT_NEG {
		t.Errorf("GetFloat(%q) = %v, want %v",
			TEST_FLOAT_NEG, got, FLOAT_NEG)
	}

	// String
	if got := BetterConfig.GetString(TEST_STRING); got != STRING {
		t.Errorf("GetString(%q) = %q, want %q",
			TEST_STRING, got, STRING)
	}

	// Int
	intTests := []struct {
		key  string
		want int
	}{
		{TEST_INT_POSITIVE_32, INT_POS_32},
		{TEST_INT_NEGATIVE_32, INT_NEG_32},
	}

	if ^uint(0)>>63 == 1 {
		intTests = append(intTests,
			struct {
				key  string
				want int
			}{TEST_INT_POSITIVE_64, INT_POS_64},
			struct {
				key  string
				want int
			}{TEST_INT_NEGATIVE_64, INT_NEG_64},
		)
	}

	for _, tt := range intTests {
		if got := BetterConfig.GetInt(tt.key); got != tt.want {
			t.Errorf("GetInt(%q) = %d, want %d",
				tt.key, got, tt.want)
		}
	}

	// UInt
	if got := BetterConfig.GetUInt(TEST_UINT_POSITIVE_32); got != UINT_POS_32 {
		t.Errorf("GetUInt(%q) = %d, want %d",
			TEST_UINT_POSITIVE_32, got, UINT_POS_32)
	}

	if ^uint(0)>>63 == 1 {
		if got := BetterConfig.GetUInt(TEST_UINT_POSITIVE_64); got != UINT_POS_64 {
			t.Errorf("GetUInt(%q) = %d, want %d",
				TEST_UINT_POSITIVE_64, got, UINT_POS_64)
		}
	}
}
