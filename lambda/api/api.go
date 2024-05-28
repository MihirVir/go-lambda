package api

import (
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
)

type ApiHandler struct {
	dbStore database.DynamoDBClient
}

func NewApiHandler(dbStore database.DynamoDBClient) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (api ApiHandler) RegisterUserHandler(event types.RegisterUser) error {
	if event.Username == "" || event.Password == "" {
		return fmt.Errorf("request has empty params")
	}

	// does a user with this username already exists
	userExists, err := api.dbStore.IsExistingUser(event.Username)

	if err != nil {
		return fmt.Errorf("theres an error checking if user exists %w", err)
	}

	if userExists {
		return fmt.Errorf("user already exists")
	}

	// we know that a user doesn't exist
	err = api.dbStore.InsertUser(event)

	if err != nil {
		return fmt.Errorf("error registering the user")
	}

	return nil
}
