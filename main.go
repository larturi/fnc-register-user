package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	events "github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"

	"github.com/larturi/golang-fnc-register-user/awsgo"
	"github.com/larturi/golang-fnc-register-user/bd"
	"github.com/larturi/golang-fnc-register-user/models"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InitAWS()

	if !ValidoParametros() {
		fmt.Println("Error en los parametros. Debe enviar SecretName")
		err := errors.New("Error en los parametros. Debe enviar SecretName")
		return event, err
	}

	var datos models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email: " + datos.UserEmail)

		case "sub":
			datos.UserUUID = att
			fmt.Println("Sub: " + datos.UserUUID)
		}
	}

	err := bd.ReadSecret()

	if err != nil {
		fmt.Println("Error reading secret. " + err.Error())
		return event, err
	}

	err = bd.SignUp(datos)

	return event, err
}

func ValidoParametros() bool {
	var traeParametro bool
	_, traeParametro = os.LookupEnv("SecretName")
	return traeParametro
}
