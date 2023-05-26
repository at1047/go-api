package db

import (
	"context"
	_"encoding/json"
    "strconv"
	"fmt"
    "bufio"
    "strings"
	"log"
	"os"
	"github.com/joho/godotenv"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


// type Trainer struct {
//     Name string
//     Age  int
//     City string
// }

type Ingredient struct {
    Name string
    Quantity float64
    Unit string
}

type Recipe struct {
    Name string
    Ingredients []Ingredient
    Notes string
}


func main() {
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
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database("recipes").Collection("recipes")
    reader := bufio.NewReader(os.Stdin)

    readInput := func() string {
        text, _ := reader.ReadString('\n')
        text = strings.Replace(text, "\n", "", -1)
        return text
    }

    fmt.Println("Simple Shell")
    fmt.Println("---------------------")

    fmt.Print("Recipe name: ")
    recipeName := readInput()

    var ingredients []Ingredient

    for {
        fmt.Print("Add ingredient: ")
        ingredientName := readInput()

        if strings.Compare("exit", ingredientName) == 0 {
            fmt.Println("Exiting...")
            break
        }

        fmt.Print("Quantity: ")
        ingredientQuantity, _ := strconv.ParseFloat(readInput(), 32)

        fmt.Print("Unit: ")
        ingredientUnit := readInput()

        ingredient := Ingredient {
            Name: ingredientName,
            Quantity: ingredientQuantity,
            Unit: ingredientUnit,
        }

        ingredients = append(ingredients, ingredient)

    }

    recipe := Recipe {
        Name: recipeName,
        Ingredients: ingredients,
        Notes: "",
    }

    fmt.Println(recipe)

    // ash := Trainer{"Ash", 10, "Pallet Town"}
    insertResult, err := coll.InsertOne(context.TODO(), recipe)

    if err != nil {
    log.Fatal(err)
    }

    fmt.Println("Inserted a single document: ", insertResult.InsertedID)

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
}
