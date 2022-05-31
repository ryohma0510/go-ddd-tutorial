package spec

import (
	"go-ddd-tutorial/domain/model"
	"go-ddd-tutorial/domain/repository"
)

type CircleFullSpec struct {
	userRepository repository.IUserRepository
}

func NewCircleFullSpec(userRepository repository.IUserRepository) *CircleFullSpec {
	return &CircleFullSpec{userRepository: userRepository}
}

func (spec CircleFullSpec) IsSatisfiedByCircle(circle model.Circle) bool {
	members, err := spec.userRepository.FindByIDs(circle.MemberIDs())
	if err != nil {
		return false
	}

	var premiumMemberLen int
	for _, member := range members {
		if member.IsPremium() {
			premiumMemberLen++
		}
	}

	const maxPremiumMemberLen = 10
	var circleUpperLimit int
	if premiumMemberLen < maxPremiumMemberLen {
		circleUpperLimit = 30
	} else {
		circleUpperLimit = 50
	}

	return len(members) >= circleUpperLimit
}
