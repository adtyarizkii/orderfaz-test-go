package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	authdto "orderfaz-test-go/dto/auth"
	dto "orderfaz-test-go/dto/result"
	"orderfaz-test-go/models"
	"orderfaz-test-go/pkg/bcrypt"
	jwtToken "orderfaz-test-go/pkg/jwt"
	"orderfaz-test-go/repositories"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(authdto.RegisterRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	msisdn := request.MSISDN

	// msisdn start with "62"
	if len(msisdn) >= 1 && msisdn[0] == '0' {
		msisdn = "62" + msisdn[1:]
	} else if len(msisdn) >= 2 && msisdn[:2] != "62" {
		msisdn = "62" + msisdn
	}

	user := models.User{
		MSISDN:   msisdn,
		Name:     request.Name,
		Username: request.Username,
		Password: password,
	}

	// Check MSISDN
	regEm, err := h.AuthRepository.Login(user.MSISDN)
	fmt.Println(regEm)
	if err == nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: 403, Message: "MSISDN / Phone Number has already been registered!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Check Username
	regUname, err := h.AuthRepository.CheckUsername(user.Username)
	fmt.Println(regUname)
	if err == nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: 403, Message: "Username unavailable!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.AuthRepository.Register(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	registerResponse := authdto.RegisterResponse{
		MSISDN:   data.MSISDN,
		Name:     data.Name,
		Username: data.Username,
		Password: data.Password,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: registerResponse}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerAuth) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(authdto.LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: 400, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	msisdn := request.MSISDN

	// msisdn start with "62"
	if len(msisdn) >= 1 && msisdn[0] == '0' {
		msisdn = "62" + msisdn[1:]
	} else if len(msisdn) >= 2 && msisdn[:2] != "62" {
		msisdn = "62" + msisdn
	}

	user := models.User{
		MSISDN:   msisdn,
		Password: request.Password,
	}

	// Check MSISDN
	user, err := h.AuthRepository.Login(user.MSISDN)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: 403, Message: "MSISDN / Phone number not found!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Check password
	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: 404, Message: "Password not match!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	//generate token
	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2 hours expired

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		fmt.Println("Unauthorize")
		return
	}

	loginResponse := authdto.LoginResponse{
		MSISDN:   user.MSISDN,
		Name:     user.Name,
		Username: user.Username,
		Token:    token,
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Code: http.StatusOK, Data: loginResponse}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerAuth) CheckAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// Check User by Id
	user, err := h.AuthRepository.Getuser(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	CheckAuthResponse := authdto.CheckAuthResponse{
		ID:       userId,
		MSISDN:   user.MSISDN,
		Name:     user.Name,
		Username: user.Username,
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Code: http.StatusOK, Data: CheckAuthResponse}
	json.NewEncoder(w).Encode(response)
}
