package grf

import (
	"net/http"
	"gopkg.in/gin-gonic/gin.v1"
)

type Resource interface {
	GetResourceOptions() *ResourceOptions
}

type ResourceOptions struct {
	Endpoint string
	QuerySet QuerySet
	SupportedActions []Action
}

func (r *ResourceOptions) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"results": r.QuerySet.All()})
}

func (r *ResourceOptions) Retrieve(c *gin.Context) {
	c.JSON(http.StatusOK, r.QuerySet.One(c.Param("id")))
}

func AddResource(r Resource, e *gin.Engine) *gin.RouterGroup {
	group := e.Group(r.GetResourceOptions().Endpoint)

	for _, action := range r.GetResourceOptions().SupportedActions {
		switch action {
		case LIST:
			group.GET("", r.(ListSupported).List)
			break
		case RETRIEVE:
			group.GET("/:id", r.(RetrieveSupported).Retrieve)
			break
		}
	}

	return group
}