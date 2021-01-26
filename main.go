package main

import (
	"github.com/MszSabab/NewsPortal/newsportal"
	"github.com/gin-gonic/gin"
)

//main function
func main() {
	router := gin.New()

	newsportal.Init(router)
	router.Run(":5001")

}
