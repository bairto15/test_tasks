package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type signin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *Handler) auth(c *gin.Context) {
	req := signin{}

	if err := c.BindJSON(&req); err != nil {
		h.service.Logger.Warn(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	idUser, err := h.service.Auth(req.Login, req.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error() })
		return
	}

	if idUser == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": "доступ запрещен" })
		return
	}

	variants, err := h.service.GetVariants()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	res := map[string]interface{}{
		"variants": variants,
		"idUser": idUser,
	}

	c.JSON(http.StatusOK, res)
}