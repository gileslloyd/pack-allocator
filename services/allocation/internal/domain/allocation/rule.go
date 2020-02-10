package allocation

import (
	"sort"
)

type Rule struct {
	totalItems int
}

func NewPackAllocationRule() Rule {
	return Rule{}
}

func (r Rule) CalculatePackAllocation(requiredItems int, packSizes []int) map[int]int {
	packs := map[int]int{}
	r.totalItems = 0
	sort.Ints(packSizes)

	for r.totalItems < requiredItems {
		for _, pack := range r.getNext(requiredItems, packSizes) {
			if _, ok := packs[pack]; !ok {
				packs[pack] = 0
			}

			packs[pack] += 1
			r.totalItems += pack
		}
	}

	return packs
}

func (r Rule) getNext(requiredItems int, packSizes []int) []int {
	single := r.getSingle(requiredItems, packSizes)
	multiple := r.getMultiple(requiredItems, packSizes)

	if (single != 0) && (multiple != nil) {
		if r.getExcessWidgets(requiredItems, []int{single}) > r.getExcessWidgets(requiredItems, multiple) {
			return multiple
		}

		return []int{single}
	}

	return []int{packSizes[len(packSizes)-1]}
}

func (r Rule) getSingle(requiredItems int, packSizes []int) int {
	for _, size := range packSizes {
		if (r.totalItems + size) >= requiredItems {
			return size
		}
	}

	return 0
}

func (r Rule) getMultiple(requiredItems int, packSizes []int) []int {
	for pack1 := 0; pack1 < len(packSizes)-1; pack1++ {
		for pack2 := 0; pack2 < len(packSizes)-1; pack2++ {
			if (r.totalItems + packSizes[pack1] + packSizes[pack2]) >= requiredItems {
				return []int{packSizes[pack1], packSizes[pack2]}
			}
		}
	}

	return nil
}

func (r Rule) getExcessWidgets(requiredItems int, packs []int) int {
	return r.packTotal(packs) - requiredItems - r.totalItems
}

func (r Rule) packTotal(packs []int) int {
	sum := 0

	for _, pack := range packs {
		sum += pack
	}

	return sum
}
