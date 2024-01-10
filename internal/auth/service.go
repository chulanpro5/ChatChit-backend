package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"test-chat/pkg/common"
	"test-chat/pkg/entity"
	"test-chat/pkg/util"
	"time"
)

const SecretKey = "secret"

type Service struct {
	common *common.Common
}

func NewAuthService(common *common.Common) *Service {
	return &Service{
		common: common,
	}
}

func (s *Service) Register(dto RegisterRequest) (entity.User, error) {

	password, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), 14) //GenerateFromPassword returns the bcrypt hash of the password at the given cost i.e.

	user := entity.User{
		Name:               dto.Name,
		Email:              dto.Email,
		Password:           password,
		ProfileImageBase64: util.ImageToBase64(util.GenerateAvatar(dto.Email)),
		LanguageId:         2,
	}

	s.common.Database.Create(&user)

	return user, nil
}

func (s *Service) Login(dto LoginRequest) (fiber.Cookie, *entity.User, error) {
	var user entity.User

	s.common.Database.Where("email = ?", dto.Email).First(&user) //Check the email is present in the DB

	if user.ID == 0 { //If the ID return is '0' then there is no such email present in the DB
		return fiber.Cookie{}, &entity.User{}, &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "email not found",
		}
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(dto.Password)); err != nil {
		return fiber.Cookie{}, &entity.User{}, &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "incorrect password",
		}
	} // If the email is present in the DB then compare the Passwords and if incorrect password then return error.

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * 24)},
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return fiber.Cookie{}, &entity.User{}, err
	}

	return fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}, &user, nil
}

func (s *Service) Logout(userId string) (fiber.Cookie, error) {
	// Dummy cookie
	return fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}, nil
}
