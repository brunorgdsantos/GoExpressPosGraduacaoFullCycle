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

	var products []Product //Retornando todos os produtos
	db.Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}

	var products Product
	db.Limit(3).Offset(2).Find(&products) //O Limit serve para limiter o numero de registros a serem retornados
	for _, product := range products {    //O offset vai dividir o numero de registro em paginas. 3 registro serão apresentado na primeira pagina, outros 3 na segunda e assim por diante
		fmt.Println(product)
	}

	var products []Product
	db.Where("price > ?", 100).Find(&products) //Where("price > ?", 100) vai retornar todos os valores de prince maior que 100
	for _, product := range products {
		fmt.Println(product)
	}

	var product []Product
	db.Where("price LIKE ?", "%book").Find(&products) //Lista os produtos que tem book no nome
	for _, product := range products {
		fmt.Println(product)
	}
}
