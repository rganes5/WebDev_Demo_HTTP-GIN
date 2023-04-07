package main

import (
	"html/template"
	"net/http"

	"github.com/icza/session"
)

var userDataBase = map[string]string{
	"Ganesh": "Ganesh@123",
	"Stebin": "Stebin@123",
	"Edwin":  "Edwin@123",
}

var tmpl *template.Template

// type detail struct {
// 	Name string
// 	Age  int
// }

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	sess := session.Get(r)
	if sess != nil {
		userName := sess.CAttr("userName")
		data := map[string]interface{}{
			"userName": userName,
		}
		tmpl.ExecuteTemplate(w, "login.html", data)
	} else {
		tmpl.ExecuteTemplate(w, "index.html", nil)
	}
}

func loginCheckHandler(w http.ResponseWriter, r *http.Request) {
	sess := session.Get(r)
	if sess != nil {
		userName := sess.CAttr("userName")
		data := map[string]interface{}{
			"userName": userName,
		}
		tmpl.ExecuteTemplate(w, "login.html", data)
	} else {
		userName := r.FormValue("name")
		password := r.FormValue("password")

		if userDataBase[userName] == password && password != "" {
			sess := session.NewSessionOptions(&session.SessOptions{
				CAttrs: map[string]interface{}{"userName": userName},
			})
			//creates a session id
			session.Add(sess, w)
			http.Redirect(w, r, "/welcome", http.StatusSeeOther)
		} else {
			data := map[string]interface{}{
				"error": "Invalid username and password",
			}
			// http.Redirect(w, r, "/", http.StatusSeeOther)
			tmpl.ExecuteTemplate(w, "index.html", data)
		}
	}
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache,no-store,must-revalidate")
	// var s detail
	// s.Name = "Ganesh"
	// s.Age = 24
	sess := session.Get(r)
	if sess != nil {
		userName := sess.CAttr("userName")
		data := map[string]interface{}{
			"userName": userName,
		}
		tmpl.ExecuteTemplate(w, "login.html", data)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	sess := session.Get(r)
	if sess != nil {
		session.Remove(sess, w)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	//For accessing all the assets like css and images.
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	//Handle fn
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginCheckHandler)
	http.HandleFunc("/welcome", welcomeHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.ListenAndServe(":9999", nil)
}
