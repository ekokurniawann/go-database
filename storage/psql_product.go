package storage

import (
	"database/sql"
	"fmt"

	"github.com/ekokurniawann/gobd/pkg/product"
)

const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT products_id_pk PRIMARY KEY (id) 
	)`
	psqlCreateProduct = `INSERT INTO products(name, observations, price, created_at) VALUES($1, $2, $3, $4) RETURNING id`
)

type psqlProduct struct {
	db *sql.DB
}

func NewPsqlProduct(db *sql.DB) *psqlProduct {
	return &psqlProduct{db}
}

func (p *psqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Migrasi produk berhasil dijalankan")
	return nil
}

func (p *psqlProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		m.CreatedAt,
	).Scan(&m.ID)
	if err != nil {
		return err
	}

	fmt.Println("Produk berhasil dibuat")
	return nil
}
