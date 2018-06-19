package models

import (
	"github.com/gin-contrib/sessions"
	"github.com/casbin/casbin"
	"github.com/gin-contrib/authz"
	"github.com/gin-gonic/gin"
)
/*
 Casbinによる認証
 */
func Authlizer(url string) gin.HandlerFunc {
	return func(c *gin.Context) {
		e := casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")
		//roles := e.GetModel()
		//fmt.Print(roles)
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
