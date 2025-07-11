package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primarykey"`
	Name     string
	Products []Product //Uma categoria pode ter muitos produtos mas um produto tem apenas uma categoria
}

type Product struct {
	ID           int      `gorm:"primary_key"`
	Name         string   //`json:"name"`
	Price        float64  //`json:"price"`
	CategoryID   int      //`json:"category_id"` //Estamos relacionando duas tabelas
	Category     Category //Estamos relacionando duas tabelas
	SerialNumber SerialNumber
	//gorm.Model                //Cria algumas colunas a mais no banco
}

type SerialNumber struct {
	ID        int
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, SerialNumber{}) //Cria as tabelas

	category := Category{ //Criando Categorias
		Name: "Cozinha",
	}
	db.Create(&category)

	db.Create(&Product{ //Criando Produto
		Name:       "Colchão",
		Price:      1300.0,
		CategoryID: 1, //Passando o Id da Categoria
	})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error //Observe o Products.SerialNumber
	if err != nil {
		panic(err)
	}

	db.Create(&SerialNumber{
		Number:    "123",
		ProductID: 1,
	})

	for _, category := range categories {
		fmt.Println("TESTE 1: ", category.Name)
		for _, product := range category.Products {
			fmt.Println("TESTE 2: ", product.Name, category.Name, "SERIAL NUMBER: ", product.SerialNumber.Number)
		}
	}

}
