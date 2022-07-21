package middleware

import "github.com/kataras/iris/v12/middleware/basicauth"

var BasicAuth = basicauth.New(basicauth.Config{
	Users: map[string]string{
		"admin": "password",
	},
})
