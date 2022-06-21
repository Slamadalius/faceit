package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/Slamadalius/faceit/internal/server"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := fmt.Sprintf("mongodb://%s:%s@mongoDB:27017/?maxPoolSize=20", os.Getenv("MONGO_USERNAME"), os.Getenv("MONGO_PASSWORD"))

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")

	router := mux.NewRouter()
	router.HandleFunc("/faceit", handler).Methods("GET")

	server := server.Server{}
	server.Start(router)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	server.Shutdown()
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
