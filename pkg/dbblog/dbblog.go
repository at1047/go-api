package dbblog

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
  _ "go.mongodb.org/mongo-driver/bson/primitive"
)

// type Trainer struct {
//     Name string
//     Age  int
//     City string
// }

// type Ingredient struct {
// 	Name     string  `json:"name"`
// 	Quantity float64 `json:"quantity"`
// 	Unit     string  `json:"unit"`
// }


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
	coll := client.Database("recipes").Collection("blog")
	return coll
}

var coll *mongo.Collection = ConnectDB()

type Blog struct {
  //ID    primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Title       string  `json:"title"`
	TitleCode       string  `json:"titleCode"`
	Project       string  `json:"project"`
	ProjectCode       string  `json:"projectCode"`
	Content     string  `json:"content"`
	Date        string  `json:"date"`
}
func GetBlogContents(ctx *gin.Context) {
	fmt.Println("Getting single blog")

	titleCode := ctx.Param("titleCode")
	fmt.Println(titleCode)

	var result Blog
	err := coll.FindOne(context.TODO(), bson.D{{"titleCode", titleCode}}).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}

	ctx.IndentedJSON(http.StatusOK, result)
}

func GetAllBlogContents(ctx *gin.Context) {
	fmt.Println("Getting all blogs")

  var results []Blog
	cur, err := coll.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem Blog
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}
	cur.Close(context.TODO())
	ctx.IndentedJSON(http.StatusOK, results)
}


type ProjectTitle struct {
	Project       string  `json:"project"`
	ProjectCode       string  `json:"projectCode"`
}
func GetProjectTitles(ctx *gin.Context) {
	fmt.Println("Getting multiple blogs")
	var results []ProjectTitle
  groupStage := bson.D{
    {"$group", bson.D{
        {"_id", "$projectCode"},
        {"project", bson.D{{"$max", "$project"}}},
        {"projectCode", bson.D{{"$max", "$projectCode"}}},
    }}}

  sortStage := bson.D{{"$sort", bson.D{{"project", 1}}}}

	cur, err := coll.Aggregate(context.TODO(), mongo.Pipeline{groupStage, sortStage})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem ProjectTitle
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}
	cur.Close(context.TODO())
	ctx.IndentedJSON(http.StatusOK, results)
}

type BlogTitle struct {
	Title       string  `json:"title"`
	TitleCode       string  `json:"titleCode"`
}
func GetProjectBlogTitles(ctx *gin.Context) {
	fmt.Println("Getting multiple blogs, filter by project")
	var results []BlogTitle

 	projectCode := ctx.Param("projectCode")
  matchStage := bson.D{{"$match", bson.D{{"projectCode", projectCode}}}}
  sortStage := bson.D{{"$sort", bson.D{{"date", -1}}}}

  groupStage := bson.D{
    {"$group", bson.D{
        {"_id", "$titleCode"},
        {"title", bson.D{{"$max", "$title"}}},
        {"titleCode", bson.D{{"$max", "$titleCode"}}},
    }}}

	cur, err := coll.Aggregate(context.TODO(), mongo.Pipeline{matchStage, groupStage, sortStage})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem BlogTitle
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}
	cur.Close(context.TODO())
	ctx.IndentedJSON(http.StatusOK, results)
}

func GetAllBlogTitles(ctx *gin.Context) {
	fmt.Println("Getting all blog titles")
	var results []BlogTitle

  groupStage := bson.D{
    {"$group", bson.D{
        {"_id", "$titleCode"},
        {"title", bson.D{{"$max", "$title"}}},
        {"titleCode", bson.D{{"$max", "$titleCode"}}},
    }}}

  sortStage := bson.D{{"$sort", bson.D{{"title", 1}}}}

	cur, err := coll.Aggregate(context.TODO(), mongo.Pipeline{groupStage, sortStage})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem BlogTitle
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}
	cur.Close(context.TODO())
	ctx.IndentedJSON(http.StatusOK, results)
}


// func GetProjectBlogTitles(ctx *gin.Context) {
// 	fmt.Println("Getting single project blog titles")
//  	project := ctx.Param("project")
//   filter := bson.D{{"project", project}}
//
// 	results, err := coll.Distinct(context.TODO(), "title", filter)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	ctx.IndentedJSON(http.StatusOK, results)
// }



// func AddRecipe(ctx *gin.Context) {
// 	var newRecipe Recipe
//
// 	if err := ctx.BindJSON(&newRecipe); err != nil {
// 		return
// 	}
//
// 	insertResult, err := coll.InsertOne(context.TODO(), newRecipe)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	ctx.IndentedJSON(http.StatusCreated, insertResult)
// }

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
