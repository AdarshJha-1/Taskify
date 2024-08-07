package handler

import (
	"encoding/json"
	"net/http"

	"github.com/AdarshJha-1/Taskify/backend/internal/model"
	"github.com/AdarshJha-1/Taskify/backend/internal/repository"
	"github.com/AdarshJha-1/Taskify/backend/internal/response"
	"github.com/AdarshJha-1/Taskify/backend/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Creating local user object
	var user model.User

	// Getting user data and decoding it into local user object
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res := response.Response{Status: http.StatusBadRequest, Success: false, Message: "Invalid input", Data: map[string]interface{}{"error": err.Error()}}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Closing request body
	defer r.Body.Close()

	// Checking if user already exists or not
	userExists := repository.CheckExistingUser(user.Email, user.Username)
	if userExists {
		w.WriteHeader(http.StatusConflict)
		res := response.Response{Status: http.StatusConflict, Success: false, Message: "User already exists", Data: map[string]interface{}{"error": "User already exists"}}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Hashing password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := response.Response{Status: http.StatusInternalServerError, Success: false, Message: "Internal Server Error", Data: map[string]interface{}{"error": err.Error()}}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Creating new user bson objectId and assigning it in user's id field
	user.Id = primitive.NewObjectID()

	// Updating password with hash password
	user.Password = hashedPassword

	// Creating new user with CreateUser function
	result, err := repository.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := response.Response{Status: http.StatusInternalServerError, Success: false, Message: "Failed to create user", Data: map[string]interface{}{"error": err.Error()}}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Sending success message with newly create user id
	w.WriteHeader(http.StatusCreated)
	res := response.Response{Status: http.StatusCreated, Success: true, Message: "User created successfully", Data: map[string]interface{}{"data": result}}
	json.NewEncoder(w).Encode(res)
}

func SignInUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Creating local user object
	var signinUserData model.SignIn

	// Getting user signin data and decoding it into local signinUserData object
	err := json.NewDecoder(r.Body).Decode(&signinUserData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res := response.Response{Status: http.StatusBadRequest, Success: false, Message: "Invalid input", Data: map[string]interface{}{"error": err.Error()}}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Closing request body
	defer r.Body.Close()

	// Fetching user from db
	var user *model.User
	user, err = repository.GetUser(signinUserData.Identifier, signinUserData.Password)
	if err != nil && user == nil {
		w.WriteHeader(http.StatusNotFound)
		res := response.Response{Status: http.StatusNotFound, Success: false, Message: "User not found", Data: map[string]interface{}{"error": err.Error()}}
		json.NewEncoder(w).Encode(res)
		return
	}
	// Checking user password
	isCorrectPassword := utils.CheckPasswordHash(signinUserData.Password, user.Password)
	if !isCorrectPassword {
		w.WriteHeader(http.StatusBadRequest)
		res := response.Response{Status: http.StatusBadRequest, Success: false, Message: "Wrong Credentials", Data: map[string]interface{}{"error": "Wrong Credentials"}}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Creating JWT Token
	token, err := utils.CreateJWT(string(user.Id.Hex()))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := response.Response{Status: http.StatusInternalServerError, Success: false, Message: "Error creating token", Data: map[string]interface{}{"error": err.Error()}}
		json.NewEncoder(w).Encode(res)
		return
	}

	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		MaxAge:   86400,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
	res := response.Response{Status: http.StatusOK, Success: true, Message: "User Signed In successfully", Data: map[string]interface{}{"token": token}}
	json.NewEncoder(w).Encode(res)
}
func SignOutUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Closing request body
	defer r.Body.Close()

	cookie := http.Cookie{
		Name:     "token",
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
	res := response.Response{Status: http.StatusOK, Success: true, Message: "User Signed Out successfully"}
	json.NewEncoder(w).Encode(res)
}
