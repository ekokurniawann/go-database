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

	m := &product.Model{
		Name:         "Belajar golang asik",
		Price:        100,
		Observations: "wih",
	}
	if err := serviceProduct.Create(m); err != nil {
		log.Fatalf("product.Create: %v", err)
	}

	fmt.Printf("%+v\n", m)
}
