package solver

import (
	"fmt"
	"os"
	ops "push-swap/common"
	"sort"
)

func Solve(stackA []int) {
	if sort.IntsAreSorted(stackA) {
		return
	}
	switch len(stackA) {
	case 2:
		SortTwo(&stackA)
	case 3:
		SortThree(&stackA)
	default:
		SortStack(&stackA)
	}
}

func SortTwo(stackA *[]int) {
	a := *stackA
	if a[0] > a[1] {
		ops.Sa(a)
		fmt.Println("sa")
	}
}

func SortThree(stackA *[]int) {
	a := *stackA
	sorted := make([]int, 3)
	copy(sorted, a)
	sort.Ints(sorted)

	s, m, l := sorted[0], sorted[1], sorted[2]

	switch {
	case m == a[0] && s == a[1] && l == a[2]:
		fmt.Println("sa")
	case l == a[0] && m == a[1] && s == a[2]:
		fmt.Println("sa")
		fmt.Println("rra")
	case l == a[0] && s == a[1] && m == a[2]:
		fmt.Println("ra")
	case s == a[0] && l == a[1] && m == a[2]:
		fmt.Println("sa")
		fmt.Println("ra")
	case m == a[0] && l == a[1] && s == a[2]:
		fmt.Println("rra")
	}

	a[0], a[1], a[2] = s, m, l
}

func SortStack(stackA *[]int) {
	a := *stackA
	b := make([]int, 0, 100)
	for len(a) != 3 {
		min, index := GetMin(a)
		half := len(a) / 2
		if half >= index {
			for min != a[0] {
				if min == a[1] {
					ops.Sa(a)
					fmt.Println("sa")
				} else {
					ops.Ra(a)
					fmt.Println("ra")
				}
			}
		} else {
			for min != a[0] {
				ops.Rra(a)
				fmt.Println("rra")
			}
		}
		ops.Pb(&a, &b)
		fmt.Println("pb")
	}

	SortThree(&a)

	for len(b) != 0 {
		ops.Pa(&a, &b)
		fmt.Println("pa")
	}
	// Print sorted array if environement variable sorted == show
	if os.Getenv("sorted") == "show" {
		fmt.Println("\033[33m", a, "\033[0m")
	}
}

func GetMin(a []int) (int, int) {
	min := a[0]
	index := 0
	for i := 0; i < len(a); i++ {
		if min > a[i] {
			min = a[i]
			index = i
		}
	}
	return min, index
}
