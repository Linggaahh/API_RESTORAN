package controller

import (
	"net/http"
	"api_resto/config"
	"api_resto/model"

	"github.com/gin-gonic/gin"
)

func GetAllProduk(c *gin.Context) {
	rows, err := config.DB.Query("SELECT produk_id, nama_produk, deskripsi, stock, price, image_url FROM produk")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}
	defer rows.Close()

	var produks []model.Produk

	for rows.Next() {
		var produk model.Produk
		err := rows.Scan(
			&produk.ProdukID,
			&produk.NamaProduk,
			&produk.Deskripsi,
			&produk.Stock,
			&produk.Price,
			&produk.ImageURL,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
			return
		}
		produks = append(produks, produk)
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "data": produks})
}

func GetProdukByID(c *gin.Context) {
	id := c.Param("id")

	var produk model.Produk
	err := config.DB.QueryRow(
		"SELECT produk_id, nama_produk, deskripsi, stock, price, image_url FROM produk WHERE produk_id=?",
		id,
	).Scan(
		&produk.ProdukID,
		&produk.NamaProduk,
		&produk.Deskripsi,
		&produk.Stock,
		&produk.Price,
		&produk.ImageURL,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": 0, "error": "Produk tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "data": produk})
}

func CreateProduk(c *gin.Context) {
	var produk model.Produk

	if err := c.ShouldBindJSON(&produk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}

	_, err := config.DB.Exec(
		"INSERT INTO produk(nama_produk, deskripsi, stock, price, image_url) VALUES(?, ?, ?, ?, ?)",
		produk.NamaProduk,
		produk.Deskripsi,
		produk.Stock,
		produk.Price,
		produk.ImageURL,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "msg": "Produk berhasil ditambahkan"})
}

func UpdateProduk(c *gin.Context) {
	id := c.Param("id")
	var produk model.Produk

	if err := c.ShouldBindJSON(&produk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}

	_, err := config.DB.Exec(
		"UPDATE produk SET nama_produk=?, deskripsi=?, stock=?, price=?, image_url=? WHERE produk_id=?",
		produk.NamaProduk,
		produk.Deskripsi,
		produk.Stock,
		produk.Price,
		produk.ImageURL,
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "msg": "Produk berhasil diupdate"})
}

func DeleteProduk(c *gin.Context) {
	id := c.Param("id")

	_, err := config.DB.Exec("DELETE FROM produk WHERE produk_id=?", id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "msg": "Produk berhasil dihapus"})
}
