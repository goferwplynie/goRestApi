package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIdk (c *gin.Context){
  c.JSON(http.StatusOK, gin.H{
    "message": "niga",
  })
}
