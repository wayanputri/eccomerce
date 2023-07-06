package features

func UserModelToEntity(user User) UserEntity {

	var transactionEntity []TransactionEntity
	for _, transaction := range user.Transactions {
		transactionEntity = append(transactionEntity, TransactionModelToEntity(transaction))
	}
	var productEntity []ProductEntity
	for _, product := range user.Products {
		productEntity = append(productEntity, ProductModelToEntity(product))
	}

	return UserEntity{
		Id:           user.ID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		DeletedAt:    user.DeletedAt.Time,
		Nama:         user.Nama,
		NoTlp:        user.NoTlp,
		Email:        user.Email,
		Password:     user.Password,
		Alamat:       user.Alamat,
		Transactions: transactionEntity,
		Product:      productEntity,
	}
}

func PaymentModelToEntity(payment Payment) PaymentEntity {
	return PaymentEntity{
		Id:            payment.ID,
		CreatedAt:     payment.CreatedAt,
		UpdatedAt:     payment.UpdatedAt,
		DeletedAt:     payment.DeletedAt.Time,
		TransactionID: payment.TransactionID,
		Transactions:  TransactionModelToEntity(payment.Transactions),
		Status:        payment.Status,
		Bank:          payment.Bank,
		VA:            payment.VA,
		OrderID:       payment.OrderID,
	}
}

func TransactionModelToEntity(transaction Transaction) TransactionEntity {
	return TransactionEntity{
		Id:           transaction.ID,
		CreatedAt:    transaction.CreatedAt,
		UpdatedAt:    transaction.UpdatedAt,
		DeletedAt:    transaction.DeletedAt.Time,
		ProductID:    transaction.ProductID,
		Products:     ProductModelToEntity(transaction.Products),
		UserID:       transaction.UserID,
		Users:        UserModelToEntity(transaction.Users),
		Status:       transaction.Status,
		TotalHarga:   transaction.TotalHarga,
		JumlahBarang: transaction.JumlahBarang,
	}
}

func ProductModelToEntity(product Product) ProductEntity {
	var transactionEntity []TransactionEntity
	for _, transaction := range product.Transactions {
		transactionEntity = append(transactionEntity, TransactionModelToEntity(transaction))
	}
	var imageEntity []ImageEntity
	for _, image := range product.Images {
		imageEntity = append(imageEntity, ImageModelToEntity(image))
	}
	var reviewEntity []ReviewEntity
	for _, review := range product.Reviews {
		reviewEntity = append(reviewEntity, ReviewModelEntity(review))
	}

	return ProductEntity{
		Id:           product.ID,
		CreatedAt:    product.CreatedAt,
		UpdatedAt:    product.UpdatedAt,
		DeletedAt:    &product.DeletedAt.Time,
		Nama:         product.Nama,
		Harga:        product.Harga,
		Deskripsi:    product.Deskripsi,
		Stok:         product.Stok,
		Transactions: transactionEntity,
		Images:       imageEntity,
		Users:        UserModelToEntity(product.Users),
		Reviews:      reviewEntity,
	}
}
func ImageModelToEntity(image Image) ImageEntity {
	return ImageEntity{
		Id:        image.ID,
		CreatedAt: image.CreatedAt,
		UpdatedAt: image.UpdatedAt,
		DeletedAt: image.DeletedAt.Time,
		ProductID: image.ProductID,
		Products:  ProductModelToEntity(image.Products),
		Link:      image.Link,
		Nama:      image.Nama,
	}
}
func ReviewModelEntity(review Review) ReviewEntity {
	var reviewImages []ReviewImagesEntity
	for _, reviewImage := range review.ImagesReview {
		reviewImages = append(reviewImages, ReviewImageModelEntity(reviewImage))
	}
	return ReviewEntity{
		Id:           review.ID,
		CreatedAt:    review.CreatedAt,
		UpdatedAt:    review.UpdatedAt,
		DeletedAt:    &review.DeletedAt.Time,
		ProductID:    review.ProductID,
		Rating:       review.Rating,
		Deskripsi:    review.Deskripsi,
		Products:     ProductModelToEntity(review.Products),
		ImagesReview: reviewImages,
	}
}

func ReviewImageModelEntity(reviewimage ReviewImages) ReviewImagesEntity {
	return ReviewImagesEntity{
		Id:        reviewimage.ID,
		CreatedAt: reviewimage.CreatedAt,
		UpdatedAt: reviewimage.UpdatedAt,
		DeletedAt: &reviewimage.DeletedAt.Time,
		ReviewID:  reviewimage.ReviewID,
		Reviews:   ReviewModelEntity(reviewimage.Reviews),
		Link:      reviewimage.Link,
	}
}
