package testing

import (
	"bytes"
	"fmt"
	"goserviceJenkinsDocker/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// func ConnectDb(merchantDb string) (mongoSession *mgo.Session) {
// 	fmt.Println("Trying too Connect....")
// 	mongoDBDialInfo := &mgo.DialInfo{
// 		Addrs:    []string{"mongo:27017"},
// 		Timeout:  60 * time.Second,
// 		Database: merchantDb,
// 	}

// 	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
// 	if err != nil {
// 		log.Fatalf("CreateSession: %s\n", err)
// 	}
// 	mongoSession.SetMode(mgo.Monotonic, true)
// 	fmt.Println("Connected")
// 	return mongoSession
// }

// func databaseCollection(collection string) error {
// 	fmt.Println("In database function function")
// 	mongoSession := ConnectDb("test_jenkins")
// 	defer mongoSession.Close()

// 	sessionCopy := mongoSession.Copy()
// 	defer sessionCopy.Close()

// 	getCollection := sessionCopy.DB("test_jenkins").C(collection)
// 	err := getCollection.Create(nil)
// 	fmt.Println(err)
// 	return err
// }

func TestCreateEntry(t *testing.T) {
	fmt.Println("enter the function")
	gin.SetMode(gin.TestMode)

	// Setup your router, just like you did in your main function, and
	// register your routes.
	var jsonStr = []byte(`{"name":"xyz","email":"xyz@pqr.com"}`)
	r := gin.Default()
	fmt.Println("database funciton")
	// databaseCollection("users")
	r.POST("/api/user", controllers.SaveUserData)

	req, err := http.NewRequest(http.MethodPost, "/api/user", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()
	fmt.Println("after req 11111111111", req, err)

	// Perform the request
	r.ServeHTTP(w, req)
	fmt.Println("after req 22222222222", req, err)

	// Check to see if the response was what you expected
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}
