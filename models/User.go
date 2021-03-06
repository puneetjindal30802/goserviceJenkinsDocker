package models

import (
	"fmt"
	"goserviceJenkinsDocker/config"
)

type User struct {
	Id    int    `json:"_id" bson:"_id"`
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}

/*
 *	Function to save user data into database
 *
 *	Return err
 */
func SaveUserData(query interface{}) (err error) {
	fmt.Println("Reached in models", query)
	err = DbInsert(config.UsersCollection, query)

	return err
}
