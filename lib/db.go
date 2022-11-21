package db

import (
	_ "bufio"
	"context"
	_ "encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	_ "strconv"
	_ "strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type Trainer struct {
//     Name string
//     Age  int
//     City string
// }

type Ingredient struct {
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
}

type Recipe struct {
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
	Notes       string       `json:"notes"`
	Icon        string       `json:"icon"`
}

func ConnectDB() *mongo.Collection {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("MONGODB_URI"))
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	coll := client.Database("recipes").Collection("recipes")
	return coll
}

var coll *mongo.Collection = ConnectDB()

func GetRecipe(ctx *gin.Context) {
	fmt.Println("Getting single recipe")

	name := ctx.Param("name")
	fmt.Println(name)

	var result Recipe
	err := coll.FindOne(context.TODO(), bson.D{{"name", name}}).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}

	ctx.IndentedJSON(http.StatusOK, result)
}

func GetRecipes(ctx *gin.Context) {
	fmt.Println("Getting multiple recipes")

	var results []Recipe
	cur, err := coll.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem Recipe
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}
	cur.Close(context.TODO())
	ctx.IndentedJSON(http.StatusOK, results)
}

func AddRecipe(ctx *gin.Context) {
	var newRecipe Recipe

	if err := ctx.BindJSON(&newRecipe); err != nil {
		return
	}

	insertResult, err := coll.InsertOne(context.TODO(), newRecipe)
	if err != nil {
		log.Fatal(err)
	}

	ctx.IndentedJSON(http.StatusCreated, insertResult)
}

// fmt.Println("Simple Shell")
// fmt.Println("---------------------")

// fmt.Print("Recipe name: ")
// recipeName := readInput()

// var ingredients []Ingredient

// for {
//     fmt.Print("Add ingredient: ")
//     ingredientName := readInput()

//     if strings.Compare("exit", ingredientName) == 0 {
//         fmt.Println("Exiting...")
//         break
//     }

//     fmt.Print("Quantity: ")
//     ingredientQuantity, _ := strconv.ParseFloat(readInput(), 32)

//     fmt.Print("Unit: ")
//     ingredientUnit := readInput()

//     ingredient := Ingredient {
//         Name: ingredientName,
//         Quantity: ingredientQuantity,
//         Unit: ingredientUnit,
//     }

//     ingredients = append(ingredients, ingredient)

// }

// recipe := Recipe {
//     Name: recipeName,
//     Ingredients: ingredients,
//     Notes: "",
// }

// fmt.Println(recipe)

// // ash := Trainer{"Ash", 10, "Pallet Town"}
// insertResult, err := coll.InsertOne(context.TODO(), recipe)

// if err != nil {
// log.Fatal(err)
// }

// fmt.Println("Inserted a single document: ", insertResult.InsertedID)

// filter := bson.D{{"name", "Ash"}}

// update := bson.D{
//     {"$inc", bson.D{
//         {"age", 1},
//     }},
// }

// updateResult, err := coll.UpdateOne(context.TODO(), filter, update)
// if err != nil {
//     log.Fatal(err)
// }

// fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

// title := "Back to the Future"
// var result bson.M
// err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)
// if err == mongo.ErrNoDocuments {
// 	fmt.Printf("No document was found with the title %s\n", title)
// 	return
// }
// if err != nil {
// 	panic(err)
// }
// jsonData, err := json.MarshalIndent(result, "", "    ")
// if err != nil {
// 	panic(err)
// }
// fmt.Printf("%s\n", jsonData)
// }
