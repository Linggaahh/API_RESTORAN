package main

import (
	"api_resto/config"
	"api_resto/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Konek()

	r := gin.Default()

	//user routes
	r.GET("/users", controller.GetAllUsers)
	r.GET("/users/:id", controller.GetUserByID)
	r.POST("/register", controller.CreateUser)
	r.PUT("/users/:id", controller.UpdateUser)
	r.DELETE("/users/:id", controller.DeleteUser)
	r.POST("/login", controller.Login)

	//produk routes
	r.GET("/produks", controller.GetAllProduk)
	r.GET("/produks/:id", controller.GetProdukByID)
	r.POST("/produks", controller.CreateProduk)
	r.PUT("/produks/:id", controller.UpdateProduk)
	r.DELETE("/produks/:id", controller.DeleteProduk)
	
	//pesanan routes
	r.GET("/pesanans", controller.GetAllPesanan)
	r.GET("/pesanans/:id", controller.GetPesananByID)
	r.POST("/pesanans", controller.CreatePesanan)
	r.PUT("/pesanans/:id", controller.UpdatePesanan)
	r.DELETE("/pesanans/:id", controller.DeletePesanan)

	// Detail Pesanan routes
	r.GET("/detail-pesanan", controller.GetAllDetailPesanan)
	r.GET("/detail-pesanan/:id", controller.GetDetailPesananByID)
	r.POST("/detail-pesanan", controller.CreateDetailPesanan)
	r.PUT("/detail-pesanan/:id", controller.UpdateDetailPesanan)
	r.DELETE("/detail-pesanan/:id", controller.DeleteDetailPesanan)


	// Cart routes
	r.GET("/carts", controller.GetAllCart)
	r.POST("/carts", controller.CreateCart)
	r.PUT("/carts/:id", controller.UpdateCart)
	r.DELETE("/carts/:id", controller.DeleteCart)


	
	r.Run("0.0.0.0:8000")

}
