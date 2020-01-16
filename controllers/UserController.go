package controllers

import (
	"encoding/json"
	"fmt"
	"goserviceJenkinsDocker/config"
	"goserviceJenkinsDocker/models"

	"github.com/gin-gonic/gin"
)

/*
 * Function to save user data
 */
func SaveUserData(c *gin.Context) {
	var user = models.User{}
	userErr := json.NewDecoder(c.Request.Body).Decode(&user)
	if userErr != nil {
		c.JSON(config.FailureCode, gin.H{
			"error": "Data is not in proper format.",
		})
	}
	userId, err := models.GetAutoIncrementCounter(config.UsersCounterId, config.UsersCollection)
	if err != nil {
		c.JSON(config.FailureCode, gin.H{
			"error": "There is something wrong please try after sometime.",
		})
	}
	user.Id = userId
	fmt.Println("Reached in controllers", userId)
	saveUserResp := models.SaveUserData(user)
	if saveUserResp != nil {
		c.JSON(config.FailureCode, gin.H{
			"error": "There is something wrong please try it later.",
		})
	} else {
		c.JSON(config.SuccessCode, gin.H{
			"success": "Data saved successfully.",
			"data":    user,
		})
	}
}
