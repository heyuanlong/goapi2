package base

import (
	kinit "goapi2/initialize"

	"github.com/gin-gonic/gin"
)

type RouteWrapStruct struct {
	Method string
	Path   string
	F      func(*gin.Context)
}

func Wrap(Method string, Path string, f func(*gin.Context), types int) RouteWrapStruct {
	wp := RouteWrapStruct{
		Method: Method,
		Path:   Path,
	}

	wp.F = func(c *gin.Context) {
		kinit.LogWarn.Println("types:", types)
		f(c)
	}
	return wp
}
