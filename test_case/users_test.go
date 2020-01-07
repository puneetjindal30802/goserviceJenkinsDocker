package testing

import (
	"bytes"
	"goserviceJenkinsDocker/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateEntry(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Setup your router, just like you did in your main function, and
	// register your routes.
	var jsonStr = []byte(`{"name":"xyz","email":"xyz@pqr.com"}`)
	r := gin.Default()
	r.POST("/api/user", controllers.SaveUserData)

	req, err := http.NewRequest(http.MethodPost, "/api/user", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check to see if the response was what you expected
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}
