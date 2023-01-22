package types

import (
	"strconv"
)

// bool
type boolType struct {
}

func (u boolType) Equals(a, b interface{}) bool {
	return a == b
}

func (u boolType) Pack(value interface{}) (interface{}, error) {
	return value.(bool), nil
}

func (u boolType) UnPack(value interface{}) (interface{}, error) {
	return value.(bool), nil
}

func (u boolType) Default() any {
	return false
}

func (u boolType) Pointer(required bool) any {
	if required {
		return new(bool)
	} else {
		return new(bool)
	}
}

func (u boolType) String(val any) string {
	return strconv.FormatBool(u.typed(val))
}

func (u boolType) typed(val any) bool {
	return val.(bool)
}

func (u boolType) IsEmpty(value any) bool {
	return value == nil
}

func (u boolType) ValidateValue(value any) error {
	return canCast[bool]("bool", value)
}
