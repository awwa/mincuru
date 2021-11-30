package main

import (
	"errors"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var identityKey = "email"

func authMiddleware() (authMiddleware *jwt.GinJWTMiddleware) {
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "mincuru-api-server",
		Key:         []byte(os.Getenv("JWT_KEY")),
		Timeout:     time.Hour, // トークン有効期限
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		SendCookie:  true,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*UserResp); ok {
				return jwt.MapClaims{
					"email": v.Email,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &UserResp{
				Email: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals User
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			query := User{}
			query.Email = loginVals.Email
			// DBからuserレコード取得
			var dbUsers []User
			if err := DB.Table("users").Where(&query).Find(&dbUsers).Error; err != nil {
				return nil, err
			}
			// 該当レコードなし。認証エラー
			if len(dbUsers) != 1 {
				return nil, errors.New("no user found")
			}
			// ハッシュ化したパスワードの比較
			if err := bcrypt.CompareHashAndPassword(
				([]byte)(dbUsers[0].Password),
				([]byte)(loginVals.Password),
			); err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			return &UserResp{Email: loginVals.Email}, nil
		},
		// Authorizator: func(data interface{}, c *gin.Context) bool {
		// 	if v, ok := data.(*User); ok && v.UserName == "admin" {
		// 		return true
		// 	}
		// 	return false
		// },
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.IndentedJSON(code, &ErrorResp{Message: message})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		panic(err)
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	if err := authMiddleware.MiddlewareInit(); err != nil {
		panic(err)
	}
	return
}
