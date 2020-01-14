package config

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Counter struct {
	Id  string `json:"id" bson:"_id"`
	Seq int    `json:"seq" bson:"seq"`
}

/*
 *	Funtion will connect with the database
 *
 *	Return databse connection.
 */
func ConnectDb(merchantDb string) (mongoSession *mgo.Session) {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{"127.0.0.1:27017"},
		Timeout:  60 * time.Second,
		Database: merchantDb,
	}

	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}
	mongoSession.SetMode(mgo.Monotonic, true)

	return mongoSession
}

/*
* Function to insert the counter id if not exist
*
* Used by : Base::GetAutoIncrementCounter
*
* Params name(counter_id) type string
*
* Returns  err        type error.
 */
func InsertCounterValue(name string, count int, merchantDb string) (err error) {
	mongoSession := ConnectDb(merchantDb)
	defer mongoSession.Close()

	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	getCollection := sessionCopy.DB(merchantDb).C(CountersCollection)
	getCollection.Upsert(bson.M{"_id": name}, bson.M{"$set": bson.M{"seq": count}})
	return nil
}

/*
* Function to get the incremented value of the section
*
* Used by : Base :: GetAutoIncrementCounter
*
* Params name(counter_id) type string
*
* Returns result    type int
* &       err        type error.
 */
func GetNextSequence(name string, merchantDb string) (result int, err error) {
	mongoSession := ConnectDb(merchantDb)
	defer mongoSession.Close()

	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	getCollection := sessionCopy.DB(merchantDb).C(CountersCollection)
	var counter Counter
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"seq": 1}},
		ReturnNew: true,
	}
	_, err = getCollection.Find(bson.M{"_id": name}).Apply(change, &counter)
	newID := counter.Seq
	if err != nil {
		return newID, err
	}
	return newID, nil
}
