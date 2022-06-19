package bases

import (
	"github.com/gin-gonic/gin"
)

var (
	BasePath   string
	Port       string
	ConfigPath string
	Router     *gin.Engine
)
