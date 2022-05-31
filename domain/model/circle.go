package model

import (
	"time"
)

type Circle struct {
	id        CircleID
	name      CircleName
	owner     User
	memberIDs []UserID
}

func (c Circle) MemberIDs() []UserID {
	return c.memberIDs
}

func (c Circle) Name() CircleName {
	return c.name
}

func (c *Circle) Join(userID UserID) error {
	c.memberIDs = append(c.memberIDs, userID)

	return nil
}

// CreatedAt is dummy function
func (c Circle) CreatedAt() time.Time {
	return time.Now()
}
