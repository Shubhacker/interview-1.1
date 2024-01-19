package apis

import (
	"log"
	"net/http"

	"interview/database"
	"interview/utils"

	"github.com/gorilla/mux"
)

func Apis() {
	var errorIn utils.ErrorIntup
	database.Setup()

	router := mux.NewRouter()
	router.HandleFunc("/createUser", CreateUser).Methods("POST")
	router.HandleFunc("/updateUser", UpdateUser).Methods("POST")
	log.Println("Server running on : 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		errorIn.ErrorMsg = err.Error()
		errorIn.FailFunction = "ListenAndServe"
		errorIn.Error()
	}
}