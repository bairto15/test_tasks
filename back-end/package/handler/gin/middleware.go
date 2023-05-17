package gin

import (
	"net/http"
	"test_puzzle/package/service"

	"github.com/gin-gonic/gin"
)

func (h *Handler) middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		login := c.Request.Header["Authorization"]
		
		if len(login) == 0 || !service.RedisAuthUser[login[0]] {
			h.service.Logger.Warn("Не авторизированный пользователь")
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"errors": "Не авторизированный пользователь"})
			return
		}
	}
}
