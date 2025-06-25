package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Para iniciar os modulos: go mod init github.com/brunorgdsantos/GoExpressPosGraduacaoFullCycle
// Depois: go mod tidy
type Product struct {
	ID    int     `gorm:"primary_key,AUTO_INCREMENT"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}) //Migrate são usadas para criar as tabelas no Banco de dados com as estruturas ja corretas

	db.Create(&Product{ //Criado produto
		Name:  "Televisão",
		Price: 2200.00,
	})

	products := []Product{
		{Name: "TV", Price: 2600.00},
		{Name: "Teclado", Price: 100.00},
		{Name: "Mouse", Price: 200.00},
	}
	db.Create(&products)
}
