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
	packs := s.repo.GetAll()

	return s.rule.CalculatePackAllocation(requiredItems, packs)
}
