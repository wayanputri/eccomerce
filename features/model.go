package features

import "gorm.io/gorm"

type Product struct{
	gorm.Model
	Nama 		 string			`gorm:"column:nama"`
	Harga 		 string			`gorm:"column:harga"`
	Deskripsi 	 string			`gorm:"column:deskripsi"`
	Stok 		 int			`gorm:"column:stok"`
	UserID 		 uint			`gorm:"column:user_id"`
	Users		 User			`gorm:"foreignKey:UserID"`
	Transactions []Transaction	`gorm:"foreignKey:ProductID"`
	Image		 []Image		`gorm:"foreignKey:ProductID"`	
}

type User struct{
	gorm.Model
	Nama 		 string			`gorm:"column:nama;not nul"`
	NoTlp 		 string			`gorm:"column:no_tlp;unique;not nul"`
	Email 		 string			`gorm:"column:email;unique;not nul"`
	Password 	 string			`gorm:"column:password;not nul"`
	Alamat 		 string			`gorm:"column:alamat;not nul"`
	Products	 []Product		`gorm:"foreignKey:UserID"`
	Transactions []Transaction	`gorm:"foreignKey:UserID"`
}

type Payment struct{
	gorm.Model
	TransactionID 	uint		  `gorm:"column:transaction_id"`
	Transactions	Transaction	  `gorm:"foreignKey:TransactionID"`
	Status 			string		  `gorm:"column:status"`
	Bank 			string		  `gorm:"column:bank"`
	VA		 		string		  `gorm:"column:va"`
	OrderID			string		  `gorm:"column:order_id"`
}

type Transaction struct{
	gorm.Model
	ProductID 	 uint			`gorm:"column:product_id"`
	Products	 Product		`gorm:"foreignKey:ProductID"`
	UserID 		 uint			`gorm:"column:user_id"`
	Users		 User			`gorm:"foreignKey:UserID"`
	Status 		 string			`gorm:"column:status;not nul"`
	TotalHarga	 string			`gorm:"column:total_harga;not nul"`
	JumlahBarang int			`gorm:"column:jumlah_barang;not nul"`
}


type Image struct{
	gorm.Model
	ProductID 	uint	`gorm:"column:product_id"`
	Products	Product `gorm:"foreignKey:ProductID"`		
	Link 		string	`gorm:"column:link"`
	Nama 		string	`gorm:"column:nama"`
}