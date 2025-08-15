package Converter

import (
	"errors"
	"strconv"
	"strings"
)

func StringToUint(str string) uint {
	if isEmpty := strings.Trim(str, " "); isEmpty == "" {
		errors.New("String is null")
	}
	Uint, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		errors.New("Can not convert string to uint")
	}
	return uint(Uint)
}
