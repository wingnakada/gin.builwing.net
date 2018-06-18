package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	m "gin/models"
	"strconv"
	//"github.com/gin-contrib/sessions"
	//"github.com/gin-contrib/sessions/cookie"
)

var user m.User

func GetUsers(c *gin.Context){
	users := user.AllUsers()
	session := user.SessionGet(c)
	data := map[string]interface{}{"users":users,"session":session}
	c.HTML(http.StatusOK, "user.html", data)
}
func GetUsersTrashed(c *gin.Context){
	users := user.Trashed()
	session := user.SessionGet(c)
	data := map[string]interface{}{"users":users,"session":session}
	c.HTML(http.StatusOK, "user.html", data)
}
func UserSearch(c *gin.Context) {
	name,_ := c.GetPostForm("name")
	users := m.SearchUsers(name)
	c.HTML(http.StatusOK, "user.html", users)
}
func GetUserCreate(c *gin.Context) {
	c.HTML(http.StatusOK, "user-create.html", "")
}
func PostUserCreate(c *gin.Context) {
	user.Name = c.PostForm("name")
	user.Yomi = c.PostForm("yomi")
	user.Email = c.PostForm("email")
	user.Role = c.PostForm("role")
	pass := c.PostForm("password")
	user.Password = m.HashPass(pass)
	user.Create(&user)
	c.Redirect(http.StatusMovedPermanently,"/users")
}
func GetUserEdit(c *gin.Context) {
	id,_ := strconv.Atoi(c.Param("id"))
	u := user.GetUser(id)
	c.HTML(http.StatusOK, "user-edit.html", u)
}
func PostUserEdit(c *gin.Context) {
	user.ID,_ = strconv.Atoi(c.PostForm("id"))
	user.Name = c.PostForm("name")
	user.Yomi = c.PostForm("yomi")
	user.Email = c.PostForm("email")
	user.Role = c.PostForm("role")
	user.Edit(&user)
	c.Redirect(http.StatusMovedPermanently,"/users")
}
func PostUserDelete(c *gin.Context) {
	user.ID,_ = strconv.Atoi(c.PostForm("id"))
	user.Delete(&user)
	c.Redirect(http.StatusMovedPermanently,"/users")
}
func PostUserRestore(c *gin.Context) {
	user.ID,_ = strconv.Atoi(c.PostForm("id"))
	user.Restore(&user)
	c.Redirect(http.StatusMovedPermanently,"/users")
}
func PostUserForceDelete(c *gin.Context) {
	user.ID,_ = strconv.Atoi(c.PostForm("id"))
	user.ForceDelete(&user)
	c.Redirect(http.StatusMovedPermanently,"/users")
}