package NetFlixController

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "connection"
const dbName = "netlfix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	//client option

	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongodb connection success")
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is ready")
}
