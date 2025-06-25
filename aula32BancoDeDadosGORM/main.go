package main

import (
	"fmt"
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

	products := []Product{ //Criado varios produtos de uma vez
		{Name: "TV", Price: 2600.00},
		{Name: "Teclado", Price: 100.00},
		{Name: "Mouse", Price: 200.00},
	}
	db.Create(&products)

	var produto Product
	db.First(&produto, 7) //Retornando o produto com id=7
	fmt.Println(produto)

	db.First(&produto, "name = ?", "TV") //Retornando o produto com name=TV
	fmt.Println(produto)

	var produts []Product //Retornando todos os produtos
	db.Find(&produts)
	for _, product := range products {
		fmt.Println(product)
	}
}
