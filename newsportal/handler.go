package newsportal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Init - initializes todo module
func Init(router *gin.Engine) {
	router.POST("portal", CreateNewsHandler)
	router.GET("portal", GetAllNewsHandler)
	// router.GET("todo/:id", GetOneTodoHandler)
	// router.DELETE("todo/:id", DeleteOneTodoHandler)
	// router.PUT("todo/:id", UpdateTodoHandler)

}

//CreateNewsRequest ----
type CreateNewsRequest struct {
	Newses NewsPortal
}

//CreateNewsResponse ----
type CreateNewsResponse struct {
	Newses NewsPortal `json:"newses"`
	Err    string     `json:"err"`
}

//CreateNewsHandler ----
func CreateNewsHandler(c *gin.Context) {
	var req CreateNewsRequest

	if err := c.BindJSON(&req.Newses); err != nil {
		c.JSON(http.StatusInternalServerError, CreateNewsResponse{
			Newses: NewsPortal{},
			Err:    "",
		})
		return
	}

	createNews, err := CreateNewsCrud(req.Newses)

	if err != nil {
		c.JSON(http.StatusInternalServerError, CreateNewsResponse{
			Newses: NewsPortal{},
			Err:    "",
		})
		return
	}

	c.JSON(http.StatusOK, CreateNewsResponse{
		Newses: createNews,
		Err:    "",
	})
	return
}

//GetAllNewsResponse ----
type GetAllNewsResponse struct {
	Newses []NewsPortal `json:"newses"`
	Err    string       `json:"err"`
}

// GetAllNewsHandler ----
func GetAllNewsHandler(c *gin.Context) {

	getAllNews, err := GetAllNewsCrud()

	if err != nil {
		c.JSON(http.StatusInternalServerError, GetAllNewsResponse{
			Newses: []NewsPortal{},
			Err:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, GetAllNewsResponse{
		Newses: getAllNews,
		Err:    "",
	})
	return

}

//	portalID := c.Param("id")
