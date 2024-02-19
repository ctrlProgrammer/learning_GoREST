package main

import (
	learning_data_users "learning/data/src"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	learning_data_users.DefineRouter(router)

	router.Run("localhost:9090")
}
