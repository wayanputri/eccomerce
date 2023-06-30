package product

import "belajar/bareng/features"

type ProductData interface {
	Insert(product features.ProductEntity) (uint,error)
	SelectAll()([]features.ProductEntity,error)
	SelectById(id uint)(features.ProductEntity,error)
	Update(id uint,product features.ProductEntity)(uint,error)
	Delete(id uint)(error)
	SelectByUserId(user_id uint)(error)
}

type ProductServise interface{
	Add(product features.ProductEntity) (uint,error)
	GetAll()([]features.ProductEntity,error)
	GetById(id uint)(features.ProductEntity,error)
	Edit(id uint,product features.ProductEntity) (uint,error)
	Delete(id uint)(error)
	SelectByUserId(user_id uint)(error)
}