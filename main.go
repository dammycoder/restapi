package main 

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)

var schools []strings

func main (){
	r := gin.Default()
	r.GET("/",schools)
	r.POST("/school",newschools)
	r.Run(":8080")
}

func newschools(c *gin.context){
	school := c.PostForm("school")
	schools := append(schools, school)
	c.String(http.StatusCreated, c.FullPath()+"/"+strconv.Itoa(len(todos)-1))
}
