package app

import (
	"lambda-func/api"
	"lambda-func/database"
)

type App struct {
	ApiHandler api.ApiHandler
}

func NewApp() App {
	// we actually initialise our DB Store
	// gets passed down into the API Handler
	db := database.NewDynamoDBClient()
	apiHandler := api.NewApiHandler(db)

	return App{
		ApiHandler: apiHandler,
	}
}
