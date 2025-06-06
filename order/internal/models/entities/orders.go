package entities

import "github.com/jackc/pgx/v5/pgtype"

type Order struct {
	Id        pgtype.Int4      `db:"id"`
	UserId    pgtype.Int4      `db:"user_id"`
	Amount    pgtype.Int8      `db:"amount"`
	Status    pgtype.Text      `db:"status"`
	CreatedAt pgtype.Timestamp `db:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at"`
}
