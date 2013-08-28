package netmask

import (
	"testing"
)

func TestCountBitsReturnsErrorIfOneAfterZeroBit(t *testing.T) {
	var inputs = map[uint32]string{
		65: "set bit follows unset bit",
	}
	for input, emsg := range inputs {
		_, err := countBits(input)
		if err == nil || err.Error() != emsg {
			t.Errorf("case: %v, expected error: %v, but got: %v",
				input, emsg, err)
		}
	}
}

func TestCountBitsReturnsNumberOfSetBits(t *testing.T) {
	var inputs = map[uint32]int {
		4294967295	: 32,	// 11111111 11111111 11111111 11111111
		16777215	: 24,	// 11111111 11111111 11111111 00000000
		65535		: 16,	// 11111111 11111111 00000000 00000000
		255		: 8,	// 1111 1111
		254		: 7,	// 1111 1110
		252		: 6,	// 1111 1100
		240		: 4,	// 1111 0000
		224		: 3,	// 1110 0000
		192		: 2,	// 1100 0000
		128		: 1,	// 1000 0000
	}
	for input, expected := range inputs {
		actual, err := countBits(input)
		if err != nil {
			t.Errorf("case: %v, expected: %v, actual error: %v",
				input, expected, err)
		}
		if actual != expected {
			t.Errorf("case: %v, expected: %v, actual: %v",
				input, expected, actual)
		}
	}
}

func TestConvertNetmaskToCIDRReturnsErrorWhenSetBitFollowsUnsetBit(t *testing.T) {
	var inputs = map[string]string{
		"255.255.255.255.0"	: "too many or too few octets",
		"255.255"		: "too many or too few octets",
		"2550.255.0.0"		: "number out of range",
		"255.-255.0.0"		: "number out of range",
		"255.255..0"		: "strconv.ParseInt: parsing \"\": invalid syntax",
		"255a.255.0.0"		: "strconv.ParseInt: parsing \"255a\": invalid syntax",
		"255.0.255.0"		: "set bit follows unset bit",
		"240.255.0.0"		: "set bit follows unset bit",
	}

	for input, emsg := range inputs {
		_, err := ConvertNetmaskToCIDR(input)
		if err == nil || err.Error() != emsg {
			t.Errorf("case: %v, expected error: %v, but got: %v", input, emsg, err)
		}
	}
}

func TestConvertNetmaskToCIDRReturnsNumberOfSetBits(t *testing.T) {
	var inputs = map[string]int {
		"255.0.0.0"		: 8,		// 1111 1111
		"254.0.0.0"		: 7,		// 1111 1110
		"252.0.0.0"		: 6,		// 1111 1100
		"240.0.0.0"		: 4,		// 1111 0000
		"224.0.0.0"		: 3,		// 1110 0000
		"192.0.0.0"		: 2,		// 1100 0000
		"128.0.0.0"		: 1,		// 1000 0000

		"255.255.255.224"	: 27,
		"255.255.255.192"	: 26,
		"255.255.255.128"	: 25,
		"255.255.255.0"		: 24,
		"255.255.0.0"		: 16,
		"255.240.0.0"		: 12,
	}
	for input, expected := range inputs {
		actual, err := ConvertNetmaskToCIDR(input)
		if err != nil {
			t.Errorf("case: %v, expected: %v, actual error: %v",
				input, expected, err)
		}
		if actual != expected {
			t.Errorf("case: %v, expected: %v, actual: %v",
				input, expected, actual)
		}
	}
}
