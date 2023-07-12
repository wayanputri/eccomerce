package features

import (
	"time"
)

type UserEntity struct {
	Id           uint      		     `json:"user_id,omitempty" form:"user_id"`
	CreatedAt    time.Time  	     `json:"created_at,omitempty"`
	UpdatedAt    time.Time  	     `json:"updated_at,omitempty"`
	DeletedAt    time.Time   	     `json:"deleted_at,omitempty"`
	Nama         string     	     `json:"nama,omitempty" form:"nama"`
	NoTlp        string       		 `json:"no_tlp,omitempty" form:"no_tlp"`
	Email        string       		 `json:"email,omitempty" form:"email"`
	Password     string        		 `json:"password,omitempty" form:"password"`
	Alamat       string        		 `json:"alamat,omitempty" form:"alamat"`
	Transactions []TransactionEntity `json:"transactions,omitempty"`
	Product		 []ProductEntity	 `json:"products,omitempty"`
}

type LoginUser struct {
	Email        string       		 `json:"email,omitempty" form:"email" validate:"required,email"`
	Password     string        		 `json:"password,omitempty" form:"password" validate:"required"`
}

type PaymentEntity struct {
	Id          			uint         		  		`json:"payment_id,omitempty" form:"payment_id"`
	CreatedAt    			time.Time    		  		`json:"created_at,omitempty"`
	UpdatedAt   		 	time.Time    		  		`json:"updated_at,omitempty"`
	DeletedAt    			time.Time    		  		`json:"deleted_at,omitempty"`
	TransactionID 			uint        		  		`json:"transaction_id,omitempty" form:"transaction_id"`
	Transactions  			TransactionEntity   	 	`json:"transactions,omitempty"`
	TransactionPaymentID 	uint 	 				 	`json:"transaction_payment_id"`
	TransactionPayment 		TransactionPaymentEntity 	`json:"transaction_payment"`
	Status        			string      		  		`json:"status,omitempty" form:"status"`
	Bank          			string      		  		`json:"bank,omitempty" form:"bank"`
	VA            			string      		 		`json:"va,omitempty" form:"va"`
	OrderID		  			string			  			`json:"order_id,omitempty" form:"order_id"`
}

type TransactionEntity struct {
	Id           			uint      	      		`json:"transaction_id,omitempty" form:"transaction_id"`
	CreatedAt    			time.Time 	      		`json:"created_at,omitempty"`
	UpdatedAt   			time.Time  	  	  		`json:"updated_at,omitempty"`
	DeletedAt    			time.Time    	  		`json:"deleted_at,omitempty"`
	TransactionPaymentID 	uint 	 		  		`json:"transaction_payment_id"`
	TransactionPayment 		TransactionPaymentEntity `json:"transaction_payment"`
	ProductID 	 			uint    	  	  		`json:"product_id,omitempty" form:"product_id"`
	Products 	 			ProductEntity 	  		`json:"products,omitempty"`
	UserID   	 			uint    		  		`json:"user_id,omitempty" form:"user_id"`
	Users    	 			UserEntity   	  		`json:"users,omitempty"`
	Status  	 			string 		  	  		`json:"status,omitempty" form:"status"`
	TotalHarga	 			string			  		`json:"total_harga,omitempty" form:"total_harga"`
	JumlahBarang 			int			 	  		`json:"jumlah_barang,omitempty" form:"jumlah_barang"`
}

type TransactionPaymentEntity struct {
	Id           			uint      	      	`json:"transaction_payment_id,omitempty" form:"transaction_payment_id"`
	CreatedAt    			time.Time 	      	`json:"created_at,omitempty"`
	UpdatedAt   			time.Time  	  	  	`json:"updated_at,omitempty"`
	DeletedAt    			time.Time    	  	`json:"deleted_at,omitempty"`
	Transactions  			[]TransactionEntity `json:"transactions,omitempty"`
	Payments     			[]PaymentEntity   	`json:"payments,omitempty"`
	HargaTotal    			string 	    	  	`json:"harga_total" form:"harga_total"`
}

type ProductEntity struct {
	Id          uint                `json:"product_id,omitempty" form:"product_id"`
	CreatedAt   time.Time           `json:"created_at,omitempty"`
	UpdatedAt   time.Time           `json:"updated_at,omitempty"`
	DeletedAt   *time.Time          `json:"deleted_at,omitempty"`
	Nama        string              `json:"nama,omitempty" form:"nama"`
	Harga       string              `json:"harga,omitempty" form:"harga"`
	Ratings     int					`json:"ratings,omitempty" form:"ratings"`
	Deskripsi   string              `json:"deskripsi,omitempty" form:"deskripsi"`
	Stok        int                 `json:"stok,omitempty" form:"stok"`
	UserId      uint                `json:"user_id" form:"user_id"`
	Users       UserEntity          `json:"users,omitempty"`
	Transactions []TransactionEntity `json:"transactions,omitempty"`
	Images      []ImageEntity       `json:"images,omitempty"`
	Reviews		[]ReviewEntity      `json:"reviews,omitempty"`
}

type ReviewEntity struct{
	Id          	uint                	`json:"review_id,omitempty" form:"review_id"`
	CreatedAt   	time.Time           	`json:"created_at,omitempty"`
	UpdatedAt   	time.Time           	`json:"updated_at,omitempty"`
	DeletedAt   	*time.Time          	`json:"deleted_at,omitempty"`
	ProductID 		uint 					`json:"product_id,omitempty" form:"product_id"`
	Rating 			int						`json:"rating,omitempty" form:"rating"`
	Deskripsi 		string					`json:"deskripsi,omitempty" form:"deskripsi"`
	Products		ProductEntity 			`json:"products,omitempty"`
	ImagesReview 	[]ReviewImagesEntity  	`json:"imagesreviews,omitempty"`

}

type ReviewImagesEntity struct{
	Id          	uint                `json:"reviewimages_id,omitempty" form:"reviewimages_id"`
	CreatedAt   	time.Time           `json:"created_at,omitempty"`
	UpdatedAt   	time.Time           `json:"updated_at,omitempty"`
	DeletedAt   	*time.Time          `json:"deleted_at,omitempty"`
	ReviewID 		uint				`json:"review_id,omitempty" form:"review_id"`
	Reviews 		ReviewEntity		`json:"reviews,omitempty"`
	Link 			string				`json:"link,omitempty" form:"link"`
	
}

type ImageEntity struct {
	Id          uint          `json:"image_id,omitempty" form:"image_id"`
	CreatedAt    time.Time     `json:"created_at,omitempty"`
	UpdatedAt    time.Time     `json:"updated_at,omitempty"`
	DeletedAt    time.Time     `json:"deleted_at,omitempty"`
	ProductID 	 uint    	   `json:"product_id,omitempty" form:"product_id"`
	Products  	 ProductEntity `json:"products,omitempty"`
	Link      	 string 	   `json:"link,omitempty" form:"link"`
	Nama      	 string 	   `json:"nama_images,omitempty" form:"nama_images"`
}