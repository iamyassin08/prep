// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)


type User struct {
	ID               int32
	UserID        pgtype.Text
	FirstName            string
	Email             pgtype.Numeric
	LastName      pgtype.Numeric

}














