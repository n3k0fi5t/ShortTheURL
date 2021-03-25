package shorter

import (
	"net/http"
	"net/url"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Resource struct {
	service Service
}

func ShorterRegister(rg *gin.RouterGroup, service Service) {
	resource := Resource{service}

	rg.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "http://www.google.com/")
	})

	rg.POST("/shorten", resource.encode)
	rg.GET("/s/:code", resource.decode)
}

func (r Resource) decode(c *gin.Context) {
	code := c.Param("code")

	uri, err := r.service.Proxy(code)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.Redirect(http.StatusTemporaryRedirect, uri)
}

func (r Resource) encode(c *gin.Context) {
	var input struct {
		URL string `form:"shortenBox" binding:"required"`
	}

	if err := c.ShouldBind(&input); err != nil {
		c.String(http.StatusBadRequest, "Invalid request")
		return
	}

	if _, err := url.Parse(input.URL); err != nil {
		c.String(http.StatusBadRequest, "Invalid URL")
		return
	}

	path, err := r.service.Shorten(input.URL)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	u := url.URL{
		Scheme: "http",
		Host:   c.Request.Host,
		Path:   "s/" + path,
	}

	// update session
	session := sessions.Default(c)
	session.Set("shortenURL", u.String())
	session.Save()

	c.Redirect(http.StatusTemporaryRedirect, "/main")
}
