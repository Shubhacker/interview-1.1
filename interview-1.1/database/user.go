package database

import (
	"interview/structs"
	"interview/utils"

	"github.com/jmoiron/sqlx"
)

func CreateUser(U structs.UserData) error{
	query := "INSERT INTO users (name, email) VALUES (" + U.Name + ", " + U.Email + ")"
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(U structs.UserData) error{
	var inputArgs []interface{}

	query := `Update users set`
	if !utils.IsBlank(U.Name) {
		query += ` name= ?`
		inputArgs = append(inputArgs, U.Name)
	}

	if !utils.IsBlank(U.Email) {
		query += `, email= ?`
		inputArgs = append(inputArgs, U.Email)
	}

	query +=` where id=?`
	inputArgs = append(inputArgs, U.UserId)
	query = sqlx.Rebind(sqlx.DOLLAR, query)
	_, err := db.Exec(query, inputArgs...)
	if err != nil {
		return err
	}

	return nil
}