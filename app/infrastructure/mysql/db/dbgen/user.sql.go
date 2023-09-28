// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: user.sql

package dbgen

import (
	"context"
)

const upsertUser = `-- name: UpsertUser :exec
INSERT INTO
   users (
      id,
      email,
      firebaseUid,
      phone_number,
      first_name,
      last_name,
      postal_code,
      prefecture,
      city,
      address_extra,
      created_at,
      updated_at
   )
VALUES
   (
      ?,
      ?,
      ?,
      ?,
      ?,
      ?,
      ?,
      ?,
      ?,
      ?,
      NOW(),
      NOW()
   ) ON DUPLICATE KEY
UPDATE
   email = ?,
   firebaseUid = ?,
   phone_number = ?,
   first_name = ?,
   last_name = ?,
   postal_code = ?,
   prefecture = ?,
   city = ?,
   address_extra = ?,
   updated_at = NOW()
`

type UpsertUserParams struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	Firebaseuid  string `json:"firebaseuid"`
	PhoneNumber  string `json:"phone_number"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	PostalCode   string `json:"postal_code"`
	Prefecture   string `json:"prefecture"`
	City         string `json:"city"`
	AddressExtra string `json:"address_extra"`
}

func (q *Queries) UpsertUser(ctx context.Context, arg UpsertUserParams) error {
	_, err := q.db.ExecContext(ctx, upsertUser,
		arg.ID,
		arg.Email,
		arg.Firebaseuid,
		arg.PhoneNumber,
		arg.FirstName,
		arg.LastName,
		arg.PostalCode,
		arg.Prefecture,
		arg.City,
		arg.AddressExtra,
		arg.Email,
		arg.Firebaseuid,
		arg.PhoneNumber,
		arg.FirstName,
		arg.LastName,
		arg.PostalCode,
		arg.Prefecture,
		arg.City,
		arg.AddressExtra,
	)
	return err
}

const userFindAll = `-- name: UserFindAll :many
SELECT
   id, email, firebaseuid, phone_number, first_name, last_name, postal_code, prefecture, city, address_extra, created_at, updated_at
FROM
   users
`

func (q *Queries) UserFindAll(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, userFindAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Firebaseuid,
			&i.PhoneNumber,
			&i.FirstName,
			&i.LastName,
			&i.PostalCode,
			&i.Prefecture,
			&i.City,
			&i.AddressExtra,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const userFindById = `-- name: UserFindById :one
SELECT
   id, email, firebaseuid, phone_number, first_name, last_name, postal_code, prefecture, city, address_extra, created_at, updated_at
FROM
   users
WHERE
   id = ?
`

func (q *Queries) UserFindById(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, userFindById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Firebaseuid,
		&i.PhoneNumber,
		&i.FirstName,
		&i.LastName,
		&i.PostalCode,
		&i.Prefecture,
		&i.City,
		&i.AddressExtra,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
