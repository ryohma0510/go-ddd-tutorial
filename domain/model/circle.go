package model

import (
	"fmt"
)

const (
	maxMemberNum = 29
)

var (
	errTooManyMembers = fmt.Errorf("too many members. members must be <= %d", maxMemberNum)
)

type Circle struct {
	id      CircleID
	name    CircleName
	owner   User
	members CircleMembers
}

func (c Circle) Id() CircleID {
	return c.id
}

func (c Circle) Name() CircleName {
	return c.name
}

func (c Circle) Owner() User {
	return c.owner
}

func (c Circle) Members() CircleMembers {
	return c.members
}

func (c *Circle) Join(user User) error {
	if c.isFull() {
		return errTooManyMembers
	}

	c.members.Add(user)

	return nil
}

func (c Circle) isFull() bool {
	return len(c.members.collection) >= maxMemberNum
}
