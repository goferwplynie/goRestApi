package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type User struct{
  Id int `json:"id"`
  Name string `json:"name"`
  Surname string `json:"surname"`
  BirthYear int `json:"birthYear"`
}

var Users []User

func LoadFromJson(){
  file, err := os.Open("users.json")

  if err != nil{
    fmt.Println("error opening users.json: ", err)
    return
  }

  defer file.Close()

  byteValue, err := io.ReadAll(file)

  if err != nil{
    fmt.Println("error reading users.json: ", err)
    return
  }

  err = json.Unmarshal(byteValue, &Users)

  if err != nil{
    fmt.Println("error unmarshaling json: ", err)
    return
  }
}

func writeJsonFile(){
  
}

func main(){
  LoadFromJson()

  router := gin.New()

  router.GET("/users", func (ctx *gin.Context){
    ctx.JSON(http.StatusOK, Users)
  })

  router.POST("/users", func(ctx *gin.Context) {
    var user User

    err := ctx.ShouldBindJSON(&user)
    if err != nil{
      ctx.JSON(http.StatusBadRequest,gin.H{
        "error": err.Error(),
      })
    }

   Users = append(Users, user) 

   ctx.JSON(http.StatusOK, gin.H{
     "message": user,
   })
  })

  router.Run()
}
