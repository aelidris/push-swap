package main

import (
	"fmt"
	"os"

	read "push-swap/common"
	sol "push-swap/pushswap/solve"
)

func main() {
	switch len(os.Args) {
	case 1:
		return
	case 2:
		if os.Args[1] == "" {
			return
		}
		stackA := read.ReadArgs()
		if stackA == nil {
			fmt.Println("Error")
			return
		}
		sol.Solve(stackA)
	default:
		fmt.Println("Error")
	}
}
