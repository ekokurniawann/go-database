package main

import (
	"fmt"
	"log"

	"github.com/ekokurniawann/gobd/pkg/product"
	"github.com/ekokurniawann/gobd/storage"
)

func main() {
	storage.ConnectPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	ms, err := serviceProduct.GetAll()
	if err != nil {
		log.Fatalf("product.GetAll: %v", err)
	}

	fmt.Println(ms)
}
