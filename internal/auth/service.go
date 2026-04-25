package auth

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"WorkWorth/internal/users"
)

type AccessTokenClaim struct {
	UserId int64  `json:"user_id"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

func CreateAccessToken(user *users.User, secret string, expiry int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry))
	claims := AccessTokenClaim{
		UserId: user.ID,
		Name:   user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

type AuthService struct {
	userRepo  users.UserRepository
	jwtSecret string
	jwtExpiry int
}

func NewAuthService(repo users.UserRepository, secret string, expiry int) *AuthService {
	return &AuthService{
		userRepo:  repo,
		jwtSecret: secret,
		jwtExpiry: expiry,
	}
}

func (s *AuthService) SignUp(request SignUpRequest) (*LoginResponse, error) {
	if request.Password != request.ConfirmPassword {
		return nil, errors.New("Password and confirm password do not match")
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	createUserRequest := &users.CreateUserRequest{
		Name:                     request.Name,
		Email:                    request.Email,
		Password:                 string(encryptedPassword),
		MonthlySalary:            request.MonthlySalary,
		MonthlyWorkingDays:       request.MonthlyWorkingDays,
		HoursPerDay:              request.HoursPerDay,
		PreferredCurrency:        request.PreferredCurrency,
		FixedExpenses:            request.FixedExpenses,
		EstimateVariableExpenses: request.EstimateVariableExpenses,
		TargetMonthlySaving:      request.TargetMonthlySaving,
	}

	user, err := s.userRepo.Create(createUserRequest)
	if err != nil {
		return nil, err
	}

	accessToken, err := CreateAccessToken(user, s.jwtSecret, s.jwtExpiry)
	if err != nil {
		return nil, err
	}

	signupResponse := LoginResponse{
		AccessToken: accessToken,
	}

	return &signupResponse, nil
}

func (s *AuthService) Login(request LoginRequest) (*LoginResponse, error) {
	user, err := s.userRepo.GetUserByEmail(request.Email)
	if err != nil {
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		return nil, errors.New("Invalid credentials")
	}

	accessToken, err := CreateAccessToken(user, s.jwtSecret, s.jwtExpiry)
	if err != nil {
		return nil, err
	}

	loginResponse := LoginResponse{
		AccessToken: accessToken,
	}

	return &loginResponse, nil
}
