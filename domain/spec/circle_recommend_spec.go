package spec

import (
	"go-ddd-tutorial/domain/model"
	"time"
)

type CircleRecommendSpec struct {
	executionTime time.Time
}

func NewCircleRecommendSpec(executionTime time.Time) CircleRecommendSpec {
	return CircleRecommendSpec{executionTime: executionTime}
}

func (spec CircleRecommendSpec) IsSatisfiedByCircle(circle model.Circle) bool {
	const recommendMemberCountLowerLimit = 10
	if len(circle.MemberIDs()) < recommendMemberCountLowerLimit {
		return false
	}

	if circle.CreatedAt().Before(spec.executionTime.Add(-24 * 30 * time.Hour)) {
		return false
	}

	return true
}
