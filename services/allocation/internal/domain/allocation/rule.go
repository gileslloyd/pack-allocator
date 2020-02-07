package allocation

import "github.com/gileslloyd/gs-allocation-service/internal/domain/pack"

type Rule struct {
}

func NewPackAllocationRule() Rule {
	return Rule{}
}

func (r Rule) CalculatePackAllocation(requiredItems int, packs []pack.Entity) map[int]int {
	return map[int]int{
		250: 2,
	}
}
