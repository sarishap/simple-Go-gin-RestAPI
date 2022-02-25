package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func getMessage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func postMessage(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"message": string(value),
	})
}
func QueryString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}
func ParamsString(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func main() {
	r := gin.Default()
	r.GET("/getMessage", getMessage)
	r.POST("/postMessage", postMessage)
	r.GET("/query", QueryString)            // ?name=xxx&age=xxx
	r.GET("/path/:name/:age", ParamsString) //path/xxx/xxx
	r.Run()
}
