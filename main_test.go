package AlgorithmsAndTasks

import (
	"fmt"
	"sort"
	"testing"
)

type Struct struct {
	Name string
}

func TestSlice(t *testing.T) {
	sl := [3]Struct{
		{Name: "b"},
		{Name: "c"},
		{Name: "a"},
	}

	fmt.Println(sl)
	fmt.Println(sl[:])

	newSL := sl[:]
	sort.Slice(newSL, func(i, j int) bool {
		return newSL[i].Name < newSL[j].Name
	})

	fmt.Println(newSL)
	fmt.Println(sl)
}
