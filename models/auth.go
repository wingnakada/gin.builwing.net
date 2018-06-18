package models

import (
	"github.com/gin-contrib/sessions"
	"github.com/casbin/casbin"
	"github.com/gin-contrib/authz"
	"github.com/gin-gonic/gin"
)

func Authlizer(url string) gin.HandlerFunc {
	return func(c *gin.Context) {
		e := casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")
		authz.NewAuthorizer(e)
		session := sessions.Default(c)
		role := session.Get("role")
		if role != nil{
			sub := role
			obj := url + "/*"
			act := "*"
			if e.Enforce(sub,obj,act) == true {
				c.Next()
			} else {
				c.Redirect(301,"/signin")
			}
		} else {
			c.Redirect(301,"/signin")
		}
	}
}
//
///* Loginがパスするかどうかの判断 */
//func (u *User) IsLogin(email,pass string) bool {
//	//Email自体がなければ
//	if db.Where("email=?",email).First(u).RecordNotFound() {
//		return false
//	}else{
//		//passwordが一致すれば
//		if comparePass(u.Password,pass) {
//			return true
//		}else{
//			return false
//		}
//	}
//}
