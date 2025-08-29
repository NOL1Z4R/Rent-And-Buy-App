package Converter

import (
	"errors"
	"strconv"
)

func UintToString(uint uint) string {
	if &uint == nil {
		errors.New("Uint is null")
		return ""
	}

	return strconv.Itoa(int(uint))
}
