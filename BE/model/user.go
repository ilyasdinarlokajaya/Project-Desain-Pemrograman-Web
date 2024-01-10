package model

// tags
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Datauser struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Image    string `json:"image"`
}

type Databanner struct {
	ID       		uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	BannerName     	string `json:"bannername"`
	BroadcastDay 	string `json:"broadcastday"`
	ValidFrom    	string `json:"validfrom"`
	ValidUntil     	string `json:"validuntil"`
	BannerStatus    string `json:"bannerstatus"`
	Image    		string `json:"image"`
}
type Datanotification struct {
	ID              uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Title           string `json:"title"`
	Konten          string `json:"konten"`
	Broadcast_Day 	string `json:"broadcast_day"`
}
type Dataproduct struct {
	ID              	uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductCode 		string `json:"product_code"`
	ProductName     	string `json:"product_name"`
	Purchase    		int 	`json:"purchase"`
	ProductDescription 	string `json:"productdescription"`
    Unit 				int 	`json:"unit"`
	MinPurchase 		int 	`json:"minpurchase"`
	MaxPurchase 		int 	`json:"maxpurchase"`
	Category 			string `json:"category"`
	Image    			string `json:"image"`
}
type Dataplannogram struct {
	ID       			uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	PlannogramName  	string `json:"plannogramname"`
	Image    			string `json:"image"`
	BroadcastDay 		string `json:"broadcastday"`
	ValidFrom    		string `json:"validfrom"`
	ValidUntil     		string `json:"validuntil"`
	ListProduct         string `json:"listproduct"`
}

type Datavoucher struct {
	ID       		uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	VoucherCode 	string `json:"vouchercode"`
	VoucherName 	string `json:"vouchername"`
	Description 	string `json:"description"`
	Type 			string `json:"type"`
	MinTransaction 	int		`json:"mintransaction"`
	Amount 			int		`json:"amount"`
	ValidFrom    	string `json:"validfrom"`
	ValidUntil     	string `json:"validuntil"`
	Image    		string `json:"image"`
}

type Dataarticle struct {
	ID       			uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	ArticleName  		string `json:"judul_artikel"`
	Konten 				string `json:"konten"`
	Image    			string `json:"image"`
	BroadcastDay 		string `json:"jadwal_broadcast"`
	ValidFrom    		string `json:"valid_from"`
	ValidUntil     		string `json:"valid_until"`
}

type Datadeliverymethod struct {
	ID       			uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	DeliveryName  		string `json:"delivery_name"`
	Delivery_Fee 		string `json:"delivery_fee"`
	Estimation	 		string `json:"estimation"`
	Image    			string `json:"image"`
}