package main

import (
	"github.com/gin-gonic/gin"
  "github.com/goferpwlynie/goRestApi/endpoints"
  "github.com/goferpwlynie/goRestApi/jsonTools"
)

func main(){
  jsontools.LoadFromJson()

  router := gin.New()

  router.GET("/users", endpoints.GetUsersHandler)
  router.GET("/users/:id", endpoints.GetUserHandler)

  router.POST("/users", endpoints.PostUsersHandler)
  router.PUT("/users/:id", endpoints.PutUserHandler)
  router.DELETE("/users/:id", endpoints.DeleteUsersHandler)

  router.Run()

}
