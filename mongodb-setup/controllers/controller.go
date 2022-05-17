package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"mongodb-setup/model"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString string = "YOUR_DATABASE_STRING"
const dbName string = "netflix"
const colName string = "watchlist"

// MOST IMPORTANT
var collection *mongo.Collection

// connecting with mongodb
func init() {
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		panic(err)
	}

	fmt.Println("MongoDB connection established")
	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready")
}

// MONGODB helpers - file

func InsertOneMovie(movie model.Netflix) interface{} {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted one movie in mongodb database with id:", inserted.InsertedID)
	// fmt.Printf("Type of Object Id is %T\n", inserted.InsertedID)
	return inserted.InsertedID
}

func UpdateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count: ", result.ModifiedCount)
}

func DeleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie got deleted with delete count:", deleteCount)
}

func DeleteAllMovies() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of movies deleted:", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

func GetAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())

	return movies
}

// Actual controllers - file

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all movies")
	w.Header().Set("Content-Type", "application/json")

	allMovies := GetAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Movie")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)

	movieId := InsertOneMovie(movie)
	valStr := fmt.Sprint(movieId)
	movie.ID, _ = primitive.ObjectIDFromHex(valStr)
	fmt.Println("Movie details is :", movie)
	json.NewEncoder(w).Encode(movie)

}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Mark Movie as Watched")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	UpdateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Single Movie")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	DeleteOneMovie(params["id"])
	json.NewEncoder(w).Encode("Movie " + params["id"] + " got deleted")
}

func DeleteAllTheMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete All Movie")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := DeleteAllMovies()
	json.NewEncoder(w).Encode("All movies deleted: " + string(count))
}
