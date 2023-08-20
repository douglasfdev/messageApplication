package repository

import (
	"baby-steps/internal/entity"
	"database/sql"
)

type ProductRepositoryMySql struct {
	DB *sql.DB
}

func NewProductRepositoryMySql(db *sql.DB) *ProductRepositoryMySql {
	return &ProductRepositoryMySql{DB: db}
}

func (r *ProductRepositoryMySql) Create(product *entity.Product) error {
	_, err := r.DB.Exec("insert into products (id, name, price) values(?,?,?)",
		product.ID, product.Name, product.Price)

	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepositoryMySql) FindAll() ([]*entity.Product, error) {
	rows, err := r.DB.Query("select id, name, price from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var product entity.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price)

		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}
