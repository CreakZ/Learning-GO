package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 1. Вывод данных о пользователе
// 2. Создание пользователя
// 3. Изменение имени
// 4. Изменение фамилии
// 5. Изменение пароля

type User struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Password string `json:"password"`
}

func GetUserInfo(user User) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, user)
	}
}

func main() {
	router := gin.Default()

	user := User{"Georgiy", "Sladkopykhov", "ilovefurryporn"}

	router.GET("/user_info", GetUserInfo(user))

	router.POST("/user_creation", func(c *gin.Context) {
		var newUser User

		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user = newUser

		c.JSON(200, gin.H{"post": "successful"})
	})

	runErr := router.Run()

	if runErr != nil {
		log.Fatal("Error: ", runErr)
	}
}
