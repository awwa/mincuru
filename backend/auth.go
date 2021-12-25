package main

import (
	"errors"
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
		SendCookie:  true,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*UserResp); ok {
				return jwt.MapClaims{
					"id":   v.Id,
					"role": v.Role,
				}
			}
			return jwt.MapClaims{}
		},
		// LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
		// 	c.JSON(http.StatusOK, gin.H{
		// 		"code":   http.StatusOK,
		// 		"token":  token,
		// 		"expire": expire.Format(time.RFC3339),
		// 		"user": gin.H{
		// 			"id":    1,
		// 			"name":  "Hoge taro",
		// 			"email": "hoge@example.com",
		// 			"role":  "user",
		// 		},
		// 	})
		// },
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &UserResp{
				Id:   uint(claims[identityKey].(float64)),
				Role: string(claims["role"].(string)),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals User
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			var users []User
			query := User{}
			query.Email = loginVals.Email
			// DBからuserレコード取得
			result := DB.Table("users").Where(&query).Find(&users)
			// 該当レコードなし。認証エラー
			if result.RowsAffected /*(userResp)*/ != 1 {
				return nil, errors.New("no user found")
			}
			// ハッシュ化したパスワードの比較
			if err := bcrypt.CompareHashAndPassword(
				([]byte)(users[0].Password),
				([]byte)(loginVals.Password),
			); err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			return &UserResp{
				Id:   users[0].Id,
				Role: users[0].Role,
			}, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			// IdentityHandlerがClaimから取り出した情報（＝JWTに格納されている情報）を受け取り
			// API実行の条件を満たしているか判定しtrue/falseで返す
			path := c.FullPath()
			method := c.Request.Method
			v, _ := data.(*UserResp)
			role := v.Role
			// 全ユーザー共通
			switch {
			case (path == "/users/refresh_token" && method == "GET"):
				return true
			case (path == "/users/logout" && method == "POST"):
				return true
			}
			// 管理者の場合、無条件で通過
			if role == "admin" {
				return true
			}
			// ユーザーの場合のアクセス制限
			if role == "user" {
				switch {
				case (path == "/users/me"):
					return true
				}
			}
			// ゲストの場合
			if role == "guest" {
				switch {
				case (path == "/users/me"):
					return true
				}
			}
			return false
			// 	if v, ok := data.(*User); ok && v.UserName == "admin" {
			// 		return true
			// 	}
			// 	return false
		},
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
