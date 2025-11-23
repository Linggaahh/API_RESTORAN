package controller

import (
	"net/http"
	"api_resto/config"
	"api_resto/model"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	rows, err := config.DB.Query(`
		SELECT user_id, email, password, username, telp, role, image 
		FROM user`)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}
	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var u model.User
		err := rows.Scan(
			&u.UserID,
			&u.Email,
			&u.Password,
			&u.Username,
			&u.Telp,
			&u.Role,
			&u.Image,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
			return
		}
		users = append(users, u)
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "data": users})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	var u model.User
	err := config.DB.QueryRow(`
		SELECT user_id, email, password, username, telp, role, image 
		FROM user WHERE user_id = ?`,
		id,
	).Scan(
		&u.UserID,
		&u.Email,
		&u.Password,
		&u.Username,
		&u.Telp,
		&u.Role,
		&u.Image,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": 0, "error": "User tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "data": u})
}

func CreateUser(c *gin.Context) {
	var u model.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}

	_, err := config.DB.Exec(`
		INSERT INTO user (email, password, username, telp, role, image) 
		VALUES (?, ?, ?, ?, ?, ?)`,
		u.Email, u.Password, u.Username, u.Telp, u.Role, u.Image,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"msg":    "User berhasil ditambahkan",
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var u model.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}

	_, err := config.DB.Exec(`
		UPDATE user 
		SET email=?, password=?, username=?, telp=?, role=?, image=?
		WHERE user_id=?`,
		u.Email, u.Password, u.Username, u.Telp, u.Role, u.Image, id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "msg": "User berhasil diupdate"})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	_, err := config.DB.Exec("DELETE FROM user WHERE user_id = ?", id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "msg": "User berhasil dihapus"})
}
// ============================ LOGIN ==============================
func Login(c *gin.Context) {
	var req model.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 0,
			"error":  "Email dan Password wajib diisi",
		})
		return
	}

	var u model.User
	err := config.DB.QueryRow(`
		SELECT user_id, email, password, username, telp, role, image
		FROM user WHERE email = ? AND password = ?`,
		req.Username, req.Password,
	).Scan(
		&u.UserID,
		&u.Email,
		&u.Password,
		&u.Username,
		&u.Telp,
		&u.Role,
		&u.Image,
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 0,
			"error":  "Email atau Password salah",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"msg":    "Login berhasil",
		"user":   u,
	})
}
