package application

import (
	helperDomain "apiSecurity/helper/domain"
	"apiSecurity/token/domain"
	userDomain "apiSecurity/user/domain"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

type Bussines struct {
	repository domain.Repository
}

func NewBussines(rep domain.Repository) domain.Service {
	return &Bussines{
		repository: rep,
	}
}

func (b *Bussines) GenerateToken(user *userDomain.User) *helperDomain.Response {
	var response = helperDomain.NewResponse()
	ok, err := b.repository.ValidUser(user)
	if err != nil {
		response.Message = "Internal server error"
		response.Code = http.StatusInternalServerError
		return response
	}
	if !ok {
		response.Message = "Invalid"
		response.Code = http.StatusUnauthorized
		return response
	}
	token, err := CreateToken(user.Email)
	if err != nil {
		response.Message = "Internal server error"
		response.Code = http.StatusInternalServerError
		return response
	}
	response.Data, err = json.Marshal(token)
	if err != nil {
		log.Println(err)
		response.Message = "Internal server error"
		response.Code = http.StatusInternalServerError
		return response
	}
	return response
}
func CreateToken(emailUser string) (*domain.Token, error) {
	td := &domain.Token{}
	td.AtExpires = time.Now().Add(time.Minute * 1440).Unix()
	td.AccessUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	var err error
	//Creating Access Token

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["email_user"] = emailUser
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["email_user"] = emailUser
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func (b *Bussines) VerifyToken(r *http.Request) *helperDomain.Response {
	var response = helperDomain.NewResponse()
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			response.Message = fmt.Sprintf("unexpected signing method: %v", token.Header["alg"])
			response.Code = http.StatusUnauthorized
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		response.Message = err.Error()
		response.Code = http.StatusUnauthorized
		return response
	}
	if !token.Valid {
		response.Message = "Unauthorized"
		response.Code = http.StatusUnauthorized
		return response
	}
	response.Data, err = json.Marshal(token)
	if err != nil {
		log.Println(err)
		response.Message = "Internal server error"
		response.Code = http.StatusInternalServerError
		return response
	}
	return response
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
