package learning_data_users

import (
	"errors"
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

func GetUserByID(id string) (*User, error) {
	for i, t := range testUsers {
		if t.ID == id {
			return &testUsers[i], nil
		}
	}

	return nil, errors.New("not found")
}

func RemoveUserByID(id string) (*User, error) {
	for i, t := range testUsers {
		if t.ID == id {
			removedUser := testUsers[i]
			testUsers = append(testUsers[:i], testUsers[i+1:]...)
			return &removedUser, nil
		}
	}

	return nil, errors.New("not found")
}

// SECTION API REST Methods

func GetUsers(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, testUsers)
}

func AddUser(context *gin.Context) {
	var newUser User

	if err := context.BindJSON(&newUser); err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": true})
		return
	}

	testUsers = append(testUsers, newUser)

	context.IndentedJSON(http.StatusCreated, newUser)
}

func GetUser(context *gin.Context) {
	id := context.Param("id")
	user, err := GetUserByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": true})
		return
	}

	context.IndentedJSON(http.StatusOK, user)
}

func RemoveUser(context *gin.Context) {
	id := context.Param("id")
	user, err := RemoveUserByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": true})
		return
	}

	context.IndentedJSON(http.StatusOK, user)
}

func DefineRouter(router gin.IRouter) {
	router.GET("/users", GetUsers)
	router.GET("/users/:id", GetUser)
	router.POST("/users", AddUser)
	router.DELETE("/users/:id", RemoveUser)
}
