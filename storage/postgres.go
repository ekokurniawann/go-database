package storage

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/ekokurniawann/gobd/pkg/product"
	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

func ConnectPostgresDB() {
	once.Do(func() {
		var err error

		db, err = sql.Open("postgres", "postgres://postgres:123456@localhost:5432/test-db?sslmode=disable")
		if err != nil {
			log.Fatalf("tidak dapat membuka db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("tidak dapat melakukan ping: %v", err)
		}

		fmt.Println("terhubung dengan postgres")
	})
}

func Pool() *sql.DB {
	return db
}

func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}
	if null.String != "" {
		null.Valid = true
	}

	return null
}

type scanner interface {
	Scan(dest ...any) error
}

func scanRowProduct(s scanner) (*product.Model, error) {
	m := &product.Model{}
	observationNull := sql.NullString{}
	updatedAtNull := sql.NullTime{}

	err := s.Scan(
		&m.ID,
		&m.Name,
		&observationNull,
		&m.Price,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return &product.Model{}, err
	}

	m.Observations = observationNull.String
	m.UpdatedAt = updatedAtNull.Time

	return m, nil
}
