package auth

import (
	"github.com/Unknwon/macaron"
	"github.com/macaron-contrib/session"
	"log"
	"net/http"
)

// User 用户结构
type User struct {
	Id       int
	Name     string
	Password string
}

// Authenticate 验证用户
func Authenticate(user User, s session.Store) (int, string) {
	log.Println("authenticate start")
	authUser := s.Get("auth_user")
	if authUser != nil {
		s.Delete("auth_user")
	}
	if user.Name != "jacob" || user.Password != "abc123" {
		log.Println("authenticate error")
		return http.StatusUnauthorized, "Not Authorized"
	}
	s.Set("auth_user", user)

	log.Println("authenticate end")

	return http.StatusOK, "Authorized"
}

func Authorize() macaron.Handler {
	return func(ctx *macaron.Context, s session.Store) {
		log.Println("authorize start")
		user := s.Get("auth_user")
		if user == nil {
			log.Println("authorize error")
			ctx.HTML(http.StatusUnauthorized, "403")
			return
		}
		log.Println("authorize end")
		ctx.Next()
	}
}
