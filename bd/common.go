package bd

import (
	"os"

	"github.com/larturi/golang-fnc-register-user/awssecret"
	"github.com/larturi/golang-fnc-register-user/models"
)

var SecretModel models.SecretRDSJson
var err error

func ReadSecret() error {
	SecretModel, err = awssecret.GetSecret(os.Getenv("SecretName"))
	return err
}
