package service

import (
	"belajar/bareng/features"
	"belajar/bareng/features/product"

	"github.com/go-playground/validator/v10"
)

type ProductService struct {
	productService product.ProductData
	validate       *validator.Validate
}

// SelectByUserId implements product.ProductServise.
func (service *ProductService) SelectByUserId(user_id uint) error {
	err := service.productService.SelectByUserId(user_id)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements product.ProductServise
func (service *ProductService) Delete(id uint) error {
	err := service.productService.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// Edit implements product.ProductServise
func (service *ProductService) Edit(id uint, product features.ProductEntity) (uint, error) {
	data, err := service.productService.Update(id, product)
	if err != nil {
		return 0, err
	}
	return data, nil
}

// GetById implements product.ProductServise
func (servise *ProductService) GetById(id uint) (features.ProductEntity, error) {
	data, err := servise.productService.SelectById(id)
	if err != nil {
		return features.ProductEntity{}, err
	}
	return data, nil
}

// GetAll implements product.ProductServise
func (service *ProductService) GetAll() ([]features.ProductEntity, error) {
	data, err := service.productService.SelectAll()
	if err != nil {
		return []features.ProductEntity{}, err
	}
	return data, nil
}

// Add implements product.ProductServise
func (servise *ProductService) Add(product features.ProductEntity) (uint, error) {

	errValidate := servise.validate.Struct(product)
	if errValidate != nil {
		return 0, errValidate
	}
	id, err := servise.productService.Insert(product)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func New(products product.ProductData) product.ProductServise {
	return &ProductService{
		productService: products,
		validate:       validator.New(),
	}
}
