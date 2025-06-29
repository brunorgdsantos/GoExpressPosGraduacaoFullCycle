package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	tx := db.Begin()
	var c Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		panic(err)
	}

	c.Name = "TESTE"
	tx.Debug().Save(&c)
	tx.Commit()
}
