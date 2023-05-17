package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) task(c *gin.Context) {
	idTest := c.Query("idTest")
	idUser := c.Query("idUser")
	answer := c.Query("answer")
	corrAnswer := c.Query("corrAnswer")

	err := h.service.AddAnswer(idTest, idUser, answer, corrAnswer)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	idTask := c.Query("idTask")
	idVariant := c.Query("idVariant")

	task, err := h.service.Task(idUser, idTask, idVariant, idTest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}