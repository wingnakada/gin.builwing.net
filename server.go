package main

import (
	"github.com/gin-gonic/gin"
	c "gin/controllers"		//home/hide/go/src/gin/controllers
	m "gin/models"			//home/hide/go/src/gin/models
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func main() {
	r := gin.Default()
	//r := gin.New()
	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())
	r.LoadHTMLGlob("views/*.html")
	//r.Static("/public/css/", "./public/css")
	//r.Static("/public/js/", "./public/js/")
	//r.Static("/public/fonts/", "./public/fonts/")
	//r.Static("/public/img/", "./public/img/")
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("GinSession", store))
	r.GET("/", c.GetIndex)
	r.GET("/index", c.GetIndex)
	r.GET("/signin", c.GetSignin)
	r.POST("/signin", c.PostSignin)
	r.GET("/signout", c.GetSignout)
	v1 := r.Group("/users")
	{
		v1.Use(m.Authlizer("/users"))
		v1.GET("", c.GetUsers)
		v1.GET("/create", c.GetUserCreate)
		v1.GET("/edit/:id", c.GetUserEdit)
		v1.GET("/trashed", c.GetUsersTrashed)
		v1.POST("/edit", c.PostUserEdit)
		v1.POST("/create", c.PostUserCreate)
		v1.POST("/search", c.UserSearch)
		v1.POST("/delete", c.PostUserDelete)
		v1.POST("/forcedelete", c.PostUserForceDelete)
		v1.POST("/restore", c.PostUserRestore)
	}
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}