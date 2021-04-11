package app

import (
	"github.com/rodrigoagostin/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
}
