package endpoints

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goferpwlynie/goRestApi/db"
	jsontools "github.com/goferpwlynie/goRestApi/jsonTools"
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

  rows, err := db.DB.Query(ctx, "select * from users where id=" + fmt.Sprint(id))
  if err != nil{
    panic(err)
  }
  defer rows.Close()
  for rows.Next(){
    var user users.User
		err := rows.Scan(&user.Id, &user.Name, &user.Surname, &user.BirthYear) 
    if err != nil{
      log.Fatalf("error reading rows: %v\n", err)
    }
  }
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

  query := "INSERT INTO users (name, surname, birth_year) VALUES('"+user.Name+"', '"+user.Surname+"', "+fmt.Sprint(user.BirthYear)+")"
  db.DB.Exec(ctx, query)

  ctx.Status(http.StatusCreated)
}

func DeleteUsersHandler(ctx *gin.Context) {
	id := handleIdParam(ctx)
  query := "DELETE FROM public.users WHERE id ="+ fmt.Sprint(id)+";"

  db.DB.Exec(ctx, query)


	ctx.Status(http.StatusCreated)
}

func PutUserHandler(ctx *gin.Context) {
	id := handleIdParam(ctx)

	var user users.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong data structure",
		})
		return
	}

	user.Id = id

	for i, v := range users.Users {
		if v.Id == id {
			users.Users[i] = user
		}
	}

	ctx.Status(http.StatusNoContent)
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

	for i, v := range users.Users {
		if v.Id == id {
			if user.Name != "" {
				users.Users[i].Name = user.Name
			}
			if user.Surname != "" {
				users.Users[i].Surname = user.Surname
			}
			if user.BirthYear != 0 {
				users.Users[i].BirthYear = user.BirthYear
			}
		}
	}
}
