package grf

import "gopkg.in/gin-gonic/gin.v1"

type StatusCode int
type ResponseData interface {}

type CreateSupported interface {
	Create(c *gin.Context) (StatusCode, ResponseData)
}

type ListSupported interface {
	List(c *gin.Context) (StatusCode, ResponseData)
}

type RetrieveSupported interface {
	Retrieve(c *gin.Context) (StatusCode, ResponseData)
}

type QuerySet interface {
	All() *ResponseData
	One() *ResponseData
}

type Resource struct {
	Endpoint string
	QuerySet QuerySet
}

func AddResource(r *Resource, e *gin.Engine) *gin.RouterGroup {
	group := e.Group(r.Endpoint)
	return group
}