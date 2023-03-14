package api

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"mio-api-interface/utils"
	"net/http"
)

func (Api) GetSentence(c *gin.Context) {
	n := rand.Intn(20)
	c.JSON(http.StatusOK, utils.ResponseOK(Sentences[n]))
}
