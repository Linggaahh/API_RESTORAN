package model

type User struct {
	UserID   int    `json:"user_id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
	Telp     string `json:"telp" binding:"required"`
	Role     string `json:"role" binding:"required"`
	Image    string `json:"image"`
}

type Produk struct {
	ProdukID   int    `json:"produk_id"`
	NamaProduk string `json:"nama_produk" binding:"required"`
	Deskripsi  string `json:"deskripsi" binding:"required"`
	Stock      int    `json:"stock" binding:"required"`
	Price      int    `json:"price" binding:"required"`
	ImageURL   string `json:"image_url"`
}

type Pesanan struct {
	PesananID   int    `json:"pesanan_id"`
	UserID      int    `json:"user_id" binding:"required"`
	PesananDate string `json:"pesanan_date" binding:"required"`
	Status      string `json:"status" binding:"required"`
	Note        string `json:"note" binding:"required"`
}

type DetailPesanan struct {
	DetailID    int `json:"detail_id"`
	PesananID   int `json:"pesanan_id" binding:"required"`
	ProdukID    int `json:"produk_id" binding:"required"`
	JumlahOrder int `json:"jumlah_order" binding:"required"`
	Subtotal    int `json:"subtotal" binding:"required"`
}

type Cart struct {
	Cart_id   int `json:"cart_id"`
	User_id   int `json:"user_id" binding:"required"`
	Produk_id int `json:"produk_id" binding:"required"`
	Quantity  int `json:"quantity" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
