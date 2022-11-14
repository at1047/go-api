package main

import (
    _"net/http"
    "github.com/gin-gonic/gin"
    _"errors"
    _"strconv"
	_"context"
    _"fmt"
    _"log"
    "go-recipes/lib"
	_"go.mongodb.org/mongo-driver/bson"
	_"go.mongodb.org/mongo-driver/mongo"
    "bufio"
    "os"
    "strings"
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



// var coll *mongo.Collection = db.ConnectDB()
func main() {
    // recipes := db.GetRecipes()
    // fmt.Println(recipes)
    // recipe := db.GetRecipe(coll)
    // fmt.Println(recipe)
    // recipes := db.GetRecipes(coll)
    // fmt.Println(recipes)
    router := gin.Default()
    router.GET("/recipes", db.GetRecipes)
    router.GET("/recipes/:name", db.GetRecipe)
    // fmt.Println(recipes)
    // router.GET("/api/recipes/{id}", GetTodo)
    router.POST("/recipes", db.AddRecipe)
    router.Run("localhost:9090")
}

// func getRecipes(context *gin.Context) {
//     recipes := db.GetRecipes(coll)
//     context.IndentedJSON(http.StatusOK, recipes)
// }


