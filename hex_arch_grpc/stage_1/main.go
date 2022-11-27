package main

import (
	"fmt"
	"hex_stage1/internal/adapters/core/arithmetic"
	"hex_stage1/internal/ports"
)

func main() {
	var core ports.ArithmeticPort
	core = arithmetic.NewAdapter()

	result, err := core.Addition(1, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
