package endpoints

import (
	"net/http"
  "strconv"

	"github.com/gin-gonic/gin"
  "github.com/goferpwlynie/goRestApi/users"
  "github.com/goferpwlynie/goRestApi/jsonTools"
)

func handleIdParam(ctx *gin.Context) int{
  Id, err := strconv.Atoi(ctx.Param("id"))

  if err != nil{
    ctx.JSON(http.StatusBadRequest, gin.H{
      "Error": "Wrong Id type",
    })
  }

  return Id
}

func GetUsersHandler(ctx *gin.Context){
    ctx.JSON(http.StatusOK, users.Users)
}

func GetUserHandler(ctx *gin.Context){
  id := handleIdParam(ctx)

  for i, v := range users.Users{
    if v.Id == id{
      ctx.JSON(http.StatusOK, users.Users[i])
      break
    }
  }
}

func PostUsersHandler(ctx *gin.Context) {
    var user users.User

    err := ctx.ShouldBindJSON(&user)
    if err != nil{
      ctx.JSON(http.StatusBadRequest,gin.H{
        "error": err.Error(),
      })
    }

   users.Users = append(users.Users, user) 
   jsontools.WriteJsonFile()

   ctx.JSON(http.StatusOK, gin.H{
     "message": user,
   })
  }

func DeleteUsersHandler(ctx *gin.Context) {
  id := handleIdParam(ctx) 

  for i, v := range users.Users{
    if v.Id == id{
      users.Users = append(users.Users[:i], users.Users[i+1:]...)
      break
    }
  }
    
  jsontools.WriteJsonFile()
  ctx.Status(http.StatusCreated)
}

func PutUserHandler(ctx *gin.Context){
  id := handleIdParam(ctx)

  var user users.User

  err := ctx.ShouldBindJSON(&user)

  if err != nil{
    ctx.JSON(http.StatusBadRequest,gin.H{
      "error": err.Error(),
    })
  }

  user.Id = id
  
  for i, v := range users.Users{
    if v.Id == id{
      users.Users[i] = user
    }
  }

  ctx.Status(http.StatusNoContent)
}

func PatchUserHandler(ctx *gin.Context){
  id := handleIdParam(ctx)
  
  var user users.User

  err := ctx.ShouldBindJSON(&user)
  
  if err != nil{
    ctx.JSON(http.StatusBadRequest,gin.H{
      "error": err.Error(),
    })
  }

  for i, v := range users.Users{
    if v.Id == id{
      if user.Name != ""{
        users.Users[i].Name = user.Name
      }
      if user.Surname != ""{
        users.Users[i].Surname = user.Surname
      }
      if user.BirthYear != 0{
        users.Users[i].BirthYear = user.BirthYear
      }
    }
  }
}
