package features

func UserEntityToModel(user UserEntity) User {

	var transactions []Transaction
	for _, transaction := range user.Transactions {
		transactions = append(transactions, TransactionEntityToModel(transaction))
	}

	var products []Product
	for _, product := range user.Product {
		products = append(products, ProductEntityToModel(product))
	}

	return User{
		Nama:         user.Nama,
		NoTlp:        user.NoTlp,
		Email:        user.Email,
		Password:     user.Password,
		Alamat:       user.Alamat,
		JenisKelamin: user.JenisKelamin,
		Role:         user.Role,
		File:         user.File,
		Transactions: transactions,
		Products:     products,
	}
}

func PaymentEntityToModel(payment PaymentEntity) Payment {
	return Payment{
		Status:  payment.Status,
		Bank:    payment.Bank,
		VA:      payment.VA,
		OrderID: payment.OrderID,
	}
}

func TransactionPaymentEntityToModel(transactionPayment TransactionPaymentEntity) TransactionPayment {
	var transaksi []Transaction
	for _, transaction := range transactionPayment.Transactions {
		transaksi = append(transaksi, TransactionEntityToModel(transaction))
	}
	var payments []Payment
	for _, payment := range transactionPayment.Payments {
		payments = append(payments, PaymentEntityToModel(payment))
	}
	return TransactionPayment{
		Transactions: transaksi,
		Payments:     payments,
		HargaTotal:   transactionPayment.HargaTotal,
	}
}

func TransactionEntityToModel(transaction TransactionEntity) Transaction {
	return Transaction{
		ProductID:    transaction.ProductID,
		Products:     ProductEntityToModel(transaction.Products),
		UserID:       transaction.UserID,
		Users:        UserEntityToModel(transaction.Users),
		Status:       transaction.Status,
		TotalHarga:   transaction.TotalHarga,
		JumlahBarang: transaction.JumlahBarang,
	}
}

func ProductEntityToModel(product ProductEntity) Product {
	var transactions []Transaction
	for _, transaction := range product.Transactions {
		transactions = append(transactions, TransactionEntityToModel(transaction))
	}
	var imagesModel []Image
	for _, image := range product.Images {
		imagesModel = append(imagesModel, ImageEntityToModel(image))
	}
	var reviewsModel []Review
	for _, review := range product.Reviews {
		reviewsModel = append(reviewsModel, ReviewEntityToModel(review))
	}
	return Product{
		Nama:         product.Nama,
		Harga:        product.Harga,
		Deskripsi:    product.Deskripsi,
		Stok:         product.Stok,
		UserID:       product.UserId,
		Users:        UserEntityToModel(product.Users),
		Transactions: transactions,
		Images:       imagesModel,
		Ratings:      product.Ratings,
		Reviews:      reviewsModel,
	}
}

func ImageEntityToModel(image ImageEntity) Image {
	return Image{
		ProductID: image.ProductID,
		Products:  ProductEntityToModel(image.Products),
		Link:      image.Link,
		Nama:      image.Nama,
	}
}

func ReviewEntityToModel(review ReviewEntity) Review {
	var reviewImages []ReviewImages
	for _, reviewImage := range review.ImagesReview {
		reviewImages = append(reviewImages, ReviewImageEntityToModel(reviewImage))
	}
	return Review{
		ProductID:    review.ProductID,
		Rating:       review.Rating,
		Deskripsi:    review.Deskripsi,
		Products:     ProductEntityToModel(review.Products),
		ImagesReview: reviewImages,
	}
}

func ReviewImageEntityToModel(reviewimage ReviewImagesEntity) ReviewImages {
	return ReviewImages{
		ReviewID: reviewimage.ReviewID,
		Reviews:  ReviewEntityToModel(reviewimage.Reviews),
		Link:     reviewimage.Link,
	}
}
