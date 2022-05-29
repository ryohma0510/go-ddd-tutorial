package model

import (
	"errors"
	"unicode/utf8"
)

type CircleName struct {
	value string
}

const (
	minCircleNameLen = 3
	maxCircleNameLen = 20
)

var (
	errCircleNameTooShort = errors.New("circle name too long. It must be >= 3")
	errCircleNameTooLong  = errors.New("circle name too long. It must be <= 20")
)

func (cn CircleName) Value() string {
	return cn.value
}

func NewCircleName(value string) (CircleName, error) {
	valueLen := utf8.RuneCountInString(value)

	if valueLen < minCircleNameLen {
		return CircleName{}, errCircleNameTooShort
	}

	if valueLen > maxCircleNameLen {
		return CircleName{}, errCircleNameTooLong
	}

	return CircleName{value: value}, nil
}

func (cn CircleName) Equals(other CircleName) bool {
	return cn.value == other.value
}
