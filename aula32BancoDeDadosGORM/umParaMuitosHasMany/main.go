package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primarykey"`
	Name     string
	Products []Product
}

type Product struct {
	ID           int          `gorm:"primary_key,AUTO_INCREMENT"`
	Name         string       `json:"name"`
	Price        float64      `json:"price"`
	CategoryID   int          `json:"category_id"` //Estamos relacionando duas tabelas
	Category     Category     //Estamos relacionando duas tabelas
	SerialNumber SerialNumber //`gorm:"foreignKey:ProductID"` //Relacionando Product com SerialNumber
	//gorm.Model                //Cria algumas colunas a mais no banco
}

type SerialNumber struct {
	ID        int `gorm:"primarykey"`
	Number    string
	ProductID int //Todos os SerialNumber terão um ProductId relacionados
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Category{}, &Product{}, &SerialNumber{}) //Cria as tabelas

	db.Create(&Category{
		Name: "Cama/Mesa",
	})

	db.Create(&Product{ //Criando Produto
		Name:       "Colchão",
		Price:      1300.0,
		CategoryID: 1, //Passando o Id da Categoria
	})

	db.Create(&SerialNumber{ //Criado SerialNumber
		Number:    "123",
		ProductID: 1,
	})

	/*var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products) //O Preload traz alguma informações a mais
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	}*/

	var categories []Category
	err = db.Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println("TESTE 1:", category.Name)
		for _, product := range category.Products {
			fmt.Println("TESTE 2:", product.Name, product.Category)
		}
	}

}
