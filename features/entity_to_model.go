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
		Transactions: transactions,
		Products:     products,
	}
}

func PaymentEntityToModel(payment PaymentEntity) Payment {
	return Payment{
		TransactionID: payment.TransactionID,
		Transactions:  TransactionEntityToModel(payment.Transactions),
		Status:        payment.Status,
		Bank:          payment.Bank,
		VA:            payment.VA,
		OrderID:       payment.OrderID,
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
	return Product{
		Nama:         product.Nama,
		Harga:        product.Harga,
		Deskripsi:    product.Deskripsi,
		Stok:         product.Stok,
		UserID:       product.UserId,
		Users:        UserEntityToModel(product.Users),
		Transactions: transactions,
		Images:       imagesModel,
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