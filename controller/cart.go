package controller

import (
	"net/http"
	"api_resto/config"
	"api_resto/model"
	"github.com/gin-gonic/gin"
)

func GetAllCart(c *gin.Context) {
	rows, err := config.DB.Query("SELECT cart_id, user_id, produk_id, quantity FROM cart")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}
	defer rows.Close()

	var carts []model.Cart
	for rows.Next() {
		var cart model.Cart
		err := rows.Scan(&cart.Cart_id, &cart.User_id, &cart.Produk_id, &cart.Quantity)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
			return
		}
		carts = append(carts, cart)
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "data": carts})
}

func CreateCart(c *gin.Context) {
	var cart model.Cart
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}

	_, err := config.DB.Exec(
		"INSERT INTO cart (user_id, produk_id, quantity) VALUES (?, ?, ?)",
		cart.User_id, cart.Produk_id, cart.Quantity,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "message": "Cart created"})
}

func UpdateCart(c *gin.Context) {
	id := c.Param("id")
	var cart model.Cart
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": err.Error()})
		return
	}

	_, err := config.DB.Exec(
		"UPDATE cart SET user_id=?, produk_id=?, quantity=? WHERE cart_id=?",
		cart.User_id, cart.Produk_id, cart.Quantity, id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "message": "Cart updated"})
}

func DeleteCart(c *gin.Context) {
	id := c.Param("id")

	_, err := config.DB.Exec("DELETE FROM cart WHERE cart_id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 1, "message": "Cart deleted"})
}
