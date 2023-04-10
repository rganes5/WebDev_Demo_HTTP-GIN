package main

import (
	"html/template"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var userDataBase = map[string]string{
	"Ganesh": "Ganesh@123",
	"Stebin": "Stebin@123",
	"Edwin":  "Edwin@123",
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}
//main function
func main() {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// For accessing all the assets like css and images.
	r.Static("/assets", "./assets")

	// Handle fn
	r.GET("/", homeHandler)
	r.POST("/login", loginCheckHandler)
	r.GET("/welcome", welcomeHandler)
	r.GET("/logout", logoutHandler)

	r.Run(":9191")
}

func homeHandler(c *gin.Context) {
	session := sessions.Default(c)
	userName := session.Get("userName")
	if userName != nil {
		data := map[string]interface{}{
			"userName": userName,
		}
		tmpl.ExecuteTemplate(c.Writer, "login.html", data)
	} else {
		tmpl.ExecuteTemplate(c.Writer, "index.html", nil)
	}
}

func loginCheckHandler(c *gin.Context) {
	session := sessions.Default(c)
	userName := session.Get("userName")
	if userName != nil {
		data := map[string]interface{}{
			"userName": userName,
		}
		tmpl.ExecuteTemplate(c.Writer, "login.html", data)
	} else {
		userName := c.PostForm("name")
		password := c.PostForm("password")

		if userDataBase[userName] == password && password != "" {
			session.Set("userName", userName)
			session.Save()

			c.Redirect(http.StatusSeeOther, "/welcome")
		} else {
			data := map[string]interface{}{
				"error": "Invalid username and password",
			}
			tmpl.ExecuteTemplate(c.Writer, "index.html", data)
		}
	}
}

func welcomeHandler(c *gin.Context) {
	session := sessions.Default(c)
	userName := session.Get("userName")
	if userName != nil {
		data := map[string]interface{}{
			"userName": userName,
		}
		tmpl.ExecuteTemplate(c.Writer, "login.html", data)
	} else {
		c.Redirect(http.StatusSeeOther, "/")
	}
}

func logoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusSeeOther, "/")
}
