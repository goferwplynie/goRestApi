package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/goferpwlynie/goRestApi/db"
	"github.com/goferpwlynie/goRestApi/endpoints"
	jsontools "github.com/goferpwlynie/goRestApi/jsonTools"
	"github.com/joho/godotenv"
)

func main() {
	jsontools.LoadFromJson()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error reading .env file: %v\n", err)
	}
	connString := os.Getenv("DATABASE_URL")

	db.ConnectToDB(connString)
	defer db.CloseDB()

	router := gin.New()

	router.GET("/users", endpoints.GetUsersHandler)
	router.GET("/users/:id", endpoints.GetUserHandler)

	router.POST("/users", endpoints.PostUsersHandler)
	router.PUT("/users/:id", endpoints.PutUserHandler)
	router.PATCH("/users/:id", endpoints.PatchUserHandler)
	router.DELETE("/users/:id", endpoints.DeleteUsersHandler)

	router.Run()

}
