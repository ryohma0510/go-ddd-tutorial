package model

type ICircleDTO interface {
	ID() CircleID
	Name() CircleName
	Owner() User
	MemberIDs() []UserID
}

type CircleDTO struct {
	id        CircleID
	name      CircleName
	owner     User
	memberIDs []UserID
}

func (dto CircleDTO) ID() CircleID {
	return dto.id
}

func (dto CircleDTO) Name() CircleName {
	return dto.name
}

func (dto CircleDTO) Owner() User {
	return dto.owner
}

func (dto CircleDTO) MemberIDs() []UserID {
	return dto.memberIDs
}
