package main

import (
	"log"

	"github.com/ekokurniawann/gobd/pkg/invoiceitem"
	"github.com/ekokurniawann/gobd/storage"
)

func main() {
	storage.ConnectPostgresDB()

	storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
	serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)

	if err := serviceInvoiceItem.Migrate(); err != nil {
		log.Fatalf("invoiceItem.Migrate: %v", err)
	}
}
