package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	//"fmt"

	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
	"gopkg.in/session.v1"
)

var (
	globalSessions *session.Manager
)

func init() {
	globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid","gclifetime":3600}`)
	go globalSessions.GC()
}

func main() {
	manager := manage.NewDefaultManager()
	// token store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	clientStore := store.NewClientStore()
	clientStore.Set("x6c0971e3c5016f020e0f", &models.Client{
		ID:     "x6c0971e3c5016f020e0f",
		Secret: "x5dd8069b23820aaad1afb6c215f6eb5cd72f33c2",
		Domain: "http://localhost:8888",
	})
	manager.MapClientStorage(clientStore)

	srv := server.NewServer(server.NewConfig(), manager)
	srv.SetUserAuthorizationHandler(userAuthorizeHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			fmt.Println(r.Body)
			us, err := globalSessions.SessionStart(w, r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			us.Set("LoggedInUserID", "test")
			w.Header().Set("Location", "/auth")
			w.WriteHeader(http.StatusFound)
			return
		}
		outputHTML(w, r, "html/login.html")
	})
	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		us, err := globalSessions.SessionStart(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if us.Get("LoggedInUserID") == nil {
			w.Header().Set("Location", "/login")
			w.WriteHeader(http.StatusFound)
			return
		}
		if r.Method == "POST" {
			form := us.Get("Form").(url.Values)
			fmt.Println(r.Body)
			u := new(url.URL)
			u.Path = "/authorize"
			u.RawQuery = form.Encode()
			w.Header().Set("Location", u.String())
			w.WriteHeader(http.StatusFound)
			us.Delete("Form")
			us.Set("UserID", us.Get("LoggedInUserID"))
			return
		}
		outputHTML(w, r, "html/auth.html")
	})
	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})
	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleTokenRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("Server is running at 6788 port.")
	//fmt.Println(token.AccessToken)
	log.Fatal(http.ListenAndServe("localhost:6788", nil))
}

func userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	us, err := globalSessions.SessionStart(w, r)
	uid := us.Get("UserID")
	if uid == nil {
		if r.Form == nil {
			r.ParseForm()
		}
		us.Set("Form", r.Form)
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}
	userID = uid.(string)
	us.Delete("UserID")
	return
}

func outputHTML(w http.ResponseWriter, req *http.Request, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()
	fi, _ := file.Stat()
	http.ServeContent(w, req, file.Name(), fi.ModTime(), file)
}
