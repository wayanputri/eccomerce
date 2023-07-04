package data

import (
	"belajar/bareng/features"
	"belajar/bareng/features/product"
	"errors"

	"gorm.io/gorm"
)

type ProductsData struct {
	db *gorm.DB
}

// SelectByUserId implements product.ProductData.
func (repo *ProductsData) SelectByUserId(user_id uint) error {
	var product features.Product
	tx := repo.db.First(&product, "user_id=?", user_id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Delete implements product.ProductData
func (repo *ProductsData) Delete(id uint) error {
	var product features.Product
	tx := repo.db.Delete(&product, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Update implements product.ProductData
func (repo *ProductsData) Update(id uint, product features.ProductEntity) (uint, error) {
	var productModal features.Product
	tx := repo.db.Model(&productModal).Where("id=?", id).Updates(features.ProductEntityToModel(product))
	if tx.Error != nil {
		return 0, tx.Error
	}
	return id, nil
}

// SelectById implements product.ProductData
func (repo *ProductsData) SelectById(id uint) (features.ProductEntity, error) {
	var ProductBase features.Product
	tx := repo.db.First(&ProductBase, id)
	if tx.Error != nil {
		return features.ProductEntity{}, tx.Error
	}
	data := features.ProductModelToEntity(ProductBase)
	return data, nil
}

// SelectAll implements product.ProductData
func (repo *ProductsData) SelectAll() ([]features.ProductEntity, error) {
	var productModel []features.Product

	tx := repo.db.Find(&productModel)
	if tx.Error != nil {
		return []features.ProductEntity{}, tx.Error
	}
	var dataEntity []features.ProductEntity
	for _, products := range productModel {
		data := features.ProductModelToEntity(products)
		dataEntity = append(dataEntity, data)
	}

	return dataEntity, nil
}

// Insert implements product.ProductData
func (repo *ProductsData) Insert(product features.ProductEntity) (uint, error) {
	productModel := features.ProductEntityToModel(product)
	tx := repo.db.Create(&productModel)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("gagal create data")
	}
	return productModel.ID, nil
}

func New(db *gorm.DB) product.ProductData {
	return &ProductsData{
		db: db,
	}
}
