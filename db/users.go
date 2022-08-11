package db

import (
	"context"
	"simple-chat-api/db/queries"
)

type UserEntity struct {
	Username string
	Passhash string
}

var userTableExists bool

func validateTable(ctx context.Context) {
	if !userTableExists {
		var exists bool
		usersTableRes, _ := GetConn(ctx).Query(
			ctx,
			queries.UserTableExists(),
		)
		usersTableRes.Scan(&exists)
		usersTableRes.Close()
		if !exists {
			creation, _ := GetConn(ctx).Query(
				ctx, queries.CreateUserTable(),
			)
			creation.Close()
		}

		userTableExists = true
	}
}

func GetUserByName(ctx context.Context, username string) (UserEntity, error) {
	validateTable(ctx)
	user := UserEntity{}
	query := queries.GetUserByName(username)

	res, err := GetConn(ctx).Query(ctx, query)
	res.Next()
	res.Scan(&(user.Username), &(user.Passhash))
	res.Close()

	return user, err
}

func CreateUser(ctx context.Context, user UserEntity) error {
	validateTable(ctx)
	query := queries.InsertUser(user.Username, user.Passhash)

	_, err := GetConn(ctx).Exec(ctx, query)
	return err
}
