// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: product.sql

package dbgen

import (
	"context"
	"database/sql"
	"strings"
	"time"
)

const productFetchWithOwner = `-- name: ProductFetchWithOwner :many
SELECT
  products.id, products.owner_id, products.name, products.description, products.price, products.stock, products.created_at, products.updated_at,
  owners.name AS owner_name
FROM
  products
  LEFT OUTER JOIN owners ON products.owner_id = owners.id
`

type ProductFetchWithOwnerRow struct {
	ID          string         `json:"id"`
	OwnerID     string         `json:"owner_id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       int64          `json:"price"`
	Stock       int32          `json:"stock"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	OwnerName   sql.NullString `json:"owner_name"`
}

func (q *Queries) ProductFetchWithOwner(ctx context.Context) ([]ProductFetchWithOwnerRow, error) {
	rows, err := q.db.QueryContext(ctx, productFetchWithOwner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ProductFetchWithOwnerRow{}
	for rows.Next() {
		var i ProductFetchWithOwnerRow
		if err := rows.Scan(
			&i.ID,
			&i.OwnerID,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.Stock,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.OwnerName,
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

const productFindById = `-- name: ProductFindById :one
SELECT
   id, owner_id, name, description, price, stock, created_at, updated_at
FROM
   products
WHERE
   id = ?
`

func (q *Queries) ProductFindById(ctx context.Context, id string) (Product, error) {
	row := q.db.QueryRowContext(ctx, productFindById, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.Stock,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const productFindByIds = `-- name: ProductFindByIds :many
SELECT
   id, owner_id, name, description, price, stock, created_at, updated_at
FROM
   products
WHERE
   id IN (/*SLICE:ids*/?)
`

func (q *Queries) ProductFindByIds(ctx context.Context, ids []string) ([]Product, error) {
	query := productFindByIds
	var queryParams []interface{}
	if len(ids) > 0 {
		for _, v := range ids {
			queryParams = append(queryParams, v)
		}
		query = strings.Replace(query, "/*SLICE:ids*/?", strings.Repeat(",?", len(ids))[1:], 1)
	} else {
		query = strings.Replace(query, "/*SLICE:ids*/?", "NULL", 1)
	}
	rows, err := q.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.OwnerID,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.Stock,
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

const upsertProduct = `-- name: UpsertProduct :exec
INSERT INTO products (
   id,
   owner_id,
   name,
   description,
   price,
   stock
) VALUES (
   ?,
   ?,
   ?,
   ?,
   ?,
   ?
) ON DUPLICATE KEY UPDATE
   owner_id = ?,
   name = ?,
   description = ?,
   price = ?,
   stock = ?
`

type UpsertProductParams struct {
	ID          string `json:"id"`
	OwnerID     string `json:"owner_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Stock       int32  `json:"stock"`
}

func (q *Queries) UpsertProduct(ctx context.Context, arg UpsertProductParams) error {
	_, err := q.db.ExecContext(ctx, upsertProduct,
		arg.ID,
		arg.OwnerID,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.Stock,
		arg.OwnerID,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.Stock,
	)
	return err
}
