package main

import (
	"github.com/go-oauth2/mongo"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"

	"log"
	"net/http"
)

/*type Config struct{
	Client_Id="1da3e3e57166dcfd116a"
	Client_Secret="806d9f4a49de69272bfaf24e5b4eb9afdebed5d9"
}*/

func main() {
	manager := manage.NewDefaultManager()
	// using mongodb token store
	manager.MustTokenStorage(mongo.NewTokenStore(mongo.NewConfig(
		"mongodb://127.0.0.1:27017",
		"oauth2",
	)),
	)

	// client memory store
	clientStore := store.NewClientStore()
	clientStore.Set("1da3e3e57166dcfd116a", &models.Client{
		ID:     "1da3e3e57166dcfd116a",
		Secret: "806d9f4a49de69272bfaf24e5b4eb9afdebed5d9",
		Domain: "http://localhost:9090",
	})
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		srv.HandleTokenRequest(w, r)
	})

	log.Fatal(http.ListenAndServe(":9092", nil))
}
