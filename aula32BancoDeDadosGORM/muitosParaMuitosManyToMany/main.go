package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primarykey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"` //Quando se trabalha com Many To Many no GO, é necessário especificar. Em relações ManyToMany é necessário criar uma tabela intermediária
}

type Product struct {
	ID         int        `gorm:"primary_key"`
	Name       string     //`json:"name"`
	Price      float64    //`json:"price"`
	Categories []Category `gorm:"many2many:products_categories;"` //Estamos relacionando duas tabelas
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}) //Cria as tabelas

	category := Category{Name: "Cozinha"}
	db.Create(&category)

	category2 := Category{Name: "Eletronicos"}
	db.Create(&category2)

	db.Create(&Product{
		Name:       "Microondas",
		Price:      1050.0,
		Categories: []Category{category, category2}, //Este Product pertence as duas categorias
	})
	/*
		var categories []Category
		err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error //Observe o Products.SerialNumber
		if err != nil {
			panic(err)
		}

		for _, category := range categories {
			fmt.Println("TESTE 1: ", category.Name)
			for _, product := range category.Products {
				fmt.Println("TESTE 2: ", product.Name, category.Name)
			}
		}*/

}
