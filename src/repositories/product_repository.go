package repositories

import (
	"database/sql"
	"fmt"
	"go-api/src/models"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]models.Product, error) {
	query := "SELECT id, name, price FROM products"

	rows, err := pr.connection.Query((query))
	if err != nil {
		fmt.Println(err)
		return []models.Product{}, err
	}

	var productList []models.Product
	var product models.Product

	for rows.Next() {
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price)

		if err != nil {
			fmt.Println(err)

			return []models.Product{}, err
		}

		productList = append(productList, product)
	}

	rows.Close()

	return productList, nil
}

func (pr *ProductRepository) FindProduct(id int) (*models.Product, error) {
	query, err := pr.connection.Prepare("SELECT id, name, price FROM products WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var product models.Product

	err = query.QueryRow(id).Scan(
		&product.ID,
		&product.Name,
		&product.Price)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()

	return &product, nil
}

func (pr *ProductRepository) CreateProduct(product models.Product) (int, error) {
	var id int

	query, err := pr.connection.Prepare("INSERT INTO products (name, price) VALUES ($1, $2)" +
		"RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()

	return id, nil
}

func (pr *ProductRepository) DeleteProduct(id int) error {
	query, err := pr.connection.Prepare("DELETE FROM products WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return err
	}

	var product models.Product

	err = query.QueryRow(id).Scan(
		&product.ID,
		&product.Name,
		&product.Price)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}

		return nil
	}

	query.Close()

	return nil
}
