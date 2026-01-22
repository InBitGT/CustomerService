package migration

import (
	"fmt"

	"CustomerService/db"
	"CustomerService/internal/modules/address"
	"CustomerService/internal/modules/branch"
	"CustomerService/internal/modules/user_branch"
)

func Migration() {
	database := db.Database()

	err := database.AutoMigrate(
		&address.Address{},
		&branch.Branch{},
		&user_branch.UserBranch{},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Migraciones de SentiService ejecutadas")
}
