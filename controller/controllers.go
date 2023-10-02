package controller

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"local/auth"
	"local/database"
	"local/model"
	"net/http"
	"strconv"
)

// LoginPayload login body
// LoginPayload is a struct that contains the fields for a user's login credentials
type LoginPayload struct {
	Email    string `param:"email" query:"email" form:"email" json:"email" binding:"required"`
	Password string `param:"password" query:"password" form:"password"  json:"password" binding:"required"`
}

// LoginResponse token response
// LoginResponse is a struct that contains the fields for a user's login response
type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshtoken"`
}

// Signup is a function that handles user signup
// It takes in an echo context as an argument and binds the user data from the request body to a user struct
// It then hashes the user's password and creates a user record in the database
// If successful, it returns a 200 status code with a success message
// If unsuccessful, it returns a 400 or 500 status code with an error message

func Signup(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return c.String(http.StatusBadRequest, "Wrong user body")
	}
	if err := user.HashPassword(user.Password); err != nil {
		return c.String(http.StatusBadRequest, "Bad body request")
	}
	if err := user.CreateUserRecord(); err != nil {
		return c.String(http.StatusBadRequest, "Current email is already taken")
	}
	return c.JSON(http.StatusOK, "Registred")
}

// Login is a function that handles user login
// It takes in a gin context as an argument and binds the user data from the request body to a LoginPayload struct
// It then checks if the user exists in the database and if the password is correct
// If successful, it generates a token and a refresh token and returns a 200 status code with the token and refresh token
// If unsuccessful, it returns a 401 or 500 status code with an error message

func Login(c echo.Context) error {
	payload := new(LoginPayload)
	if err := c.Bind(payload); err != nil {
		return c.String(http.StatusBadRequest, "Wrong payload. Please set email and password")
	}
	result, err := database.GlobalDB.FindOneFrom(model.UserTable, "email", payload.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "No user with such email")
	}
	if result == nil {
		return c.JSON(http.StatusNotFound, "Wrong email/password (user not found)")
	}

	user := result.(*model.User)

	if err := user.CheckPassword(payload.Password); err != nil {
		return c.String(http.StatusBadRequest, "Wrong password")
	}

	config, err := godotenv.Read()
	if err != nil {
		return c.JSON(500, "Error reading .env file")
	}

	jwtWrapper := auth.JwtWrapper{
		SecretKey:         config["SECRET_KEY"],
		Issuer:            "AuthService",
		ExpirationMinutes: 30,
		ExpirationHours:   24,
	}
	signedToken, err := jwtWrapper.GenerateToken(user.Password)
	if err != nil {
		return c.String(http.StatusBadRequest, "GenerateToken error")
	}

	signedtoken, err := jwtWrapper.RefreshToken(user.Password)
	if err != nil {
		return c.String(http.StatusBadRequest, "RefreshToken error")
	}

	tokenResponse := LoginResponse{
		Token:        signedToken,
		RefreshToken: signedtoken,
	}
	return c.JSON(http.StatusOK, tokenResponse)
}

// e.PUT("/users/", store.updateUser)
func UpdateUser(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "Bad body request")
	}
	if err := database.GlobalDB.Update(u); err != nil {
		return c.String(http.StatusBadRequest, "No such user")
	}
	return c.JSON(http.StatusOK, "Updated")
}

// e.POST("/users", store.saveUser)
func SaveUser(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	if err := database.GlobalDB.Save(u); err != nil {
		return c.String(http.StatusBadRequest, "User not saved")
	}
	return c.JSON(http.StatusOK, "Saved")
}

// e.DELETE("/users/:id", store.deleteUser)
func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	res, err := database.GlobalDB.DeleteFrom(model.UserTable, fmt.Sprintf("where id = %d", id))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "Удалено строк "+strconv.Itoa(int(res)))
}

// e.GET("/users/:id", store.getUser)
func GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	res, err := database.GlobalDB.FindAllFrom(model.UserTable, "id", id)
	if res == nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}
	return c.JSON(http.StatusOK, res)
}

// e.GET("/", store.getUsers)
func GetUsers(c echo.Context) error {
	res, err := database.GlobalDB.SelectAllFrom(model.UserTable, "")
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
