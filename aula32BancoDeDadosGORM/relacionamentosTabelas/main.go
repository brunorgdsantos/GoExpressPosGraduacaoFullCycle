package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primarykey"`
	Name string
}

type Product struct {
	ID         int      `gorm:"primary_key,AUTO_INCREMENT"`
	Name       string   `json:"name"`
	Price      float64  `json:"price"`
	CategoryID int      `json:"category_id"` //Estamos relacionando duas tabelas
	Category   Category //Estamos relacionando duas tabelas
	gorm.Model          //Cria algumas colunas a mais no banco
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Category{}, &Product{})

	Category := Category{ //Criando Categorias
		Name: "Eletronicos",
	}
	db.Create(&Category)

	db.Create(&Product{ //Criando Produto
		Name:       "Dell",
		Price:      30.0,
		CategoryID: Category.ID, //Passando o Id da Categoria
	})

	var products []Product
	db.Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name)
	}

	var products []Product
	db.Preload("Category").Find(&products) //O Preload traz alguma informações a mais
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name)
	}

}
