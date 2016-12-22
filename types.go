package grf

import (
	"gopkg.in/gin-gonic/gin.v1"
)

type StatusCode int
type ResponseData interface {}
type Action string

type ListSupported interface {
	List(*gin.Context)
}

type RetrieveSupported interface {
	Retrieve(*gin.Context)
}