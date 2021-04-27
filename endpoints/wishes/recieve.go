package wishes

import (
	"net/http"

	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/models"
	"github.com/gin-gonic/gin"
)

func Receive(c *gin.Context) {
	var wishes []models.Wish
	var wish models.Wish
	var err error
	c.BindJSON(&wish)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
	}
	err, wishes = database.GetWishes(wish)
	message := gin.H{"wishes": wishes}
	c.JSON(http.StatusOK, message)
}