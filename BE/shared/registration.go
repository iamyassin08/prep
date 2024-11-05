package shared

import (
	"context"
	"fmt"

	"github.com/Nerzal/gocloak/v13"
	"github.com/go-playground/validator/v10"
)

type RegistrationResponse struct {
	User *gocloak.User
}

type LoginResponse struct {
	ResToken *gocloak.JWT
}

type RegistrationRequest struct {
	Username  string `validate:"required,min=3,max=16"`
	Password  string `validate:"required"`
	FirstName string `validate:"required,min=1,max=30"`
	LastName  string `validate:"required,min=1,max=30"`
	Email     string `validate:"required,email"`
}

type LoginRequest struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

type registerUseCase struct {
	identityManager identityManager
}

func NewRegistraterUseCase(im identityManager) *registerUseCase {
	return &registerUseCase{
		identityManager: im,
	}
}

func (uc *registerUseCase) Login(ctx context.Context, request LoginRequest) (*LoginResponse, error) {
	var validate = validator.New()
	err := validate.Struct(request)
	if err != nil {
		return nil, err
	}
	usrJwt, err := uc.identityManager.LoginUser(ctx, request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	return &LoginResponse{ResToken: usrJwt}, nil
}

// Register user to Keycloak
func (uc *registerUseCase) Register(ctx context.Context, request RegistrationRequest) (*RegistrationResponse, error) {
	var validate = validator.New()
	err := validate.Struct(request)
	if err != nil {
		return nil, err
	}
	var user = gocloak.User{
		Username:      gocloak.StringP(request.Username),
		FirstName:     gocloak.StringP(request.FirstName),
		LastName:      gocloak.StringP(request.LastName),
		Email:         gocloak.StringP(request.Email),
		EmailVerified: gocloak.BoolP(true),
		Enabled:       gocloak.BoolP(true),
		// Attributes:    &map[string][]string{},
	}
	fmt.Println("Reachted User Registration", user)
	userResponse, err := uc.identityManager.CreateUser(ctx, user, request.Password, "freetier")
	if err != nil {
		return nil, err
	}
	var response = &RegistrationResponse{User: userResponse}
	return response, nil
}
