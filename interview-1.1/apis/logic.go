package apis

import (
	"interview/database"
	"interview/structs"
	"interview/utils"
	"net/http"
)

	var errorIn utils.ErrorIntup

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var User structs.UserData
	User.Name = r.FormValue("name")
	if User.Name == ""{
		errorIn.ErrorMsg = "Name parameter mandatory !"
		errorIn.FailFunction = "FormValue"
		errorIn.Error()
		return
	}
	User.Email = r.FormValue("email")
	if User.Email == ""{
		errorIn.ErrorMsg = "Email parameter mandatory !"
		errorIn.FailFunction = "FormValue"
		errorIn.Error()
		return
	}
	err := database.CreateUser(User)
	if err != nil {
		errorIn.ErrorMsg = err.Error()
		errorIn.FailFunction = "CreateUser"
		errorIn.Error()
		return
	}
}


func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var User structs.UserData
	User.Name = r.FormValue("name")
	User.Email = r.FormValue("email")

	User.UserId = r.FormValue("id")
	if User.UserId == ""{
		errorIn.ErrorMsg = "UserId parameter mandatory !"
		errorIn.FailFunction = "FormValue"
		errorIn.Error()
		return
	}

	err := database.UpdateUser(User)
	if err != nil {
		errorIn.ErrorMsg = err.Error()
		errorIn.FailFunction = "CreateUser"
		errorIn.Error()
		return
	}
}