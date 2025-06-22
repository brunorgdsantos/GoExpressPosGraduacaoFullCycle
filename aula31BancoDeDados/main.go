package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"log"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	produto := NewProduct("Notebook", 1240.0) //Inserindo dados no Banco
	err = insertProduct(db, produto)
	if err != nil {
		log.Fatal(err)
	}

	produto.Price = 100.0
	err = updateProduct(db, produto)
	if err != nil {
		log.Fatal(err)
	}
}

func insertProduct(db *sql.DB, product *Product) error { //Inserindo dados no Banco
	stmt, err := db.Prepare("Insert into products(ID, Name,Price) values(?,?,?)") //Segurança contra SQL Injection
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price) //(product.ID, product.Name, product.Price) vão substituir as (?,?,?) da linha 35
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, product *Product) error { //Atualizando dados do banco
	stmt, err := db.Prepare("Update products set Name = ?,Price = ? where ID = ?") //Segurança contra SQL Injection
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.ID) //(product.ID, product.Name, product.Price) vão substituir as (?,?,?) da linha 35
	if err != nil {
		return err
	}
	return nil
}
