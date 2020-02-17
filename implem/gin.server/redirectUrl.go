package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (rH RouterHandler) redirectUrl(c *gin.Context) {
	log := rH.log(rH.MethodAndPath(c))

	url, err := rH.ucHandler.UrlGet(c.Param("id"))
	if err != nil {
		log(err)
		c.Status(http.StatusUnprocessableEntity)
		return
	}
	c.Redirect(http.StatusMovedPermanently, url.Address)
	c.Abort()
}
