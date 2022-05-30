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
	id        CircleID
	name      CircleName
	owner     User
	memberIDs []UserID
}

func (c Circle) Name() CircleName {
	return c.name
}

func (c *Circle) Join(userID UserID) error {
	if c.isFull() {
		return errTooManyMembers
	}

	c.memberIDs = append(c.memberIDs, userID)

	return nil
}

func (c Circle) isFull() bool {
	return len(c.memberIDs) >= maxMemberNum
}
