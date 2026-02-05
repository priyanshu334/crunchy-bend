package db

import "github.com/priyanshu334/go-bend/internal/moduels/user"

func Migrate() error {
	return DB.AutoMigrate(
		&user.User,
	)

}
