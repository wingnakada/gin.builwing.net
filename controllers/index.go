package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	m "gin/models"
)

func GetIndex(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main Website",
			"info":"",
		})
}

func GetSignin(c *gin.Context) {
	c.HTML(http.StatusOK, "signin.html", "")
}
func PostSignin(c *gin.Context) {
	u := m.User{}
	email := c.PostForm("email")
	pass := c.PostForm("password")
	if u.IsLogin(email, pass) {
		u.SessionSet(c)
		c.Redirect(http.StatusMovedPermanently, "/users")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/signin")
	}
}
func GetSignout(c *gin.Context) {
	m.SessionClear(c)
	c.HTML(http.StatusOK, "signin.html", "")
}

