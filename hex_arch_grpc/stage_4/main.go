package main

import (
	"fmt"
	"hex_stage4/internal/adapters/app/api"
	"hex_stage4/internal/adapters/core/arithmetic"
	"hex_stage4/internal/adapters/framework/right/db"
)

func main() {
	arith := arithmetic.NewAdapter()
	db := db.NewAdapter()
	appArith := api.NewAdapter(db, arith)

	result, err := appArith.GetAddition(1, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
