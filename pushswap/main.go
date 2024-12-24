package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"push-swap/pushswap/solve"
)

func main() {
	switch len(os.Args) {
	case 1:
		return
	case 2:
		if os.Args[1] == "" {
			return
		}
		stackA := ReadArgs()
		if stackA == nil {
			fmt.Println("Error")
			return
		}
		solver.Solve(stackA)
	default:
		fmt.Println("Error")
	}
}

func ReadArgs() (stackA []int) {
	stackAMap := make(map[int]bool)
	for _, num := range strings.Fields(os.Args[1]) {
		n, err := strconv.Atoi(num)
		if err != nil {
			return nil
		}
		if exist := stackAMap[n]; exist {
			return nil
		}
		stackAMap[n] = true
		stackA = append(stackA, n)
	}
	return stackA
}
