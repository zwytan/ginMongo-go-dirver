package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/xyfll7/login/database"
	"github.com/xyfll7/login/middle"
	"github.com/xyfll7/login/routers/api"
)

// InitGin Creates a router
func InitGin(db *database.MgClient) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middle.Cors())
	adminHandler := api.AdminAPI{DB: db} // db 实现了 接口DB 所以可以赋值
	jwtAPI := middle.JwtAPI{DB: db}
	authMiddleware := jwtAPI.GinJwt()
	r.NoRoute(authMiddleware.MiddlewareFunc(), jwtAPI.NoRouteHandler)

	sv := r.Group("/api/v1")
	auth := sv.Group("/auth")
	{
		auth.POST("/login", authMiddleware.LoginHandler)
		auth.POST("/regist", adminHandler.InsertAdmin)
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	}
	// sv.Use(authMiddleware.MiddlewareFunc())
	{
		sv.GET("/test", api.Test)
	}
	return r
}
