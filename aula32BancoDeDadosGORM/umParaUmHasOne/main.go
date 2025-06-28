package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primarykey, AUTO_INCREMENT"`
	Name string
}

type Product struct {
	ID           int          `gorm:"primary_key"`
	Name         string       //`json:"name"`
	Price        float64      //`json:"price"`
	CategoryID   int          //`gorm:"foreignKey:CategoryID"` //Estamos relacionando duas tabelas
	Category     Category     //Estamos relacionando duas tabelas
	SerialNumber SerialNumber //`gorm:"foreignKey:ProductID"` //Relacionando Product com SerialNumber
	//gorm.Model //Cria algumas colunas a mais no banco
	//Cuidado ao utilizar o gorm.Model, pode ocultar alguma informações
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
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{}) //Cria as tabelas

	category := Category{ //Criando Categorias
		Name: "Cozinha",
	}
	db.Create(&category)

	db.Create(&Product{ //Criando Produto
		Name:       "Fogão",
		Price:      1500.0,
		CategoryID: 1, //Passando o Id da Categoria
	})

	db.Create(&SerialNumber{ //Criado SerialNumber
		Number:    "123",
		ProductID: 1,
	})

	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products) //O Preload traz alguma informações a mais
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	}

}
