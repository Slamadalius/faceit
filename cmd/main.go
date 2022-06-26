package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/Slamadalius/faceit/internal/repository"
	"github.com/Slamadalius/faceit/internal/server"
	userHttpHandler "github.com/Slamadalius/faceit/internal/user/handler/http"
	"github.com/Slamadalius/faceit/internal/user/listener"
	userService "github.com/Slamadalius/faceit/internal/user/service"
	"github.com/Slamadalius/faceit/pkg/mongoDB"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const workersCount = 5

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	mongoUrl := fmt.Sprintf("mongodb://%s:%s@mongoDB:27017/?maxPoolSize=20", os.Getenv("MONGO_USERNAME"), os.Getenv("MONGO_PASSWORD"))
	mongoClient, err := mongoDB.New(mongoUrl)
	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUserRepository(mongoClient)
	userService := userService.NewUserService(userRepository)

	router := mux.NewRouter()

	userHttpHandler.NewUserHandler(router, userService)

	listener := listener.NewListener(workersCount)
	listener.Start()

	server := server.Server{}
	server.Start(router)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	server.Shutdown()
	listener.Stop()
	listener.Wait()
}
