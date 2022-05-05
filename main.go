package main

import (
	"WCPool/controller"
	"WCPool/driver"
	"database/sql"
	"fmt"
	"net/http"
	"sync"

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

	// match
	matchController := controller.MatchController{}
	router.HandleFunc("/match/{matchday}", matchController.GetMatchesByMatchday(db)).Methods("GET")
	router.HandleFunc("/match", matchController.AddMatches(db)).Methods("POST")

	// party
	partyController := controller.PartyController{}
	router.HandleFunc("/party/{id}", partyController.GetParty(db)).Methods("GET")

	// user
	userController := controller.UserController{}
	router.HandleFunc("/member/{partyID}", userController.AddMemberToParty(db)).Methods("POST")
	router.HandleFunc("/member/{partyID}/{userID}", userController.RemoveUserFromParty(db)).Methods("DELETE")
	router.HandleFunc("/member", userController.UpdateUser(db)).Methods("PUT")

	// prediction
	predictionController := controller.PredictionController{}
	router.HandleFunc("/prediction", predictionController.AddPredictions(db)).Methods("POST")

	wg := sync.WaitGroup{}
	wg.Add(2)
	go runServer(router, &wg)

	// Security Server
	secRouter := mux.NewRouter()
	secController := controller.SecurityController{}
	secRouter.HandleFunc("/signin", secController.VerifyUser(db)).Methods("POST")
	secRouter.HandleFunc("/signup", secController.CreateUser(db)).Methods("POST")
	secRouter.HandleFunc("/refresh", secController.RefreshToken(db)).Methods("POST")

	go runSecServer(secRouter, &wg)
	wg.Wait()
}

func runServer(router *mux.Router, wg *sync.WaitGroup) {
	fmt.Println("Server is running at port 8080")
	// Add CORS configuration to router
	err := http.ListenAndServe(":8080",
		handlers.CORS(
			handlers.AllowedHeaders(
				[]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(router),
	)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Server 8080 is down")
	}
	wg.Done()
}

func runSecServer(router *mux.Router, wg *sync.WaitGroup) {
	fmt.Println("Server is running at port 8081")
	// Add CORS configuration to router
	err := http.ListenAndServeTLS(":8081", "localhost.crt", "localhost.key",
		handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)(router),
	)
	if err != nil {
		fmt.Println(err)
		// server is down
		fmt.Println("Server 8081 is down")
	}
	wg.Done()
}
