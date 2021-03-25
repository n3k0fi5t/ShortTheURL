package route

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/n3k0fi5t/ShortTheURL/app/home"
	"github.com/n3k0fi5t/ShortTheURL/app/shorter"
	"github.com/n3k0fi5t/ShortTheURL/common/cache"
	"github.com/n3k0fi5t/ShortTheURL/common/config"
	"github.com/n3k0fi5t/ShortTheURL/common/dbx"
)

type Route struct {
	route *gin.Engine
	addr  string
}

func BuildRouter(cfg *config.Config, db *dbx.DB, rdb *cache.RDB) *Route {
	address := fmt.Sprintf(":%s", cfg.Server.Port)
	r := Route{
		addr:  address,
		route: gin.Default(),
	}

	// load template files
	r.route.LoadHTMLGlob("templates/*")

	// session middleware
	store := cookie.NewStore([]byte("secret"))
	r.route.Use(sessions.Sessions("ms", store))

	// home page
	home.HomeRegister(r.route.Group("/main"))

	// shorten service
	service := shorter.NewService(int(cfg.UrlLen), db, rdb)
	shorter.ShorterRegister(r.route.Group("/"), &service)

	return &r
}

func (r *Route) Run() {
	r.route.Run(r.addr)
}
