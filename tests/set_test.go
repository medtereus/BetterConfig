package BetterConfig_test

import (
	"BetterConfig"
	"path"
	"testing"
)

const (
	TEST_BOOL_TRUE  string = "TEST_SETBOOL_TRUE"
	BOOL_T          bool   = true
	TEST_BOOL_FALSE string = "TEST_SETBOOL_FALSE"
	BOOL_F          bool   = false

	TEST_FLOAT_POS string  = "TEST_FLOAT_POS"
	FLOAT_POS      float64 = 3.141592
	TEST_FLOAT_NEG string  = "TEST_FLOAT_NEG"
	FLOAT_NEG      float64 = -3.141592
	TEST_STRING    string  = "TEST_STRING"
	STRING         string  = "SomeText String \n Некая строка \t etwas möbel äüö KEK"

	TEST_INT_POSITIVE_32 string = "TEST_INT_POSITIVE_32"
	INT_POS_32           int    = 2147483647
	TEST_INT_POSITIVE_64 string = "TEST_INT_POSITIVE_64"
	INT_POS_64           int    = 9223372036854775807
	TEST_INT_NEGATIVE_32 string = "TEST_INT_NEGATIVE_32"
	INT_NEG_32           int    = -2147483648
	TEST_INT_NEGATIVE_64 string = "TEST_INT_NEGATIVE_64"
	INT_NEG_64           int    = -9223372036854775808

	TEST_UINT_POSITIVE_32 string = "TEST_UINT_POSITIVE_32"
	UINT_POS_32           uint   = 4294967295
	TEST_UINT_POSITIVE_64 string = "TEST_UINT_POSITIVE_64"
	UINT_POS_64           uint   = 18446744073709551615
)

func Test_SetAll(t *testing.T) {
	err := BetterConfig.InitBetterConfig(
		path.Join("Medtereus", "TestingMyConfigLibrary"),
	)
	if err != nil {
		t.Fatalf("InitBetterConfig() failed: %v", err)
	}

	// Bool
	BetterConfig.SetBool(TEST_BOOL_TRUE, BOOL_T)
	BetterConfig.SetBool(TEST_BOOL_FALSE, BOOL_F)

	// Float
	BetterConfig.SetFloat(TEST_FLOAT_POS, FLOAT_POS)
	BetterConfig.SetFloat(TEST_FLOAT_NEG, FLOAT_NEG)

	// String
	BetterConfig.SetString(TEST_STRING, STRING)

	// Int
	BetterConfig.SetInt(TEST_INT_POSITIVE_32, INT_POS_32)
	BetterConfig.SetInt(TEST_INT_POSITIVE_64, INT_POS_64)
	BetterConfig.SetInt(TEST_INT_NEGATIVE_32, INT_NEG_32)
	BetterConfig.SetInt(TEST_INT_NEGATIVE_64, INT_NEG_64)

	// UInt
	BetterConfig.SetUInt(TEST_UINT_POSITIVE_32, UINT_POS_32)
	BetterConfig.SetUInt(TEST_UINT_POSITIVE_64, UINT_POS_64)
}
