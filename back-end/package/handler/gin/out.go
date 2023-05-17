package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) out(c *gin.Context) {
	req := signin{}

	if err := c.BindJSON(&req); err != nil {
		h.service.Logger.Warn(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	err := h.service.Out(req.Login)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}