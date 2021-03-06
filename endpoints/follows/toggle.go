package follows

import (
	"github.com/bredbrains/tthk-wish-list/modules"
	"net/http"
	"strconv"
	"time"

	"github.com/bredbrains/tthk-wish-list/database"
	"github.com/bredbrains/tthk-wish-list/models"
	"github.com/gin-gonic/gin"
)

func ToggleFollowing(c *gin.Context) {
	var follow models.Follow
	var userCalled models.User
	var err error
	var id int
	token := c.GetHeader("Token")
	err, userCalled = database.UserData(token)
	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"success": false, "error": "Invalid token"})
		return
	}
	database.GetFollowsFromUser(userCalled)
	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "User with this ID doesn't exist."})
		return
	}
	following, isSameUser := modules.CheckIsFollowed(userCalled, id)
	if isSameUser {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "You cannot start follow yourself."})
		return
	} else {
		follow = models.Follow{UserFrom: userCalled.ID, UserTo: id, CreationTime: time.Now().Format("2006-01-02 15:04:05")}
		if following {
			err = database.DeleteFollow(follow)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"success": true})
			return
		} else {
			err, follow = database.AddFollow(follow)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"success": true, "follow": follow})
			return
		}
	}
}
