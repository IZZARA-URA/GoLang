package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//User Structure
type User struct {
	Username string `json : "username"`
	Password string `json : "password"`
}

//Create Function
func GetHellworld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
}

func GetUsername(c *gin.Context) { //string Qry
	username := c.Query("username")
	password := c.Query("password")
	c.JSON(http.StatusOK, gin.H{"username": username, "password": password})
}

func Post(c *gin.Context) {
	var u User
	c.BindJSON(&u)
	c.JSON(http.StatusOK, gin.H{"user": u})
}

func PustPathParameter(c *gin.Context) {
	username := c.Param("username")
	password := c.Param("password")
	c.JSON(http.StatusOK, gin.H{"messege": username, "password": password})
}

func DeleteQuery(c *gin.Context) {
	username := c.Query("username")
	c.JSON(http.StatusOK, gin.H{"username": username})
}

func main() {
	r := gin.Default()

	r.GET("/", GetHellworld)
	r.GET("/Username", GetUsername)
	r.POST("/", Post)
	r.PUT("/user/:username/:password", PustPathParameter)
	r.DELETE("/username", DeleteQuery)

	r.Run()
	fmt.Print("125")
}
