package controller

import (
	"net/http"
	"api_resto/config"
	"api_resto/model"

	"github.com/gin-gonic/gin"
)

// GET ALL
func GetAllDetailPesanan(c *gin.Context) {
	rows, err := config.DB.Query(`
		SELECT detail_id, pesanan_id, produk_id, jumlah_order, subtotal
		FROM detail_pesanan`)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}
	defer rows.Close()

	var details []model.DetailPesanan

	for rows.Next() {
		var d model.DetailPesanan
		if err := rows.Scan(
			&d.DetailID,
			&d.PesananID,
			&d.ProdukID,
			&d.JumlahOrder,
			&d.Subtotal,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
			return
		}
		details = append(details, d)
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "data": details})
}

// GET BY ID
func GetDetailPesananByID(c *gin.Context) {
	id := c.Param("id")

	var d model.DetailPesanan
	err := config.DB.QueryRow(`
		SELECT detail_id, pesanan_id, produk_id, jumlah_order, subtotal
		FROM detail_pesanan WHERE detail_id=?`,
		id,
	).Scan(
		&d.DetailID,
		&d.PesananID,
		&d.ProdukID,
		&d.JumlahOrder,
		&d.Subtotal,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": 0, "error": "Detail pesanan tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "data": d})
}

// CREATE
func CreateDetailPesanan(c *gin.Context) {
	var input model.DetailPesanan

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}

	_, err := config.DB.Exec(`
		INSERT INTO detail_pesanan (pesanan_id, produk_id, jumlah_order, subtotal)
		VALUES (?, ?, ?, ?)`,
		input.PesananID,
		input.ProdukID,
		input.JumlahOrder,
		input.Subtotal,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "msg": "Detail pesanan berhasil dibuat"})
}

// UPDATE
func UpdateDetailPesanan(c *gin.Context) {
	id := c.Param("id")

	var input model.DetailPesanan
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}

	_, err := config.DB.Exec(`
		UPDATE detail_pesanan 
		SET pesanan_id=?, produk_id=?, jumlah_order=?, subtotal=?
		WHERE detail_id=?`,
		input.PesananID,
		input.ProdukID,
		input.JumlahOrder,
		input.Subtotal,
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "msg": "Detail pesanan berhasil diupdate"})
}

// DELETE
func DeleteDetailPesanan(c *gin.Context) {
	id := c.Param("id")

	_, err := config.DB.Exec("DELETE FROM detail_pesanan WHERE detail_id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "msg": "Detail pesanan berhasil dihapus"})
}
