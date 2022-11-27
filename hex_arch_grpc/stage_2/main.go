package main

import (
	"fmt"
	"hex_stage2/internal/adapters/app/api"
	"hex_stage2/internal/adapters/core/arithmetic"
)

func main() {
	arith := arithmetic.NewAdapter()
	appArith := api.NewAdapter(arith)

	result, err := appArith.GetAddition(1, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
