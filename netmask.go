// Package netmask returns the shorthand netmask from an IPv4 netmask string.
package netmask

import (
	"errors"
	"strconv"
	"strings"
)

func countBits(value uint32) (count int, err error) {
	// reading right to left, error when find 0 after find 1
	var foundOne bool
	for value > 0 {
		rem := value & 1
		value = value >> 1
		if rem == 0 && foundOne {
			return 0, errors.New("set bit follows unset bit")
		}
		if rem == 1 {
			foundOne = true
			count++
		}
	}
	return
}

// ConvertNetmaskToCIDR returns CIDR notation from an IPv4 netmask string.
func ConvertNetmaskToCIDR(input string) (num int, err error) {
	var octets = strings.Split(input, ".")
	if len(octets) != 4 {
		return 0, errors.New("not enough octets")
	}
	var accum uint32
	for _, octet := range octets {
		value, err := strconv.ParseInt(octet, 10, 16)
		if err != nil {
			return 0, err
		}
		if value < 0 || value > 255 {
			return 0, errors.New("number out of range")
		}
		accum = accum << 8
		accum += uint32(value)
	}
	return countBits(accum)
}
