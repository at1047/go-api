package dbrecipes

import (
	_ "bufio"
	_ "context"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	_ "strconv"
	_ "strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv"
	//		"go.mongodb.org/mongo-driver/bson"
	//		"go.mongodb.org/mongo-driver/mongo"
	//		"go.mongodb.org/mongo-driver/mongo/options"
	//	  "go.mongodb.org/mongo-driver/bson/primitive"
)

// type Weather struct {
//     Id string `json:"id"`
//     Main  int `json:"main"`
//     Description string `json:"description"`
// 	Icon string `json:"icon"`
// }

// type Response struct {
//     Coord string `json:"coord"`
//     Weather  Weather
//     Base string
// 	Main string
// }

type Response struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

type WeatherShort struct {
	Time        string `json:"time"`
	Temp        string `json:"temp"`
	FeelsLike   string `json:"feels_like"`
	Main        string `json:"main"`
	Description string `json:"description"`
}

func GetWeather(ctx *gin.Context) {
	fmt.Println("Getting weather")

	uri := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=Atlanta,US&APPID=%s", os.Getenv("OWM_API_KEY"))
	resp, err := http.Get(uri)

	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	// fmt.Println(string(body))

	var result Response
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	// fmt.Println(PrettyPrint(result))
	// fmt.Println(res.Body)

	data := WeatherShort{
		Time:        time.Now().String(),
		Temp:        fmt.Sprintf("%.2f", result.Main.Temp-273.15),
		FeelsLike:   fmt.Sprintf("%.2f", result.Main.FeelsLike-273.15),
		Main:        result.Weather[0].Main,
		Description: result.Weather[0].Description,
	}

	ctx.IndentedJSON(http.StatusOK, data)
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

// func GetRecipes(ctx *gin.Context) {
// 	fmt.Println("Getting multiple recipes")

// 	var results []Recipe
// 	cur, err := coll.Find(context.TODO(), bson.D{{}})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for cur.Next(context.TODO()) {
// 		var elem Recipe
// 		err := cur.Decode(&elem)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		results = append(results, elem)
// 	}
// 	cur.Close(context.TODO())
// 	ctx.IndentedJSON(http.StatusOK, results)
// }

// func AddRecipe(ctx *gin.Context) {
// 	var newRecipe Recipe

// 	if err := ctx.BindJSON(&newRecipe); err != nil {
// 		return
// 	}

// 	insertResult, err := coll.InsertOne(context.TODO(), newRecipe)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	ctx.IndentedJSON(http.StatusCreated, insertResult)
// }

// func UpdateRecipe(ctx *gin.Context) {
// 	var newRecipe Recipe

// 	if err := ctx.BindJSON(&newRecipe); err != nil {
// 		return
// 	}

//  	id := newRecipe.ID
//   filter := bson.D{{"_id", id}}

//   result, err := coll.UpdateOne(context.TODO(), filter, bson.M{"$set": newRecipe})
//   if err != nil {
//     panic(err)
//   }

// 	ctx.IndentedJSON(http.StatusCreated, result)
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
