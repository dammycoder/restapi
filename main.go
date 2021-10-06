package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type todolist struct {
	name string `json:"name"`
	id int  `json:"id"`
}


 type todoitem struct{
	 name string `json :"name"`
	 id int `json:"id"`
	 todolistid int `json:"id"`
	 isdone bool `json:"isdone"`
 }

var todolistcollection map [string] interface{}

func main() {
	r := gin.Default()

	r.GET("/todo-list", listTodoList)
	r.POST("/todo-list", createTodoList)
	r.DELETE("/todo-list/:todoListID", deleteTodoList)
	r.PATCH("/todo-list/:todoListID", patchTodoList)
	r.GET("/todo-list/:todoListID", showTodoList)
	r.GET("/todolist/:todolistID/:todoITEM",todoitems)
	r.POST("/todolist/:todolistID/:todoITEM",addtodoitem)
	r.PATCH("/:todoITEM",marktodoitem)
	r.Run("127.0.0.1:8080")
}

func listTodoList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": todolistcollection,
	})
}

func createTodoList(ctx *gin.Context) {
	data := todolist{}
	if  err := ctx.ShouldBindJSON(&data);if err != nil{
		ctx.JSON(200,err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"name" : data.name,
		"id":data.id,
	})
}

func deleteTodoList(ctx *gin.Context) {

	
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": ctx.Param("todoListID"),
	})
}

func patchTodoList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": ctx.Param("todoListID"),
	})
}

func showTodoList(ctx *gin.Context) {

	todoListID := ctx.Param("todoListID")

	ctx.JSON(http.StatusOK, gin.H{
		"id": todoListID,
	})
}

func todoitems (ctx *gin.Context){

	todoItems := ctx.Param("todoITEM")

	ctx.JSON(http.StatusOK, gin.H{
		"data" : todoItems,
	})
}

func addtodoitem(ctx *gin.Context){
	todoItems := ctx.Param("todoITEM")

	ctx.JSON(http.StatusOK,gin.H{
		"id": todoItems,
	})
}

func marktodoitem(ctx *gin.Context){

	todoItems := ctx.Param("todoITEM")

	ctx.JSON(http.StatusOK, gin.H{
		"id":todoItems,
	})
}
