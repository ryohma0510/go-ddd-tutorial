package model

type CircleMembers struct {
	collection []User
}

func (m CircleMembers) Collection() []User {
	return m.collection
}

func (m *CircleMembers) Add(user User) {
	m.collection = append(m.collection, user)
}
