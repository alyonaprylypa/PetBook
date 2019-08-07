package main

import (
	"PetBook/pkg/tokens"
	//"database/sql"
	//"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	//"github.com/jmoiron/sqlx"
	"PetBook/gomigrations"
	"github.com/urfave/negroni"
	//"github.com/lib/pq"
	"PetBook/pkg/handler"
	"PetBook/pkg/driver"
	"log"
	"net/http"
	"os"
	//"PetBook/models"
	"PetBook/pkg/controllers"
)

func logErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func main() {
	db := driver.ConnectDB()
	err := gomigrations.Migrate(db)
	if err != nil {
		log.Fatal("Migration failed.")
	}

	router := mux.NewRouter()

	storeUser := controllers.UserStore{DB: db}
	controller := handler.Controller{UserStore: &storeUser}
	storePet:=controllers.PetStore{DB:db}
	controllerPet:=handler.Controller{PetStore:&storePet}

	router.HandleFunc("/register", controller.RegisterPostHandler()).Methods("POST")
	router.HandleFunc("/register", controller.RegisterGetHandler()).Methods("GET")

	router.HandleFunc("/login", controller.LoginPostHandler()).Methods("POST")
	router.HandleFunc("/login", controller.LoginGetHandler()).Methods("GET")

	router.HandleFunc("/petcabinet", controllerPet.CreatePetPostHandler()).Methods("POST")
	router.HandleFunc("/petcabinet", controllerPet.CreatePetGetHandler()).Methods("GET")




	//router.Handle("/cabinetPet", negroni.New(
	//	negroni.HandlerFunc(utils.ValidateTokenMiddleware),
	//	negroni.Wrap(http.HandlerFunc(controller.CreatePetHandler())),
	//))

	router.Handle("/mypage", negroni.New(
		negroni.HandlerFunc(tokens.ValidateTokenMiddleware),
		negroni.Wrap(http.HandlerFunc(controller.MyPageGetHandler())),
	))



	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/",
		http.FileServer(http.Dir("./web/assets/"))))

	router.Handle("/", negroni.New(
		negroni.HandlerFunc(tokens.ValidateTokenMiddleware),
		negroni.Wrap(http.HandlerFunc(controller.MyPageGetHandler())),
	))

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	log.Fatal(http.ListenAndServe(":8080", loggedRouter))
}
