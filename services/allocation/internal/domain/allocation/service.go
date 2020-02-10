package allocation

import (
	"github.com/gileslloyd/gs-allocation-service/internal/domain/pack"
)

type Service struct {
	repo pack.Repository
	rule Rule
}

func NewAllocationService(repository pack.Repository, rule Rule) Service {
	return Service{
		repo: repository,
		rule: rule,
	}
}

func (s Service) GetPackAllocation(requiredItems int) map[int]int {
	return s.rule.CalculatePackAllocation(requiredItems, s.getPackSizes())
}

func (s Service) getPackSizes() []int {
	packSizes := []int{}

	for _, p := range s.repo.GetAll() {
		packSizes = append(packSizes, p.GetSize())
	}

	return packSizes
}
