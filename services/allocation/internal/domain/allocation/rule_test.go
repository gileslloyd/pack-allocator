package allocation

import (
	"reflect"
	"strconv"
	"testing"
)

type allocation struct {
	requiredItems int
	num_250       int
	num_500       int
	num_1000      int
	num_2000      int
	num_5000      int
}

func getAllocationTable() []allocation {
	return []allocation{
		allocation{1, 1, 0, 0, 0, 0},
		allocation{250, 1, 0, 0, 0, 0},
		allocation{251, 0, 1, 0, 0, 0},
		allocation{501, 1, 1, 0, 0, 0},
		allocation{12001, 1, 0, 0, 1, 2},
		allocation{752, 0, 0, 1, 0, 0},
		allocation{15521, 1, 1, 0, 0, 3},
	}
}

func TestCalculatingPackAllocation(t *testing.T) {
	rule := NewPackAllocationRule()
	packSizes := []int{250, 500, 1000, 2000, 5000}

	for _, table := range getAllocationTable() {
		actualAllocation := rule.CalculatePackAllocation(table.requiredItems, packSizes)

		for _, packSize := range packSizes {
			r := reflect.ValueOf(table)
			f := reflect.Indirect(r).FieldByName("num_" + strconv.Itoa(packSize))
			expectedAllocation := int(f.Int())

			if actualAllocation[packSize] != expectedAllocation {
				t.Errorf("Incorrect allocation of %d item packs. Got %d, expected %d", packSize, actualAllocation[packSize], expectedAllocation)
			}
		}
	}
}
