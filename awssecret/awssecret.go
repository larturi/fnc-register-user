package awssecret

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"

	"github.com/larturi/golang-fnc-register-user/awsgo"
	"github.com/larturi/golang-fnc-register-user/models"
)

func GetSecret(nombreSecret string) (models.SecretRDSJson, error) {
	var datosSecret models.SecretRDSJson
	fmt.Println("> Secret: ", nombreSecret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nombreSecret),
	})

	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println("> Lectura Secret OK ", nombreSecret)

	fmt.Println("> dbUser: ", datosSecret.Username)
	fmt.Println("> authToken: ", datosSecret.Password)
	fmt.Println("> dbEndpoint: ", datosSecret.Host)
	fmt.Println("> dbPort: ", datosSecret.Port)
	fmt.Println("> dbClusterIdentifier: ", datosSecret.DbClusterIdentifier)

	return datosSecret, nil
}
