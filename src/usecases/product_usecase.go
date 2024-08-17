package usecases

import (
	"go-api/src/models"
	"go-api/src/repositories"
)

type ProductUsecase struct {
	repository repositories.ProductRepository
}

func NewProductUsecase(repo repositories.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]models.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) FindProduct(id int) (*models.Product, error) {
	product, err := pu.repository.FindProduct(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pu *ProductUsecase) CreateProduct(product models.Product) (models.Product, error) {
	productId, err := pu.repository.CreateProduct(product)

	if err != nil {
		return models.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductUsecase) DeleteProduct(id int) error {
	err := pu.repository.DeleteProduct(id)

	if err != nil {
		return err
	}

	return nil
}
