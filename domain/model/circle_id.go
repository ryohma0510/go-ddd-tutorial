package model

import "errors"

type CircleID struct {
	value string
}

func (c CircleID) Value() string {
	return c.value
}

var (
	errEmptyCircleID = errors.New("empty circle id")
)

func NewCircleID(value string) (CircleID, error) {
	if value == "" {
		return CircleID{}, errEmptyCircleID
	}

	return CircleID{value: value}, nil
}
