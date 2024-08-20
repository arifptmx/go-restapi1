package controllers

import (
	users "arifptmx/business/user"
	"arifptmx/controller/user/request"
	"arifptmx/utils"
	"io"
	"net/http"
)

type UserController struct {
	usecase users.UserUseCaseInterface
}

// dipasangkan dengan routing
func NewUserController(uc users.UserUseCaseInterface) *UserController {
	return &UserController{
		usecase: uc,
	}
}

// add a user handler
func (controller *UserController) SignUp(res http.ResponseWriter, req *http.Request) {
	// check the method
	if req.Method != "POST" {
		// Add the response return message
		HandlerMessage := []byte(`{
			"success": false,
			"message": "Check your HTTP method: Invalid HTTP method executed"
		}`)

		utils.ReturnJsonResponse(res, http.StatusMethodNotAllowed, HandlerMessage)
	}

	var userSignup request.UserEdit

	// Read the request body
	bodyBytes, err := io.ReadAll(req.Body)
}
