package main

import (
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
	clientStore.Set("example-app", &models.Client{
		ID:     "example-app",
		Secret: "ZXhhbXBsZS1hcHAtc2VjcmV0",
		Domain: "https://localhost:5555",
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

	log.Println("Server is running at 5556 port.")
	//fmt.Println(token.AccessToken)
	log.Fatal(http.ListenAndServe(":5556", nil))
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

/* https://accounts.google.com/o/oauth2/auth?client_id=118778695955-4inip14ellc053slc8e3aebndgcfo3p9.apps.googleusercontent.com&scope=openid+email+profile&response_type=code&openid.realm=https://secure.imdb.com&state=eyI0OWU2YyI6IjExNjMiLCJ1IjoiaHR0cHM6Ly93d3cuaW1kYi5jb20vP3JlZl89bG9naW4iLCJtYW51YWxMaW5rIjoiZmFsc2UifQ&redirect_uri=https://secure.imdb.com/registration/googlehandler&from_login=1&as=-3c14d0634982220d*/
