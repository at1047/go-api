package main

import (
	"bufio"
	_ "context"
	_ "errors"
	_ "fmt"
	dbrecipes "go-api/pkg/dbrecipes"
	dbblog "go-api/pkg/dbblog"
	_ "log"
	_ "net/http"
	"os"
	_ "strconv"
	"strings"

	_ "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/mongo"
)

// type Todo struct {
//     ID          int     `json:"id"`
//     Item        string  `json:"item"`
//     Completed   bool    `json:"completed"`
// }
//
// var todos = []Todo {
//     {ID: 1, Item: "Clean Room", Completed: false},
//     {ID: 2, Item: "Read Book", Completed: false},
//     {ID: 3, Item: "Record Video", Completed: false},
// }
//

// func addTodo(context *gin.Context) {
//     var newTodo Todo
//
//     if err := context.BindJSON(&newTodo); err != nil {
//         return
//     }
//
//     todos = append(todos, newTodo)
//     context.IndentedJSON(http.StatusCreated, newTodo)
// }
//
// func getTodoById(id int) (*Todo, error) {
//     for i, t := range todos {
//         if t.ID == id {
//             return &todos[i], nil
//         }
//     }
//
//     return nil, errors.New("Todo not found")
// }
//
// func getTodo(context *gin.Context) {
//     id, err := strconv.Atoi(context.Param("id"))
//
//
//
//     todo, err := getTodoById(id)
//
//     if err != nil {
//         context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
//         return
//     }
//
//     context.IndentedJSON(http.StatusOK, todo)
// }

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())

	router.GET("/recipes", dbrecipes.GetRecipes)
	router.GET("/recipes/:name", dbrecipes.GetRecipe)
	router.POST("/recipes", dbrecipes.AddRecipe)

	router.GET("/blogs/:titleCode", dbblog.GetBlogContents)
	router.GET("/blogs", dbblog.GetAllBlogContents)
  router.GET("/blogtitles", dbblog.GetAllBlogTitles)
  router.GET("/blogtitles/:projectCode", dbblog.GetProjectBlogTitles)

  router.GET("/projects", dbblog.GetProjectTitles)
	router.POST("/blogs", dbblog.AddBlog)
	router.PUT("/blogs", dbblog.UpdateBlog)


	router.Run("localhost:9090")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
