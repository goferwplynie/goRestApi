package endpoints

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goferpwlynie/goRestApi/db"
	"github.com/goferpwlynie/goRestApi/users"
)

func handleIdParam(ctx *gin.Context) int {
	Id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Wrong Id type",
		})
	}

	return Id
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

	row := db.DB.QueryRow(ctx, "select * from users where id="+fmt.Sprint(id))

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

	query := "INSERT INTO users (name, surname, birth_year) VALUES('" + user.Name + "', '" + user.Surname + "', " + fmt.Sprint(user.BirthYear) + ")"
	db.DB.Exec(ctx, query)

	ctx.Status(http.StatusCreated)
}

func DeleteUsersHandler(ctx *gin.Context) {
	id := handleIdParam(ctx)
	query := "DELETE FROM public.users WHERE id =" + fmt.Sprint(id) + ";"

	db.DB.Exec(ctx, query)

	ctx.Status(http.StatusCreated)
}

func PatchUserHandler(ctx *gin.Context) {
	id := handleIdParam(ctx)

	var user users.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong data structure",
		})
		return
	}

	log.Println(user)

	if user.Name != "" {
		db.DB.Exec(ctx, "UPDATE users SET name = "+user.Name+" WHERE id = "+fmt.Sprint(id)+";")
	}
	if user.Surname != "" {
		db.DB.Exec(ctx, "UPDATE users SET surname = "+user.Surname+" WHERE id = "+fmt.Sprint(id)+";")
	}
	if user.BirthYear != 0 {
		db.DB.Exec(ctx, "UPDATE users SET birth_year = "+fmt.Sprint(user.BirthYear)+" WHERE id = "+fmt.Sprint(id)+";")
	}
}
