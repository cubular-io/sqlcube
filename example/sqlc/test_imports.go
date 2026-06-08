package sqlc

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// These structs have the same fields as User in models.go
// After reduce, they become aliases and the imports become unused
type GetUserByIDRow struct {
	ID    uuid.UUID   `json:"id"`
	Name  pgtype.Text `json:"name"`
	Email pgtype.Text `json:"email"`
}

type GetUserByEmailRow struct {
	ID    uuid.UUID   `json:"id"`
	Name  pgtype.Text `json:"name"`
	Email pgtype.Text `json:"email"`
}

type ListUsersRow struct {
	ID    uuid.UUID   `json:"id"`
	Name  pgtype.Text `json:"name"`
	Email pgtype.Text `json:"email"`
}
