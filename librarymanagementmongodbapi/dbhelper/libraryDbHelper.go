package libraryDbHelper

import (
	"context"
	"fmt"
	"log"

	connectwithdb "github.com/rubansundararaj/librarymanagement/dbconnection"
	librarymodel "github.com/rubansundararaj/librarymanagement/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "simplified"
const colName = "testunique"

var collection *mongo.Collection

func init() {
	client := connectwithdb.InitializeAndReturnConnection()

	db := client.Database(dbName) //.Collection(colName)

	// Create a "users" collection with a JSON schema validator. The validator
	// will ensure that each document in the collection has "name" and "age"
	// fields.
	jsonSchema := bson.M{
		"bsonType": "object",
		"required": []string{"bookname", "bookauthor", "location", "uniqueid"},
		"properties": bson.M{
			"bookname": bson.M{
				"bsonType": "string",
			},
			"bookauthor": bson.M{
				"bsonType": "string",
			},
			"location": bson.M{
				"bsonType": "string",
			},
			"uniqueid": bson.M{
				"bsonType": "string",
			},
		},
	}
	validator := bson.M{
		"$jsonSchema": jsonSchema,
	}
	opts := options.CreateCollection().SetValidator(validator)

	names, err := db.ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		// Handle error
		panic(err)
	}
	for _, name := range names {
		fmt.Println(name + "  " + colName)
		if name == colName {
			collection = db.Collection(colName)
			fmt.Println("Collection exist mate so we are skipping it")
			return
		}
	}

	creationerror := db.CreateCollection(context.TODO(), colName, opts)
	if creationerror != nil {
		log.Fatal(creationerror)
	}
	collection = db.Collection(colName)

}

func InsertOneBookToDb(book librarymodel.LibraryModel) bool {
	fmt.Println("Going to insert to DB***********")
	inserted, err := collection.InsertOne(context.Background(), book)
	if err != nil {
		fmt.Println(string(err.Error()))
		return false
	}
	fmt.Println("Inserted 1 movie", inserted.InsertedID)
	return true
}

func GetOneBookRecordFromDb(id string) (bool, bson.M) {
	filter := bson.M{"uniqueid": id}
	cursor := collection.FindOne(context.Background(), filter)

	fmt.Println(id)
	var book bson.M
	err := cursor.Decode(&book)
	if err != nil {
		return false, book
	}
	return true, book

}

func UpdateOneBookRecordInDb(uniqueId string, book librarymodel.LibraryModel) bool {
	filter := bson.M{"uniqueid": uniqueId}
	update := bson.M{"$set": bson.M{
		"bookname":   book.BookName,
		"bookauthor": book.BookAuthor,
		"location":   book.Location,
		"uniqueid":   book.Uniqueid}}

	fmt.Println(filter)
	fmt.Println(update)

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func DeleteOneBookRecordFromBb(id string) bool {

	filter := bson.M{"uniqueid": id}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return false
	}
	return true
}

func GetAllBookRecordsFromDb() (bool, []primitive.M) {
	var dbdata []primitive.M

	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		fmt.Println(err)
		return false, dbdata
	}

	for cursor.Next(context.Background()) {
		var data bson.M
		err := cursor.Decode(&data)
		if err != nil {
			fmt.Println(err)
			return false, dbdata
		}
		dbdata = append(dbdata, data)
	}

	defer cursor.Close(context.Background())

	return true, dbdata
}
