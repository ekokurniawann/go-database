package main

import (
	"log"

	"github.com/ekokurniawann/gobd/pkg/invoice"
	"github.com/ekokurniawann/gobd/pkg/invoiceheader"
	"github.com/ekokurniawann/gobd/pkg/invoiceitem"
	"github.com/ekokurniawann/gobd/storage"
)

func main() {
	storage.ConnectPostgresDB()

	storageHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	storageItems := storage.NewPsqlInvoiceItem(storage.Pool())
	storageInvoice := storage.NewPsqlInvoice(
		storage.Pool(),
		storageHeader,
		storageItems,
	)

	m := &invoice.Model{
		Header: &invoiceheader.Model{
			Client: "Roger",
		},
		Items: invoiceitem.Models{
			&invoiceitem.Model{ProductID: 1},
		},
	}

	serviceInvoice := invoice.NewService(storageInvoice)
	if err := serviceInvoice.Create(m); err != nil {
		log.Fatalf("invoice.Create: %v", err)
	}
}
