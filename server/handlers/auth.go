package handlers

import (
	// "encoding/json"
	// "fmt"
	authdto "housy/dto/auth"
	dto "housy/dto/result"
	"housy/models"
	"housy/pkg/bcrypt"
	jwtToken "housy/pkg/jwt"
	"housy/repositories"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"

	// "github.com/golang-jwt/jwt/v4/request"
	"github.com/labstack/echo/v4"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) SignUp(c echo.Context) error {
	request := new(authdto.SignUpRequest)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "error1" + err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "error2" + err.Error()})
	}

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "error3" + err.Error()})
	}

	user := models.User{
		Fullname:   request.Fullname,
		Username:   request.Username,
		Email:      request.Email,
		Password:   password,
		ListAsRole: request.ListAsRole,
		Gender:     request.Gender,
		Phone:      request.Phone,
		Address:    request.Address,
		// Image:      filename,
	}

	data, err := h.AuthRepository.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: "error4" + err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func (h *handlerAuth) SignIn(c echo.Context) error {
	request := new(authdto.SignInRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	userLogin := models.User{
		Username: request.Username,
		Password: request.Password,
	}

	// Check email/username
	user, err := h.AuthRepository.SignIn(userLogin.Username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "error 1" + err.Error()})
	}

	// Check password
	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	//generate token
	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2 hours expired

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		// fmt.Println("Unauthorize")
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	signInResponse := authdto.SignInResponse{
		ID:         user.ID,
		Username:   user.Username,
		ListAsRole: user.ListAsRole,
		Fullname:   user.Fullname,
		Email:      user.Email,
		Password:   user.Password,
		Gender:     user.Gender,
		Phone:      user.Phone,
		Address:    user.Address,
		// Image:      user.Image,
		Token: token,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: signInResponse})

}

func (h *handlerAuth) CheckAuth(c echo.Context) error {
	userInfo := c.Get("userLogin").(jwt.MapClaims)
	// userInfo := c.Request().Context().Value("userLogin").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// Check User by Id
	user, err := h.AuthRepository.Getuser(userId)
	if err != nil {
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		return c.JSON(http.StatusBadRequest, response)
	}

	CheckAuthResponse := authdto.CheckAuthResponse{
		ID:         user.ID,
		Fullname:   user.Fullname,
		Email:      user.Email,
		Password:   user.Password,
		Username:   user.Username,
		ListAsRole: user.ListAsRole,
		Gender:     user.Gender,
		Phone:      user.Phone,
		Address:    user.Address,
		Token:      "",
	}

	response := dto.SuccessResult{Code: http.StatusOK, Data: CheckAuthResponse}
	return c.JSON(http.StatusOK, response)
}
