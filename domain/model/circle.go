package model

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
