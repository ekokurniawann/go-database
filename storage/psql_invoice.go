package storage

import (
	"database/sql"
	"fmt"

	"github.com/ekokurniawann/gobd/pkg/invoice"
	"github.com/ekokurniawann/gobd/pkg/invoiceheader"
	"github.com/ekokurniawann/gobd/pkg/invoiceitem"
)

type PsqlInvoice struct {
	db            *sql.DB
	storageHeader invoiceheader.Storage
	storageItems  invoiceitem.Storage
}

func NewPsqlInvoice(db *sql.DB, h invoiceheader.Storage, i invoiceitem.Storage) *PsqlInvoice {
	return &PsqlInvoice{
		db:            db,
		storageHeader: h,
		storageItems:  i,
	}
}

func (p *PsqlInvoice) Create(m *invoice.Model) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	if err := p.storageHeader.CreateTranscation(tx, m.Header); err != nil {
		tx.Rollback()
		return fmt.Errorf("Header: %w", err)
	}
	fmt.Printf("Faktur berhasil dibuat dengan ID: %d \n", m.Header.ID)

	if err := p.storageItems.CreateTranscation(tx, m.Header.ID, m.Items); err != nil {
		tx.Rollback()
		return fmt.Errorf("Item: %w", err)
	}
	fmt.Printf("Jumlah item yang berhasil dibuat: %d \n", len(m.Items))

	return tx.Commit()
}
