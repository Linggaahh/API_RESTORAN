package controller

import (
	"net/http"
	"api_resto/config"
	"api_resto/model"

	"github.com/gin-gonic/gin"
)

func GetAllPesanan(c *gin.Context) {
	rows, err := config.DB.Query("SELECT pesanan_id, user_id, pesanan_date, status, note FROM pesanan")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}
	defer rows.Close()

	var pesanans []model.Pesanan

	for rows.Next() {
		var pesanan model.Pesanan
		err := rows.Scan(
			&pesanan.PesananID,
			&pesanan.UserID,
			&pesanan.PesananDate,
			&pesanan.Status,
			&pesanan.Note,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
			return
		}
		pesanans = append(pesanans, pesanan)
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "data": pesanans})
}

func GetPesananByID(c *gin.Context) {
	id := c.Param("id")

	var pesanan model.Pesanan

	err := config.DB.QueryRow(
		"SELECT pesanan_id, user_id, pesanan_date, status, note FROM pesanan WHERE pesanan_id=?",
		id,
	).Scan(
		&pesanan.PesananID,
		&pesanan.UserID,
		&pesanan.PesananDate,
		&pesanan.Status,
		&pesanan.Note,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": 0, "error": "Pesanan tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "data": pesanan})
}

func CreatePesanan(c *gin.Context) {
	var pesanan model.Pesanan

	if err := c.ShouldBindJSON(&pesanan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}

	_, err := config.DB.Exec(
		"INSERT INTO pesanan(user_id, pesanan_date, status, note) VALUES(?, ?, ?, ?)",
		pesanan.UserID,
		pesanan.PesananDate,
		pesanan.Status,
		pesanan.Note,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "msg": "Pesanan berhasil dibuat"})
}

func UpdatePesanan(c *gin.Context) {
	id := c.Param("id")

	var pesanan model.Pesanan

	if err := c.ShouldBindJSON(&pesanan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}

	_, err := config.DB.Exec(
		"UPDATE pesanan SET user_id=?, pesanan_date=?, status=?, note=? WHERE pesanan_id=?",
		pesanan.UserID,
		pesanan.PesananDate,
		pesanan.Status,
		pesanan.Note,
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "msg": "Pesanan berhasil diupdate"})
}

func DeletePesanan(c *gin.Context) {
	id := c.Param("id")

	_, err := config.DB.Exec("DELETE FROM pesanan WHERE pesanan_id=?", id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "msg": "Pesanan berhasil dihapus"})
}
