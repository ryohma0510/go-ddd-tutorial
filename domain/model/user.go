package model

type User struct {
	id UserID
}

func (u User) ID() UserID {
	return u.id
}
