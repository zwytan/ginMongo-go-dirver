package middle

import (
	"fmt"
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/xyfll7/login/models"
)

// JwtAPI provides handlers for managing admin.
type JwtAPI struct {
	DB AdminDatabase
}

// AdminDatabase interface for encapsulating database access.
type AdminDatabase interface {
	FindAdmin(auth *models.Auth) (*models.Admin, error)
}

var identityKey = "id"
var roles = "roles"

// GinJwt xxx
func (j *JwtAPI) GinJwt() *jwt.GinJWTMiddleware {
	// the jwt middle
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte("secret keys"),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     identityKey,
		Authenticator:   j.authenticator,
		PayloadFunc:     j.payloadFunc,
		IdentityHandler: j.identityHandler,
		Authorizator:    j.authorizator,
		Unauthorized:    j.unauthorized,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return authMiddleware
}

func (j *JwtAPI) authenticator(c *gin.Context) (interface{}, error) {
	fmt.Println("111")
	var loginVals models.Auth
	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	adminID := loginVals.Adminname
	password := loginVals.Password

	admin, err := j.DB.FindAdmin(loginVals.New())
	if err != nil {
		return nil, jwt.ErrFailedAuthentication // 账号错误
	}
	if adminID == admin.Name || adminID == admin.Email &&
		password == admin.Password {
		return admin, nil
	}
	return nil, jwt.ErrFailedAuthentication // 账号错误
}
func (j *JwtAPI) payloadFunc(data interface{}) jwt.MapClaims {
	fmt.Println("222")
	if v, ok := data.(*models.Admin); ok {
		fmt.Println(v, ok)
		return jwt.MapClaims{
			identityKey: v.Name,
			roles:       v.Roles,
		}
	}
	return jwt.MapClaims{}
}
func (j *JwtAPI) identityHandler(c *gin.Context) interface{} {
	fmt.Println("333")
	claims := jwt.ExtractClaims(c)
	return &models.Admin{
		Roles: []string{claims[identityKey].(string)},
	}
}
func (j *JwtAPI) authorizator(data interface{}, c *gin.Context) bool {
	fmt.Println("444")
	if v, ok := data.(*models.Admin); ok && v.Roles[0] == "admin" {
		return true
	}
	return false
}
func (j *JwtAPI) unauthorized(c *gin.Context, code int, message string) {
	fmt.Println("555")
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

// NoRouteHandler 404 handler
func (j *JwtAPI) NoRouteHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	fmt.Printf("NoRoute claims: %#v\n", claims)
	c.JSON(404, gin.H{
		"code":    "PAGE_NOT_FOUND",
		"message": "Page not found",
	})
}
