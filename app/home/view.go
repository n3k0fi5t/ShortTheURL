package home

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func HomeRegister(rg *gin.RouterGroup) {
	rg.GET("/", homePageHandler)
	rg.POST("/", homePageHandler) // fix redirect issue
}

func homePageHandler(c *gin.Context) {
	var resp = make(gin.H)
	session := sessions.Default(c)

	if surl := session.Get("shortenURL"); surl != nil {
		resp["surl"] = surl.(string)
	}

	c.HTML(http.StatusOK, "main.html", resp)
}
