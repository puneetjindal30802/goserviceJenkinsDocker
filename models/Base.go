package models

import (
	"errors"
	"fmt"
	"goserviceJenkinsDocker/config"

	"gopkg.in/mgo.v2/bson"
)

func DbInsert(collectionName string, query interface{}) error {
	fmt.Println("reached to base file")
	mongoSession := config.ConnectDb(config.Database)
	defer mongoSession.Close()

	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	getCollection := sessionCopy.DB(config.Database).C(collectionName)
	err := getCollection.Insert(query)
	return err

}

/*
 *	Convert the mongdb ID into int
 */
func GetAutoIncrementCounter(counterName string, collectionName string) (counterId int, err error) {
	data, _ := GetSingleRecord(collectionName)
	if data == nil {
		error := config.InsertCounterValue(counterName, 0, config.Database)
		if error != nil {
			return 0, error
		}
	}
	counterId, err = config.GetNextSequence(counterName, config.Database)
	if err != nil {
		lastId, queryError := DbLastInsertedId(collectionName)
		if queryError != nil {
			return lastId, err
		}
		return lastId + 1, nil
	}
	return counterId, nil
}

/*
* Function to check coupons exist in database
*
* Used by : GetAutoIncrementCounter function
*
* Params collectionName type string
*
* Returns result    type interface
* &       err        type error.
 */

func GetSingleRecord(collectionName string) (result interface{}, err error) {
	mongoSession := config.ConnectDb(config.Database)
	defer mongoSession.Close()

	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	getCollection := sessionCopy.DB(config.Database).C(collectionName)

	err = getCollection.Find(bson.M{}).Select(bson.M{"_id": 1}).One(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

/*
 * Function to fetch one/all document(s) from the database.
 *
 * Used by all the models.
 *
 * Params collection type string
 *
 * Returns true/false type bool
 * &       err        type error.
 */
func DbLastInsertedId(collectionName string) (result int, err error) {
	mongoSession := config.ConnectDb(config.Database)
	defer mongoSession.Close()

	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	resp := bson.M{}
	getCollection := sessionCopy.DB(config.Database).C(collectionName)
	err = getCollection.Find(bson.M{}).Sort("-_id").One(&resp)

	if err != nil {
		return 1, err
	}
	lastId, ok := resp["_id"].(int)
	if ok == false {
		err = errors.New("Error during type assertion")
		return lastId, err
	}
	return lastId, nil
}
