package features

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama         string        `gorm:"column:nama;not null"`
	NoTlp        string        `gorm:"column:no_tlp;unique;not null"`
	Email        string        `gorm:"column:email;unique;not null"`
	Password     string        `gorm:"column:password;not null"`
	Alamat       string        `gorm:"column:alamat;not null"`
	Products     []Product     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Transactions []Transaction `gorm:"foreignKey:UserID"`
}

type Product struct {
	gorm.Model
	Nama         string        `gorm:"column:nama"`
	Harga        string        `gorm:"column:harga"`
	Deskripsi    string        `gorm:"column:deskripsi"`
	Stok         int           `gorm:"column:stok"`
	Ratings      int		   `gorm:"column:ratings"`
	UserID       uint          `gorm:"column:user_id"`
	Users        User          `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Transactions []Transaction `gorm:"foreignKey:ProductID"`
	Images       []Image       `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Reviews		 []Review      `gorm:"foreignKey:ProductID"`
}

type Review struct{
	gorm.Model
	ProductID 		uint 			`gorm:"column:product_id"`
	Rating 			int				`gorm:"column:rating"`
	Deskripsi 		string			`gorm:"column:deskripsi"`
	Products		Product 		`gorm:"foreignKey:ProductID"`
	ImagesReview 	[]ReviewImages  `gorm:"foreignKey:ReviewID"`

}

type ReviewImages struct{
	gorm.Model
	ReviewID 		uint		`gorm:"column:review_id"`
	Reviews 		Review		`gorm:"foreignKey:ReviewID"`
	Link 			string		`gorm:"column:link"`
	
}

type Payment struct {
	gorm.Model
	TransactionPaymentID 		uint 			`gorm:"column:transaction_payment_id"`
	TransactionPayments 		TransactionPayment 	`gorm:"foreignKey:TransactionPaymentID"`
	Status        			string      		`gorm:"column:status"`
	Bank          			string      		`gorm:"column:bank"`
	VA            			string      		`gorm:"column:va"`
	OrderID       			string      		`gorm:"column:order_id"`
}


type Transaction struct {
	gorm.Model
	ProductID    			uint     `gorm:"column:product_id"`
	Products     			Product  `gorm:"foreignKey:ProductID"`
	TransactionPaymentID 	uint 	 `gorm:"column:transaction_payment_id"`
	TransactionPayment 		TransactionPayment `gorm:"foreignKey:TransactionPaymentID"`
	UserID       			uint     `gorm:"column:user_id"`
	Users        			User     `gorm:"foreignKey:UserID"`
	Status       			string   `gorm:"column:status;not null"`
	TotalHarga   			string   `gorm:"column:total_harga;not null"`
	JumlahBarang 			int      `gorm:"column:jumlah_barang;not null"`
}

type TransactionPayment struct {
	gorm.Model
	Transactions  			[]Transaction 		`gorm:"foreignKey:TransactionPaymentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` 
	Payments				[]Payment			`gorm:"foreignKey:TransactionPaymentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	HargaTotal    			string 	    		`gorm:"column:harga_total"`
}
type Image struct {
	gorm.Model
	ProductID uint   `gorm:"column:product_id"`
	Products   Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Link      string `gorm:"column:link"`
	Nama      string `gorm:"column:nama_images"`
}
