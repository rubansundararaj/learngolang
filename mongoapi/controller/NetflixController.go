package NetFlixController

import (
	"context"
	"fmt"
	"log"
	"mongoapi/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func insertOneMovie(movie model.NetFlix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie", inserted.InsertedID)
}

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted one count : ", deleteCount)

}

func deleteAllMovie() int64 {
	deleteCollection, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of deleted movies : ", deleteCollection.DeletedCount)
	return deleteCollection.DeletedCount
}
