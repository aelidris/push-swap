package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	ops "push-swap/common"
)

func main() {
	switch len(os.Args) {
	case 1:
		return
	case 2:
		if os.Args[1] == "" {
			return
		}
		stackA := ops.ReadArgs()
		if stackA == nil {
			fmt.Println("Error")
			return
		}
		if err := Checker(&stackA); err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println("Error")
	}
}

func Checker(stackA *[]int) error {

	var stackB []int
	var inputOps []string
	operations := []string{"pa", "pb", "sa", "sb", "ss", "ra", "rb", "rr", "rra", "rrb", "rrr"}

	if errorDetected := Readop(operations, &inputOps); errorDetected {
		return fmt.Errorf("Error")
	}

	for _, instruct := range inputOps {
		switch instruct {
		case "pa":
			ops.Pa(stackA, &stackB)
		case "pb":
			ops.Pb(stackA, &stackB)
		case "sa":
			ops.Sa(*stackA)
		case "sb":
			ops.Sb(stackB)
		case "ss":
			ops.Ss(*stackA, stackB)
		case "ra":
			ops.Ra(*stackA)
		case "rb":
			ops.Rb(stackB)
		case "rr":
			ops.Rr(*stackA, stackB)
		case "rra":
			ops.Rra(*stackA)
		case "rrb":
			ops.Rrb(stackB)
		}
	}
	if sort.IntsAreSorted(*stackA) && len(stackB) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
	return nil
}

// Reads instructions on the standard input and saves them. If gets non-existing insturction - returns
func Readop(operations []string, inputop *[]string) bool {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		for i, op := range operations {
			if scanner.Text() == "stop" {
				return false
			}

			if scanner.Text() == "" {
				continue
			}

			if scanner.Text() == op {
				*inputop = append(*inputop, scanner.Text())
				break
			}

			if i == len(operations)-1 {
				fmt.Println("Error")
				return true
			}
		}
	}
	return false
}
