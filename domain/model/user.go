package model

type User struct {
	id UserID
}

func (u User) ID() UserID {
	return u.id
}

func (u User) IsPremium() bool {
	// 適当
	return true
}
