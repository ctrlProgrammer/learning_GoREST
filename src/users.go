package learning_data_users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Address  string `json:"address"`
	Created  int32  `json:"created"`
}

var testUsers = []User{
	{ID: "1", Username: "TestUser_1", Address: "0x1", Created: 1708346339},
	{ID: "2", Username: "TestUser_2", Address: "0x2", Created: 1708346339},
	{ID: "3", Username: "TestUser_3", Address: "0x3", Created: 1708346339},
	{ID: "4", Username: "TestUser_4", Address: "0x4", Created: 1708346339},
}

func GetUsers(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, testUsers)
}

func AddUser(context *gin.Context) {
	var newUser User

	if err := context.BindJSON(&newUser); err != nil {
		return
	}

	testUsers = append(testUsers, newUser)
}

func DefineRouter(router gin.IRouter) {
	router.GET("/users", GetUsers)
}
