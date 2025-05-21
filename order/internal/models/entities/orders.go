package entities

import "github.com/jackc/pgx/v5/pgtype"

type Order struct {
	Id        pgtype.Int4
	UserId    pgtype.Int4
	Amount    pgtype.Int8
	Status    pgtype.Text
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}
