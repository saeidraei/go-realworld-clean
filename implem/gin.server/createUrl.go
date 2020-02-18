package server

import (
	"github.com/saeidraei/go-realworld-clean/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saeidraei/go-realworld-clean/implem/json.formatter"
)

type UrlReq struct {
	Url struct {
		Address *string `json:"address"`
	} `json:"url,required"`
}

func urlFromRequest(req *UrlReq) domain.Url {
	return domain.Url{
		Address: *req.Url.Address,
	}
}

func (rH RouterHandler) createUrl(c *gin.Context) {
	log := rH.log(rH.MethodAndPath(c))

	req := &UrlReq{}
	if err := c.BindJSON(req); err != nil {
		log(err)
		c.Errors = append(c.Errors, &gin.Error{Err:err,Type:gin.ErrorTypePublic})
		c.Status(http.StatusBadRequest)
		return
	}

	url, err := rH.ucHandler.UrlPost(urlFromRequest(req))
	if err != nil {
		log(err)
		c.Errors = append(c.Errors, &gin.Error{Err:err,Type:gin.ErrorTypePublic})
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"url": formatter.NewUrlFromDomain(*url)})
}
