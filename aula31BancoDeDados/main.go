package main

import (
	"database/sql"
	"fmt"
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

	p, err := selectProduct(db, produto.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Product: %v possui o preço de %2.f", p.Name, p.Price)

	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, prod := range products { //chave e valor (_, prod)
		fmt.Printf("Produto %v, possui o preço de %.2f", prod.Name, prod.Price)
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

	_, err = stmt.Exec(product.Name, product.Price, product.ID) //(product.ID, product.Name, product.Price) vão substituir as (?,?,?) da linha 60
	if err != nil {
		return err
	}
	return nil
}

func selectProduct(db *sql.DB, id string) (*Product, error) { //Read Produtos, selecionando produtos
	stmt, err := db.Prepare("select id, name, price from products where id = ?") //Segurança contra SQL Injection
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var p Product

	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price) //Faz as atribuições
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select ID, Name, Price from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
