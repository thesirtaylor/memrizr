package handler

import (
	"log"
	"net/http"

	"github.com/thesirtaylor/memrizr/model"
	"github.com/thesirtaylor/memrizr/utils"

	"github.com/gin-gonic/gin"
)

//handler calls service for getting user detaila
func (h *Handler) Me(c *gin.Context){
	user, exists := c.Get("user");

	if !exists {
		log.Printf("Unable to extract user from request context for unknown reason: %v\n", c)
		err := utils.NewInternal();
		c.JSON(err.Status(), gin.H{
			"error": err,
		});
		return;
	}

	uid := user.(*model.User).UID;
	u, err := h.UserService.Get(c, uid);

	if err != nil  {
		log.Printf("Unable to find user: %v\n%v", uid, err)
		e := utils.NewNotFound("user", uid.String());

		c.JSON(e.Status(), gin.H{
			"error": e,
		});
		return;
	}

	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}