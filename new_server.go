package main

import (
	// "context"
	"fmt"
	"log"
	// "time"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	"net/http"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Got a request")
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
	_, err := fmt.Fprintf(w, "nghia's website")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	http.HandleFunc("/", homePageHandler)

	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
	/*
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://anh:WifeLien110493@cluster0.yzacf.mongodb.net/my_new_db?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Mongodb")

	fmt.Println("Databases:")
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	defer client.Disconnect(ctx)
	*/
}
