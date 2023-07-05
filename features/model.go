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
	UserID       uint          `gorm:"column:user_id"`
	Users        User          `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Transactions []Transaction `gorm:"foreignKey:ProductID"`
	Images       []Image       `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Payment struct {
	gorm.Model
	TransactionID uint        `gorm:"column:transaction_id"`
	Transactions  Transaction `gorm:"foreignKey:TransactionID"`
	Status        string      `gorm:"column:status"`
	Bank          string      `gorm:"column:bank"`
	VA            string      `gorm:"column:va"`
	OrderID       string      `gorm:"column:order_id"`
}

type Transaction struct {
	gorm.Model
	ProductID    uint     `gorm:"column:product_id"`
	Products     Product  `gorm:"foreignKey:ProductID"`
	UserID       uint     `gorm:"column:user_id"`
	Users        User     `gorm:"foreignKey:UserID"`
	Status       string   `gorm:"column:status;not null"`
	TotalHarga   string   `gorm:"column:total_harga;not null"`
	JumlahBarang int      `gorm:"column:jumlah_barang;not null"`
}

type Image struct {
	gorm.Model
	ProductID uint   `gorm:"column:product_id"`
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Link      string `gorm:"column:link"`
	Nama      string `gorm:"column:nama_images"`
}
