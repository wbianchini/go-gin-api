package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wbianchini/go-gin-api/pkg/common/models"
)

func (h handler) DeleteBook(c *gin.Context) {
	var book models.Book
	if result := h.DB.First(&book, c.Param("id")); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if result := h.DB.Delete(&book); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.Status(http.StatusOK)
}
