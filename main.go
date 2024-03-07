package main

import (
	"log"

	"github.com/ekokurniawann/gobd/pkg/product"
	"github.com/ekokurniawann/gobd/storage"
)

func main() {
	storage.ConnectPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	err := serviceProduct.Delete(2)
	if err != nil {
		log.Fatalf("product.Update: %v", err)
	}
}
