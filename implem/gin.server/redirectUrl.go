package server

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (rH RouterHandler) redirectUrl(c *gin.Context) {
	log := rH.log(rH.MethodAndPath(c))

	url, err := rH.ucHandler.UrlGet(c.Param("id"))
	if err != nil {
		log(err)
		c.Errors = append(c.Errors, &gin.Error{Err:errors.New("url not found"),Type:gin.ErrorTypePrivate})
		c.Status(http.StatusNotFound)
		return
	}
	c.Redirect(http.StatusMovedPermanently, url.Address)
	c.Abort()
}
