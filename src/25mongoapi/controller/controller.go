package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kinjalrk2k/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://admin:u5B5skDe1L7P0Seq@cluster0.4du6u.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const collectionName = "watchlist"

// important!
var collection *mongo.Collection

// connect with MongoDB

// specialized method - runs on first time the application starts
func init() {
	// client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to MongoDB!")

	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection Reference is ready!")
}

/*
	MongoDB helpers
*/

// insert a record
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted with ID", inserted.InsertedID)
}

// update a record
func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id} // read more about bson.M vs bson.D
	update := bson.M{"$set": bson.M{"watched": true}}

	updated, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count", updated.ModifiedCount)
}

// delete one delete
func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	deleted, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted count", deleted.DeletedCount)
}

// delete all records
func deleteAllMovies() int64 {
	deleted, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted count", deleted.DeletedCount)
	return deleted.DeletedCount
}

// get all movies
func getAllMovies() []primitive.M {
	curr, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer curr.Close(context.Background())

	var movies []primitive.M
	for curr.Next(context.Background()) {
		var movie bson.M
		err := curr.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	return movies
}

/*
	Controllers
*/

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)

	json.NewEncoder(w).Encode(movie)
}

func MarkWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])

	json.NewEncoder(w).Encode(`{"msg": "Marked as watched"}`)
}

func DeleteOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])

	json.NewEncoder(w).Encode(`{"msg": "Deleted successfully"}`)
}

func DeleteAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	deleteAllMovies()

	json.NewEncoder(w).Encode(`{"msg": "Deleted all successfully"}`)
}
