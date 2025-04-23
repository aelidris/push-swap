package common

import (
	"os"
	"strconv"
	"strings"
)

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
