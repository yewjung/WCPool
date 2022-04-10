package main

import (
	"WCPool/controller"
	"WCPool/driver"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {
	db = driver.ConnectDB()
	router := mux.NewRouter()
	controller := controller.Controller{}
	router.HandleFunc("/party/{id}", controller.GetParty(db)).Methods("GET")
	router.HandleFunc("/member/{partyID}", controller.AddMemberToParty(db)).Methods("POST")
	router.HandleFunc("/member/{partyID}/{userID}", controller.RemoveUserFromParty(db)).Methods("DELETE")
	router.HandleFunc("/member", controller.UpdateUser(db)).Methods("PUT")
	router.HandleFunc("/match/{matchday}", controller.GetMatchesByMatchday(db)).Methods("GET")
	router.HandleFunc("/prediction", controller.AddPredictions(db)).Methods("POST")
	fmt.Println("Server is running at port 8080")

	// Add CORS configuration to router
	log.Fatal(
		http.ListenAndServe(":8080",
			handlers.CORS(
				handlers.AllowedHeaders(
					[]string{"X-Requested-With", "Content-Type", "Authorization"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
				handlers.AllowedOrigins([]string{"*"}),
			)(router),
		),
	)
}
