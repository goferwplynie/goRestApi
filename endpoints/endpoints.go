package endpoints

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goferpwlynie/goRestApi/auth"
	"github.com/goferpwlynie/goRestApi/db"
	requestmodels "github.com/goferpwlynie/goRestApi/requestModels"
	"github.com/goferpwlynie/goRestApi/users"
	"github.com/joho/godotenv"
)

func handleIdParam(ctx *gin.Context) int {
	Id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Wrong Id type",
		})
		return 0
	}

	return Id
}

func HandleLogin(ctx *gin.Context) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error reading .env file: %v\n", err)
	}
	login := os.Getenv("ADMIN_LOGIN")
	password := os.Getenv("ADMIN_PASSWORD")
	jwtKey := []byte(os.Getenv("JWT_SECRET"))

	var credentials requestmodels.LoginRequest

	err = ctx.ShouldBindJSON(&credentials)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
		return
	}

	if credentials.Name == login && credentials.Password == password {
		jwtToken, _ := auth.GenerateToken(login, jwtKey)
		ctx.JSON(http.StatusOK, gin.H{
			"token": jwtToken,
		})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": "wrong credentials",
	})
}

func GetUsersHandler(ctx *gin.Context) {
	rows, err := db.DB.Query(ctx, "select * from users")

	if err != nil {
		log.Fatalf("error querying db: %v\n", err)
	}

	var usersSlice []users.User

	defer rows.Close()
	for rows.Next() {
		var user users.User
		err := rows.Scan(&user.Id, &user.Name, &user.Surname, &user.BirthYear)
		if err != nil {
			log.Fatalf("error reading rows: %v\n", err)
		}
		usersSlice = append(usersSlice, user)
	}

	ctx.JSON(http.StatusOK, usersSlice)
}

func GetUserHandler(ctx *gin.Context) {
	id := handleIdParam(ctx)

	query := "select * from users where id=$1"

	row := db.DB.QueryRow(ctx, query, id)

	var user users.User
	err := row.Scan(&user.Id, &user.Name, &user.Surname, &user.BirthYear)

	if err != nil {
		log.Fatalf("error reading rows: %v\n", err)
	}

	ctx.JSON(http.StatusOK, user)
}

func PostUsersHandler(ctx *gin.Context) {
	var user users.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong data structure",
		})
		return
	}

	query := "INSERT INTO users (name, surname, birth_year) VALUES($1,$2,$3)"
	db.DB.Exec(ctx, query, user.Name, user.Surname, user.BirthYear)

	ctx.Status(http.StatusCreated)
}

func DeleteUsersHandler(ctx *gin.Context) {
	id := handleIdParam(ctx)
	query := "DELETE FROM public.users WHERE id = $1"

	db.DB.Exec(ctx, query, id)

	ctx.Status(http.StatusCreated)
}

func PatchUserHandler(ctx *gin.Context) {
	id := handleIdParam(ctx)

	var user requestmodels.PatchRequest

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong data structure",
		})
		return
	}

	if user.Name != nil {
		db.DB.Exec(ctx, "UPDATE users SET name = $1 WHERE id = $2", *user.Name, id)
	}
	if user.Surname != nil {
		db.DB.Exec(ctx, "UPDATE users SET surname = $1 WHERE id = $2", *user.Surname, id)
	}
	if user.BirthYear != nil {
		db.DB.Exec(ctx, "UPDATE users SET birth_year = $1 WHERE id = $2", *user.BirthYear, id)
	}
}
