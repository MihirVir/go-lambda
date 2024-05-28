package main

import (
	"fmt"
	"lambda-func/app"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Username string `json:"username"`
}

// Takes some payload and do some processing
func HandleRequest(event MyEvent) (string, error) {
	if event.Username == "" {
		return "", fmt.Errorf("username can't be empty")
	}

	// returns a string
	return fmt.Sprintf("Successfully called by - %s", event.Username), nil
}

func main() {
	myApp := app.NewApp()
	lambda.Start(myApp.ApiHandler.RegisterUserHandler)
}
