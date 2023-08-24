package bd

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/larturi/golang-fnc-register-user/models"
	"github.com/larturi/golang-fnc-register-user/tools"
)

func SignUp(signUp models.SignUp) error {
	fmt.Println("Comienza registro de usuario...")

	err := DbConnect()

	if err != nil {
		return err
	}

	defer Db.Close()

	sql := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + signUp.UserEmail + "', '" + signUp.UserUUID + "', '" + tools.FechaMySQL() + ")"
	fmt.Println(sql)

	_, err = Db.Exec(sql)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Registro exitoso!")

	return nil
}
