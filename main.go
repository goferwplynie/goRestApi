package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/goferpwlynie/goRestApi/auth"
	"github.com/goferpwlynie/goRestApi/db"
	"github.com/goferpwlynie/goRestApi/endpoints"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error reading .env file: %v\n", err)
	}
	connString := os.Getenv("DATABASE_URL")
	jwtKey := []byte(os.Getenv("JWT_SECRET"))

	db.ConnectToDB(connString)
	defer db.CloseDB()

	router := gin.New()

	secured := router.Group("/")
	secured.Use(auth.AuthMiddleware(jwtKey))

	router.GET("/users", endpoints.GetUsersHandler)
	router.GET("/users/:id", endpoints.GetUserHandler)

	router.POST("/login", endpoints.HandleLogin)
	secured.POST("/users", endpoints.PostUsersHandler)

	secured.PATCH("/users/:id", endpoints.PatchUserHandler)
	secured.DELETE("/users/:id", endpoints.DeleteUsersHandler)

	router.Run()

}
