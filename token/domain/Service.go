package domain

import (
	"apiSecurity/helper/domain"
	userDomain "apiSecurity/user/domain"
	"net/http"
)

type Service interface {
	GenerateToken(user *userDomain.User) *domain.Response
	VerifyToken(r *http.Request) *domain.Response
}
