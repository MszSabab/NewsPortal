package newsportal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Init - initializes todo module
func Init(router *gin.Engine) {
	router.POST("portal", CreateNewsHandler)
	router.GET("portal", GetAllNewsHandler)
	router.GET("portal/:id", GetSingleNewsHandler)
	router.DELETE("portal/:id", DeleteNewsHandler)
	router.PUT("portal/:id", UpdateNewsHandler)

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
			Err:    "Error in data binding",
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

//GetSingleNewsResponse ----
type GetSingleNewsResponse struct {
	Newses NewsPortal `json:"newses"`
	Err    string     `json:"err"`
}

//GetSingleNewsHandler ----
func GetSingleNewsHandler(c *gin.Context) {
	portalID := c.Param("id")
	getSingleNews, err := GetSingleNewsCrud(portalID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, GetSingleNewsResponse{
			Newses: NewsPortal{},
			Err:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, GetSingleNewsResponse{
		Newses: getSingleNews,
		Err:    "",
	})
	return
}

//DeleteNewsResponse ----
type DeleteNewsResponse struct {
	Err string `json:"err"`
}

//DeleteNewsHandler ----
func DeleteNewsHandler(c *gin.Context) {
	portalID := c.Param("id")
	err := DeleteNewsCrud(portalID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, DeleteNewsResponse{

			Err: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, DeleteNewsResponse{

		Err: "",
	})
	return

}

//UpdateNewsRequest ----
type UpdateNewsRequest struct {
	Newses NewsPortal
}

//UpdateNewsResponse ----
type UpdateNewsResponse struct {
	Newses NewsPortal `json:"newses"`
	Err    string     `json:"err"`
}

//UpdateNewsHandler ----
func UpdateNewsHandler(c *gin.Context) {

	var req UpdateNewsRequest
	portalID := c.Param("id")
	if err := c.ShouldBindJSON(&req.Newses); err != nil {
		c.JSON(http.StatusInternalServerError, UpdateNewsResponse{
			Newses: NewsPortal{},
			Err:    "Error in data binding",
		})
		return
	}
	updateNews, err := UpdateNewsCrud(portalID, req.Newses)
	if err != nil {
		c.JSON(http.StatusInternalServerError, UpdateNewsResponse{
			Newses: NewsPortal{},
			Err:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, UpdateNewsResponse{
		Newses: updateNews,
		Err:    "",
	})
	return

}
