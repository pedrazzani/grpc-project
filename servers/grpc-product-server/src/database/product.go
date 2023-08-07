package database

import (
	"database/sql"
	"github.com/google/uuid"
)

type Product struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewProduct(db *sql.DB) *Product {
	return &Product{db: db}
}

func (p *Product) Create(name string, description string) (Product, error) {
	id := uuid.New().String()
	_, err := p.db.Exec("INSERT INTO products (id, name, description) VALUES ($1, $2, $3)",
		id, name, description)
	if err != nil {
		return Product{}, err
	}
	return Product{ID: id, Name: name, Description: description}, nil
}

func (p *Product) FindAll() ([]Product, error) {
	rows, err := p.db.Query("SELECT id, name, description FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}
		products = append(products, Product{ID: id, Name: name, Description: description})
	}
	return products, nil
}

func (p *Product) FindByCategoryId(categoryId string) ([]Product, error) {
	rows, err := p.db.Query("SELECT p.id, p.name, p.description FROM products p WHERE p.category_id = $1", categoryId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}
		products = append(products, Product{ID: id, Name: name, Description: description})
	}
	return products, nil
}
