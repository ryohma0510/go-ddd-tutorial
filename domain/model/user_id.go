package model

type UserID struct {
	value string
}

func (u UserID) Value() string {
	return u.value
}

func NewUserID(value string) UserID {
	return UserID{value: value}
}
