package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var identityKey = "id"

func authMiddleware() (authMiddleware *jwt.GinJWTMiddleware) {
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "mincuru-api-server",
		Key:         []byte(os.Getenv("JWT_KEY")),
		Timeout:     time.Hour, // トークン有効期限
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		// SendCookie:  true,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			fmt.Println("*****1")
			fmt.Println(data)
			v, ok := data.(*UserResp)
			fmt.Println(v)
			fmt.Println(ok)
			if v, ok := data.(*UserResp); ok {
				// if v, ok := data.(*User); ok {
				fmt.Println(v)
				return jwt.MapClaims{
					"email": v.Email,
					"id":    v.ID,
				}
			}
			fmt.Println("*****2")
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			fmt.Println("######")
			claims := jwt.ExtractClaims(c)
			fmt.Println(claims)
			fmt.Println(claims["exp"])
			fmt.Println(claims["orig_iat"])
			fmt.Println(claims[identityKey])
			return &UserResp{
				Id:    Id{ID: (uint)(1) /*claims["id"].(float64)*/},
				Email: claims["email" /*identityKey*/].(string),
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
			//return &UserResp{Email: loginVals.Email}, nil
			fmt.Println("$$$$$$ auth success")
			fmt.Println(dbUsers[0].ID)
			fmt.Println(loginVals.Email)
			return &UserResp{Id: Id{ID: dbUsers[0].ID}, Email: loginVals.Email}, nil
		},
		// Authorizator: func(data interface{}, c *gin.Context) bool {
		// 	fmt.Println("&&&&&&")
		// 	// 	if v, ok := data.(*User); ok && v.UserName == "admin" {
		// 	return true
		// 	// 	}
		// 	// return false
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
