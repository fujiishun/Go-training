package controllers

import (
	"backend/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetUsers(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "データ取得に失敗しました"})
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "データの読み込みに失敗しました"})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}
