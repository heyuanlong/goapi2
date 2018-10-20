package base

import (
	"github.com/gin-gonic/gin"
	klog "github.com/heyuanlong/go-utils/common/log"
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
		klog.Warn.Println("types:", types)
		f(c)
	}
	return wp
}
