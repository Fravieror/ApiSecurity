package domain

import (
	"apiSecurity/helper/domain"
	userDomain "apiSecurity/user/domain"
)

type Service interface {
	GenerateToken(user *userDomain.User) *domain.Response
}
