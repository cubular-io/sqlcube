package sqlc

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID    uuid.UUID   `json:"id"`
	Name  pgtype.Text `json:"name"`
	Email pgtype.Text `json:"email"`
}
