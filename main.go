package main

import (
	"context"

	events "github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/larturi/golang-fnc-register-user/awsgo"
)

func main() {
	lambda.Start(ejecutoLambda)
}

func ejecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InitAWS()
}
