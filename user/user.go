package user

import (
	"context"
	"encore.dev/storage/sqldb"
)

// Get returns a user based on its email
//
//encore:api auth method=GET path=/user/:email
func Get(ctx context.Context, email string) (*User, error) {
	u := &User{Email: email}
	err := sqldb.QueryRow(ctx, `
		SELECT user_id FROM "user"
		WHERE email = $1
	`, email).Scan(&u.Id)
	return u, err
}

// Create creates a user receiving an email
//
//encore:api auth method=POST path=/user
func Create(ctx context.Context, p *CreateRequest) (*CreateResponse, error) {
	rows, err := insert(ctx, p.Email)
	if err != nil {
		return nil, err
	}
	return &CreateResponse{Rows: rows, Status: "Created"}, nil
}
