package user

import (
	"context"
	"encore.dev/storage/sqldb"
)

// insert inserts a user into the database.
func insert(ctx context.Context, email string) (int64, error) {
	exec, err := sqldb.Exec(ctx, `
		INSERT INTO "user" (email)
		VALUES ($1)
	`, email)

	return exec.RowsAffected(), err
}
